package utils

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

func init() {
	Cache = cache.New(time.Hour*24, time.Minute*1)
}
