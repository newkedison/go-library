package writers

import (
	"errors"
	"io"
	"sync"

	"github.com/newkedison/go-library/redis"
)

type redisListWriter struct {
	key       string
	limitSize int64
	mutex     sync.Mutex
	client    redis.Client
}

func (r *redisListWriter) Write(data []byte) (n int, err error) {
	if r.client.IsConnected() == false {
		return 0, errors.New("redisListWriter not connect to redis, you must call" +
			"writers.InitRedisListWriter first")
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	n = 0
	err = r.client.Client.LPush(r.key, data).Err()
	if err != nil {
		return
	}
	if r.limitSize > 0 {
		err = r.client.Client.LTrim(r.key, 0, r.limitSize-1).Err()
		if err != nil {
			return
		}
	}
	return len(data), nil
}

var writer redisListWriter

// RedisListWriter return a io.Writer that can write to redis, you must call
// writers.InitRedisListWriter before any write operator
func RedisListWriter() io.Writer {
	return &writer
}

// InitRedisListWriter init the redis writer
func InitRedisListWriter(
	addr string, port int, password string, key string) error {
	if key == "" {
		return errors.New("redisListWriter.Init: key cannot be empty")
	}
	r := &writer
	r.mutex.Lock()
	defer r.mutex.Unlock()
	err := r.client.Connect(addr, port, password)
	if err != nil {
		return err
	}
	r.key = key
	r.limitSize = 0
	return nil
}

func SetRedisListLimitSize(size int64) {
	writer.limitSize = size
}
