### Run the Micro API

```
$ micro api --handler=http
```

### Run the Service

```
$ go build -o srv srv/*.go
$ ./srv --registry=etcd
```

### Run the API

```
$ go build -o api api/gin/*.go
$ ./api --registry=etcd
```

### Curl the API

Test the index
```
curl http://localhost:8080/greeter
{
  "message": "Hi, this is the Greeter API"
}
```

Test a resource
```
 curl http://localhost:8080/greeter/asim
{
  "msg": "Hello asim"
}
```
