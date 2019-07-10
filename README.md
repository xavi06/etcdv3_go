# etcdv3_go

etcd v3 golang module

## Usage
```bash
go get github.com/xavi06/etcdv3_go
```

```go
import (
    etcdv3 "github.com/xavi06/etcdv3_go"
)

func main() {
    etcdEndpoints := []string{"localhost:2379"}
    cli, err := etcdv3.Conn(etcdEndpoints)
    if err != nil {

    }
    etcdRes, err := etcdv3.Get(cli, key)
    ...

}
```
