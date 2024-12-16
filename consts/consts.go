package consts

import "errors"

const (
	NanoAlphabet = "2346789abcdefghijkmnpqrtwxyzABCDEFGHJKLMNPQRTUVWXYZ"
	NanoLength   = 25
)

var (
	ErrNotFoundBucket     = errors.New("bucket is null")
	ErrPostgresConnection = errors.New("postgres connect error")
	ErrPostgresPing       = errors.New("postgres ping error")
	ErrSQLOpen            = errors.New("sql open error")
	ErrMigrate            = errors.New("migrate error")
	ErrMigrateInstance    = errors.New("migrate with instance sql error")
	ErrMigrateOpenFile    = errors.New("open migrations file error")
	ErrMigrateNewInstance = errors.New("new migration instance error")
	ErrMigrateUp          = errors.New("up migrate error")
	ErrOverflowOccurred   = errors.New("overflow occurred")
)
