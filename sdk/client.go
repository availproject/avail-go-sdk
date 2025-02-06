package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/availproject/avail-go-sdk/metadata"

	meta "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type Client struct {
	client   *http.Client
	endpoint string
	metadata *metadata.Metadata
	Rpc      RPC
	Call     RuntimeApi
}

func NewClient(endpoint string) *Client {
	client := new(Client)
	client.client = new(http.Client)
	client.endpoint = endpoint
	client.Rpc = newRPC(client)
	client.Call = newRuntimeAPi(client)

	return client
}

func (this *Client) BlockNumber(blockHash prim.H256) (uint32, error) {
	header, err := this.Rpc.Chain.GetHeader(prim.Some(blockHash))
	return header.Number, err
}

func (this *Client) BestBlockNumber() (uint32, error) {
	header, err := this.Rpc.Chain.GetHeader(prim.None[prim.H256]())
	return header.Number, err
}

func (this *Client) FinalizedBlockNumber() (uint32, error) {
	hash, err := this.FinalizedBlockHash()
	if err != nil {
		return uint32(0), err
	}
	header, err := this.Rpc.Chain.GetHeader(prim.Some(hash))
	return header.Number, err
}

func (this *Client) BlockHash(blockNumber uint32) (prim.H256, error) {
	return this.Rpc.Chain.GetBlockHash(prim.Some(blockNumber))
}

func (this *Client) BestBlockHash() (prim.H256, error) {
	return this.Rpc.Chain.GetBlockHash(prim.None[uint32]())
}

func (this *Client) FinalizedBlockHash() (prim.H256, error) {
	return this.Rpc.Chain.GetFinalizedHead()
}

func (this *Client) EventsAt(at prim.Option[prim.H256]) (EventRecords, error) {
	eventsRaw, err := this.Rpc.State.GetEvents(at)
	if err != nil {
		return EventRecords{}, err
	}
	events, err := NewEvents(prim.Hex.FromHex(eventsRaw), this.Metadata())
	if err != nil {
		return EventRecords{}, err
	}

	eventRecord, err := events.Decode()
	if err != nil {
		return EventRecords{}, err
	}

	return eventRecord, nil
}

func (this *Client) StorageAt(at prim.Option[prim.H256]) (BlockStorage, error) {
	if at.IsNone() {
		hash, err := this.Rpc.Chain.GetBlockHash(prim.None[uint32]())
		if err != nil {
			return BlockStorage{}, err
		}
		at.Set(hash)
	}

	return BlockStorage{
		client: this,
		at:     at.Unwrap(),
	}, nil
}

func (this *Client) RPCBlockAt(blockHash prim.Option[prim.H256]) (RPCBlock, error) {
	primBlock, err := this.Rpc.Chain.GetBlock(blockHash)
	if err != nil {
		return RPCBlock{}, err
	}
	return NewRPCBlockFromPrimBlock(primBlock)
}

func (this *Client) InitMetadata(at prim.Option[prim.H256]) error {
	scaleMetadata, err := this.Rpc.State.GetMetadata(at)
	if err != nil {
		return err
	}
	metadata, err := metadata.NewMetadata(scaleMetadata)
	if err != nil {
		return err
	}

	this.metadata = &metadata
	return nil
}

func (this *Client) RequestWithRetry(method string, params string) (string, error) {
	retryCount := 3
	for {
		res, err := this.Request(method, params)
		if err != nil {
			var sdkError *SDKError
			if !errors.As(err, &sdkError) || sdkError.Code != 0 {
				return "", err
			}
		}

		if res.IsSome() {
			return res.Unwrap(), nil
		}

		logger := NewCustomLoggerEmpty("RPC", true)

		if retryCount == 0 {
			logger.LogRpcRetryAbort(method)
			e := ErrorCode005
			e.Message = fmt.Sprintf("Method: %v, Params: %v", method, params)
			return "", e
		}
		logger.LogRpcRetry(method)
		retryCount -= 1
		time.Sleep(time.Second * time.Duration(3))
	}
}

func (this *Client) Request(method string, params string) (prim.Option[string], error) {
	rawJSON := []byte(`{
		"id": 1,
		"jsonrpc": "2.0",
		"method": "%s",
		"params": %s
	}`)

	requestBodyString := fmt.Sprintf(string(rawJSON), method, params)
	requestBodyBytes := []byte(requestBodyString)

	request, err := http.NewRequest("POST", this.endpoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return prim.None[string](), err
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := this.client.Do(request)
	if err != nil {
		return prim.None[string](), newError(err, ErrorCode000)
	}

	defer response.Body.Close()

	responseBodyBytes, _ := io.ReadAll(response.Body)
	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// fmt.Println("response Body:", string(responseBodyBytes))

	if response.StatusCode != http.StatusOK {
		err := ErrorCode001
		err.Message = fmt.Sprintf(`Status Code: %v`, response.StatusCode)
		return prim.None[string](), &err
	}

	var mappedData map[string]interface{}
	if err := json.Unmarshal(responseBodyBytes, &mappedData); err != nil {
		return prim.None[string](), newError(err, ErrorCode002)
	}

	if mappedData["error"] != nil {
		err := mappedData["error"].(map[string]interface{})
		errMessage := ""
		if err["message"] != nil {
			errMessage += err["message"].(string)
		}
		if err["data"] != nil {
			errMessage += " " + err["data"].(string)
		}

		return prim.None[string](), newError(errors.New(errMessage), ErrorCode002)
	}

	if mappedData["result"] == nil {
		return prim.None[string](), nil
	}

	resultBytes, _ := json.Marshal(mappedData["result"])
	result := string(resultBytes)

	// Remove double quotes if there are any
	if len(result) >= 1 {
		if result[0] == '"' && result[len(result)-1] == '"' {
			result = result[1 : len(result)-1]
		}
	}

	return prim.Some(result), nil
}

func (this *Client) Send(tx prim.EncodedExtrinsic) (prim.H256, error) {
	return this.Rpc.Author.SubmitExtrinsic(tx.ToHexWith0x())
}

func (this *Client) Metadata() *meta.Metadata {
	return this.metadata
}

type RPCBlock struct {
	Header     prim.Header
	Extrinsics []prim.DecodedExtrinsic
}

func NewRPCBlockFromPrimBlock(primBlock prim.Block) (RPCBlock, error) {
	extrinsics := []prim.DecodedExtrinsic{}
	for i := 0; i < len(primBlock.Extrinsics); i++ {
		encoded := prim.NewEncodedExtrinsicFromHex(primBlock.Extrinsics[i])
		decoded, err := encoded.Decode(uint32(i))
		if err != nil {
			return RPCBlock{}, newError(err, ErrorCode004)
		}
		extrinsics = append(extrinsics, decoded)
	}

	return RPCBlock{
		Header:     primBlock.Header,
		Extrinsics: extrinsics,
	}, nil
}
