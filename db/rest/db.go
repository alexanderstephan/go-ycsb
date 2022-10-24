package rest

import (
	"context"

	"github.com/magiconair/properties"
)

// properties
const (
	restEndpoints   = "raft.endpoints"
	restDialTimeout = "raft.dial_timeout"
)

type restCreator struct{}

type restDB struct {
	p *properties.Properties
}

func (db *restDB) Close() error {
	return nil
}

func (db *restDB) CleanupThread(ctx context.Context) {

}

func (db *restDB) Read(ctx context.Context, table string, key string, fields []string) (map[string][]byte, error) {
	return nil, nil
}

func (db *restDB) Scan(ctx context.Context, table string, startKey string, count int, fields []string) ([]map[string][]byte, error) {
	return nil, nil
}

func (db *restDB) Update(ctx context.Context, table string, key string, values map[string][]byte) error {
	return nil
}

func (db *restDB) Insert(ctx context.Context, table string, key string, values map[string][]byte) error {
	return nil
}

func (db *restDB) Delete(ctx context.Context, table string, key string) error {
	return nil
}
