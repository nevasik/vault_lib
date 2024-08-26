package vault

import (
	"context"
	"encoding/json"
)

const servicePath = "kv"

func (c *Client) Service(cfg any, serviceName string) (err error) {
	resp, err := c.C.KVv2(servicePath).Get(context.Background(), serviceName)
	if err != nil {
		return
	}
	data, err := json.Marshal(resp.Data)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, cfg)
	return
}
func (c *Client) WriteService(cfg any, serviceName string) (err error) {
	data, err := json.Marshal(cfg)
	if err != nil {
		return
	}
	var m map[string]any
	if err = json.Unmarshal(data, &m); err != nil {
		return
	}
	_, err = c.C.KVv2(servicePath).Get(context.Background(), serviceName)
	if err != nil {
		_, err = c.C.KVv2(servicePath).Put(context.Background(), serviceName, m)
	} else {
		_, err = c.C.KVv2(servicePath).Patch(context.Background(), serviceName, m)
	}
	return
}
