package service

import (
	"encoding/json"
	"hsr-profile-tracker/internal/database"
	"hsr-profile-tracker/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

func CacheGetBytes(uid string) ([]byte, error) {
	if database.Rdb == nil {
		return nil, redis.Nil
	}
	return database.Rdb.Get(database.Ctx, uid).Bytes()
}

func CacheSetBytes(uid string, b []byte) {
	if database.Rdb == nil || len(b) == 0 {
		return
	}
	_ = database.Rdb.Set(database.Ctx, uid, b, time.Hour).Err()
}

func CacheGetSummary(uid string) (*model.ProfileSummary, bool) {
	cachedBytes, err := CacheGetBytes(uid)
	if err == nil && len(cachedBytes) > 0 {
		var cachedSummary model.ProfileSummary
		if json.Unmarshal(cachedBytes, &cachedSummary) == nil {
			return &cachedSummary, true
		}
	}
	return nil, false
}
