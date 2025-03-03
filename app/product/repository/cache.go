package repository

import (
	"encoding/json"
	"fmt"
	"time"
)

type CachedProductQuery struct {
	prefix       string
	productQuery ProductQuery
}

func NewCachedProductQuery(productQuery ProductQuery) CachedProductQuery {
	return CachedProductQuery{
		prefix:       "gomall",
		productQuery: productQuery,
	}
}

func (q CachedProductQuery) GetProductById(id uint32) (product Product, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d", q.prefix, "product", id)
	// get data from cache
	cachedData, err := rdb.Get(q.productQuery.ctx, cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &product)
		if err == nil {
			return product, nil
		}
		// unmarshal failed
	}
	// get data from database
	product, err = q.productQuery.GetProductById(id)
	if err != nil {
		return
	}
	// marshal data
	data, err := json.Marshal(product)
	if err == nil {
		// set data to cache
		rdb.Set(q.productQuery.ctx, cacheKey, data, time.Hour)
	}
	return
}
