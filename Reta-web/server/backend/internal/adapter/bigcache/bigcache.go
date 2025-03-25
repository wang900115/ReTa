package bigcache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/spf13/viper"
)

// default:
// 		Config{
// 			Shards:             1024,
// 			LifeWindow:         eviction,
// 			CleanWindow:        1 * time.Second,
// 			MaxEntriesInWindow: 1000 * 10 * 60,
// 			MaxEntrySize:       500,
// 			StatsEnabled:       false,
// 			Verbose:            true,
// 			Hasher:             newDefaultHasher(),
// 			HardMaxCacheSize:   0,
// 			Logger:             DefaultLogger(),
// 		}

type Options struct {
	LifeWindow time.Duration
}

func NewBigcacheOptions(conf *viper.Viper) Options {
	return Options{
		LifeWindow: conf.GetDuration("bigcache.life_window"),
	}
}

func NewBigcache(options Options) *bigcache.BigCache {
	config := bigcache.DefaultConfig(options.LifeWindow)
	config.MaxEntrySize = 1024
	bigcache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil
	}
	return bigcache
}

func Set(c *bigcache.BigCache, key string, value interface{}) error {
	valueBytes, err := serialize(value)
	if err != nil {
		return err
	}
	return c.Set(key, valueBytes)
}

func Get(c *bigcache.BigCache, key string) (interface{}, error) {
	valueBytes, err := c.Get(key)
	if err != nil {
		return nil, err
	}
	value, err := deserialize(valueBytes)
	if err != nil {
		return nil, err
	}
	return value, nil
}
