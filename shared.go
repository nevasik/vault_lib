package vault

import (
	"context"
	"encoding/json"
)

func (c *Client) Shared(cfg any) (err error) {
	secret, err := c.C.KVv2("kv").Get(context.Background(), "shared")
	if err != nil {
		return
	}
	data, err := json.Marshal(secret.Data)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		return
	}
	return
}
