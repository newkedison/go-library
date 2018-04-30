package redis_test

import (
	"github.com/newkedison/go-library/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientConnect(t *testing.T) {
	assert := assert.New(t)
	var r redis.Client
	assert.Nil(r.Close())
	assert.NotNil(r.Connect("", 6379, ""))
	assert.NotNil(r.Connect("localhost", 1023, ""))
	assert.NotNil(r.Connect("localhost", 65535, ""))
	assert.False(r.IsConnected())
	err := r.Connect("localhost", 6379, "")
	assert.True(r.IsConnected())
	assert.Nil(err)
	err = r.Connect("localhost", 6379, "")
	assert.True(r.IsConnected())
	assert.Nil(err)
	assert.Nil(r.Close())
	assert.False(r.IsConnected())
	assert.NotNil(r.Connect("localhost", 6666, ""))
	assert.False(r.IsConnected())
	assert.Nil(r.Close())
}
