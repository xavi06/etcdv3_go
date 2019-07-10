package etcdv3

import (
	"context"
	"errors"
	"strings"
	"time"

	//"github.com/coreos/etcd/clientv3"
    "go.etcd.io/etcd/clientv3"
)

var requestTimeout = 30 * time.Second

// Conn func
func Conn(endpoints []string) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		return cli, err
	}
	//defer cli.Close()
	return cli, err
}

// Get func
func Get(cli *clientv3.Client, path string) (Path, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, path)
	cancel()
	if err != nil {
		return Path{}, err
	}
	if len(resp.Kvs) == 1 {
		return Path{string(resp.Kvs[0].Key), string(resp.Kvs[0].Value)}, nil

	}
	//fmt.Println(len(resp.Kvs))
	return Path{}, errors.New("error")

}

// GetPrefixName func
func GetPrefixName(cli *clientv3.Client, prefix string) []Key {
	var data []Key
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	//resp, err := cli.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	resp, err := cli.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	//resp, err := cli.Get(ctx, "/nginx")
	cancel()
	if err != nil {
		return data
	}
	ids := make(map[string]string)
	for _, ev := range resp.Kvs {
		keyList := strings.Split(string(ev.Key), "/")
		for k := range keyList {
			i := k + 1
			if i >= len(keyList) {
				continue
			}
			id := strings.Join(keyList[0:i+1], "/")
			if ids[id] != "" {
				continue
			} else {
				var parent string
				if i == 1 {
					parent = "#"
				} else {
					parent = strings.Join(keyList[0:i], "/")
				}
				text := keyList[i]
				ids[id] = id
				data = append(data, Key{id, parent, text})
			}

		}
	}
	return data

}

// Put func
func Put(cli *clientv3.Client, key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		return err
	}
	return nil
}

// Delete func
func Delete(cli *clientv3.Client, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	// delete the keys
	_, err := cli.Delete(ctx, key)
	return err
}
