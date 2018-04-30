package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"strconv"
)

// Client is wrapper for redis.Client
type Client struct {
	Client      *redis.Client
	isConnected bool
}

// Connect to the redis server addr:port with password
func (c *Client) Connect(addr string, port int, password string) error {
	if addr == "" {
		return errors.New("redis.Client.Init: addr cannot be empty")
	}
	if port < 1024 || port > 65534 {
		return errors.New("redis.Client.Init: port must in range [1024, 65534]")
	}
	if c.isConnected {
		c.Client.Close()
		c.isConnected = false
	}
	c.Client = redis.NewClient(&redis.Options{
		Addr:     addr + ":" + strconv.Itoa(port),
		Password: password,
		DB:       0,
	})
	if err := c.Client.Ping().Err(); err != nil {
		return err
	}
	c.isConnected = true
	return nil
}

// IsConnected return connet state of redis client
func (c *Client) IsConnected() bool {
	return c.isConnected
}

// Close redis client
func (c *Client) Close() error {
	if !c.isConnected {
		return nil
	}
	if err := c.Client.Close(); err != nil {
		return err
	}
	c.isConnected = false
	return nil
}
