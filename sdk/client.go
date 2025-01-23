package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/availproject/avail-go-sdk/metadata"
	"io"
	"net/http"

	meta "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type Client struct {
	client   *http.Client
	endpoint string
	metadata *metadata.Metadata
	Rpc      RPC
}

func NewClient(endpoint string) *Client {
	client := new(Client)
	client.client = new(http.Client)
	client.endpoint = endpoint
	client.Rpc = newRPC(client)

	return client
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
		hash, err := this.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
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

func (this *Client) Request(method string, params string) (string, error) {
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
		return "", err
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := this.client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	responseBodyBytes, _ := io.ReadAll(response.Body)
	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	// fmt.Println("response Body:", string(responseBodyBytes))

	if response.StatusCode != http.StatusOK {
		return "", errors.New("HTTP status was NOT OK")
	}

	var mappedData map[string]interface{}
	if err := json.Unmarshal(responseBodyBytes, &mappedData); err != nil {
		return "", err
	}

	if mappedData["error"] != nil {
		err := mappedData["error"].(map[string]interface{})
		return "", errors.New(err["message"].(string) + ". " + err["data"].(string))
	}

	if mappedData["result"] == nil {
		return "", nil
	}

	resultBytes, _ := json.Marshal(mappedData["result"])
	result := string(resultBytes)

	// Remove double quotes if there are any
	if len(result) >= 1 {
		if result[0] == '"' && result[len(result)-1] == '"' {
			result = result[1 : len(result)-1]
		}
	}

	return result, nil
}

func (this *Client) Send(tx prim.EncodedExtrinsic) (prim.H256, error) {
	params := "[\"" + tx.ToHexWith0x() + "\"]"

	txHash, err := this.Request("author_submitExtrinsic", params)
	if err != nil {
		return prim.H256{}, err
	}

	return prim.NewH256FromHexString(txHash)
}

func (this *Client) RPCBlockAt(blockHash prim.Option[prim.H256]) (RPCBlock, error) {
	primBlock, err := this.Rpc.Chain.GetBlock(blockHash)
	if err != nil {
		return RPCBlock{}, err
	}
	return NewRPCBlockFromPrimBlock(primBlock)
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
			return RPCBlock{}, err
		}
		extrinsics = append(extrinsics, decoded)
	}

	return RPCBlock{
		Header:     primBlock.Header,
		Extrinsics: extrinsics,
	}, nil
}
