package service

import (
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-cache/cache"
)

type ICache interface {
	cache.IGCache
}

type cacheImpl struct {
	*cache.GfCache
	prefix string
}

var (
	c              = cacheImpl{}
	cacheContainer *cache.GfCache
	lock           = &sync.Mutex{}
)

func Cache() ICache {
	var (
		ch  = c
		ctx = gctx.New()
	)
	prefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
	// model := g.Cfg().MustGet(ctx, "system.cache.model").String()
	if cacheContainer == nil {
		lock.Lock()
		if cacheContainer == nil {
			{
				// memory
				cacheContainer = cache.New(prefix)
			}
		}
		lock.Unlock()
	}
	ch.GfCache = cacheContainer
	return &ch
}
