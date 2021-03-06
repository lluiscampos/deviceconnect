// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package clientnats

import (
	"context"

	"github.com/nats-io/nats.go"
)

// ClientInterface is the interface of a NATS client
type ClientInterface interface {
	Connect(ctx context.Context, uri string) error
	Subscribe(subject string, callback nats.MsgHandler) error
	Publish(subject string, data []byte) error
}

// Client provides a NATS client
type Client struct {
	conn *nats.Conn
}

// NewClient returns a new Client
func NewClient() *Client {
	return new(Client)
}

// Connect connects to a NATS instance
func (c *Client) Connect(ctx context.Context, uri string) error {
	nc, err := nats.Connect(uri)
	if err != nil {
		return err
	}
	c.conn = nc
	return nil
}

// Subscribe to a subject
func (c *Client) Subscribe(subject string, callback nats.MsgHandler) error {
	_, err := c.conn.Subscribe(subject, callback)
	return err
}

// Publish to a subject
func (c *Client) Publish(subject string, data []byte) error {
	return c.conn.Publish(subject, data)
}
