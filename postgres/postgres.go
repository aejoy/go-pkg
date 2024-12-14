package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aejoy/xgo/consts"
	"github.com/aejoy/xgo/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Shards []*pgxpool.Pool

func NewPostgres(urls []string) (Shards, error) {
	shards := make(Shards, len(urls))

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

		shards[i] = bucket
	}

	return shards, nil
}

func (shards Shards) GetShard(shardKey string) (*pgxpool.Pool, int, error) {
	shardsCount := len(shards)

	shardIndex, err := utils.GetShardIndex(shardKey, shardsCount)
	if err != nil {
		return nil, shardIndex, err
	}

	if shardIndex <= shardsCount {
		return shards[shardIndex], shardIndex, nil
	}

	return nil, shardIndex, consts.ErrNotFoundBucket
}
