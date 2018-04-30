package writers_test

import (
	"github.com/newkedison/go-library/redis"
	"github.com/newkedison/go-library/writers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	key  = "test_redis_writer"
	addr = "localhost"
	port = 6379
)

func TestWriteBeforeInit(t *testing.T) {
	assert := assert.New(t)
	w := writers.RedisListWriter()
	n, err := w.Write([]byte{0x00})
	assert.Equal(n, 0)
	assert.NotNil(err)
}

func TestInit(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(writers.InitRedisListWriter(addr, port, "", ""))
	assert.NotNil(writers.InitRedisListWriter(addr, port+1, "", key))
	assert.Nil(writers.InitRedisListWriter(addr, port, "", key))
}

func TestWrite(t *testing.T) {
	assert := assert.New(t)
	w := writers.RedisListWriter()
	n, err := w.Write([]byte{0x00})
	assert.Equal(n, 1)
	assert.Nil(err)
}

func checkListItems(a *assert.Assertions, r *redis.Client, expect ...[]byte) {
	listItems, _ := r.Client.LRange(key, 0, -1).Result()
	a.Equal(len(listItems), len(expect))
	for i := range listItems {
		a.Equal([]byte(listItems[i]), expect[i])
	}
}

func TestLimit(t *testing.T) {
	assert := assert.New(t)
	writers.SetRedisListLimitSize(3)
	var r redis.Client
	assert.Nil(r.Connect(addr, port, ""))
	r.Client.Del(key)
	w := writers.RedisListWriter()
	n, err := w.Write([]byte{0x00})
	assert.Equal(n, 1)
	assert.Nil(err)
	checkListItems(assert, &r, []byte{0x00})
	n, err = w.Write([]byte{0x01})
	assert.Equal(n, 1)
	assert.Nil(err)
	checkListItems(assert, &r, []byte{0x01}, []byte{0x00})
	n, err = w.Write([]byte{0x02})
	assert.Equal(n, 1)
	assert.Nil(err)
	checkListItems(assert, &r, []byte{0x02}, []byte{0x01}, []byte{0x00})
	n, err = w.Write([]byte{0x03})
	assert.Equal(n, 1)
	assert.Nil(err)
	checkListItems(assert, &r, []byte{0x03}, []byte{0x02}, []byte{0x01})
}
