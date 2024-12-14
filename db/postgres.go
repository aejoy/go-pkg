package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aejoy/xgo/consts"
	"github.com/aejoy/xgo/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Buckets []*pgxpool.Pool

func NewPostgres(urls []string) (Buckets, error) {
	buckets := make(Buckets, len(urls))

	for i, url := range urls {
		bucket, err := pgxpool.New(context.TODO(), url)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", consts.ErrPostgresConnection, err)
		}

		if err := bucket.Ping(context.TODO()); err != nil {
			return nil, fmt.Errorf("%w: %v", consts.ErrPostgresPing, err)
		}

		db, err := sql.Open("postgres", url)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", consts.ErrSQLOpen, err)
		}

		if err := utils.PostgresMigrate(db); err != nil {
			return nil, fmt.Errorf("%w: %v", consts.ErrMigrate, err)
		}

		buckets[i] = bucket
	}

	return buckets, nil
}

func (buckets Buckets) GetShard(shardKey string) (*pgxpool.Pool, int, error) {
	bucketsCount := len(buckets)

	shardIndex, err := utils.GetShardIndex(shardKey, bucketsCount)
	if err != nil {
		return nil, shardIndex, err
	}

	if shardIndex <= bucketsCount {
		return buckets[shardIndex], shardIndex, nil
	}

	return nil, shardIndex, consts.ErrNotFoundBucket
}
