package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/magiconair/properties"
	"github.com/pingcap/go-ycsb/pkg/ycsb"
)

// properties
const (
	restEndpoint    = "rest.endpoint"
	restDialTimeout = "rest.dial_timeout"
)

var url string

type restCreator struct{}

type restDB struct {
	p *properties.Properties
}

func init() {
	ycsb.RegisterDBCreator("rest", restCreator{})
}

func (c restCreator) Create(p *properties.Properties) (ycsb.DB, error) {
	url = p.GetString(restEndpoint, "localhost:2379") + "/"
	fmt.Println(url)

	return &restDB{
		p: p,
	}, nil
}
func (db *restDB) Close() error {
	return nil
}

func (db *restDB) CleanupThread(ctx context.Context) {

}

func (db *restDB) InitThread(ctx context.Context, _ int, _ int) context.Context {
	return ctx
}

func getURL(key string) string {
	var sb strings.Builder
	sb.WriteString(url)
	sb.WriteString(key)
	return sb.String()
}

func (db *restDB) Read(ctx context.Context, table string, key string, _ []string) (map[string][]byte, error) {
	res, err := http.Get(getURL(key))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println(body)
		return nil, errors.New("wrong status code")
	}
	return map[string][]byte{
		key: body,
	}, nil
}

func (db *restDB) Scan(ctx context.Context, _ string, startKey string, count int, fields []string) ([]map[string][]byte, error) {
	// TODO
	return nil, nil
}

func (db *restDB) Update(ctx context.Context, table string, key string, values map[string][]byte) error {
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, getURL(key), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
		return errors.New("wrong status code")
	}
	return nil
}

func (db *restDB) Insert(ctx context.Context, table string, key string, values map[string][]byte) error {
	return db.Update(ctx, table, key, values)
}

func (db *restDB) Delete(ctx context.Context, table string, key string) error {
	req, err := http.NewRequest(http.MethodDelete, getURL(key), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("wrong status code")
	}
	return nil
}
