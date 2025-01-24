package sdk

import ()

type authorRPC struct {
	client *Client
}

func (this *authorRPC) RotateKeys() (string, error) {
	params := RPCParams{}
	return this.client.Request("author_rotateKeys", params.Build())
}
