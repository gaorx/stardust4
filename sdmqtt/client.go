package sdmqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gaorx/stardust4/sderr"
)

type Client struct {
	mqtt.Client
}

func Dial(opts *mqtt.ClientOptions) *Client {
	c := mqtt.NewClient(opts)
	return &Client{c}
}

func (c *Client) ConnectSync() error {
	token := c.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		return sderr.Wrap(err, "sdmqtt connect sync error")
	}
	return nil
}

func (c *Client) SubscribeSync(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := c.Subscribe(topic, qos, callback)
	token.Wait()
	if err := token.Error(); err != nil {
		return sderr.Wrap(err, "sdmqtt subscribe sync error")
	}
	return nil
}

func (c *Client) UnsubscribeSync(topics ...string) error {
	token := c.Unsubscribe(topics...)
	token.Wait()
	if err := token.Error(); err != nil {
		return sderr.Wrap(err, "sdmqtt unsubscribe sync error")
	}
	return nil
}

func (c *Client) PublishSync(topic string, qos byte, retained bool, payload any) error {
	token := c.Publish(topic, qos, retained, payload)
	token.Wait()
	if err := token.Error(); err != nil {
		return sderr.Wrap(err, "sdmqtt publish sync error")
	}
	return nil
}
