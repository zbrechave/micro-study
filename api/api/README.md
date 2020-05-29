### Run the Micro API

```
$ micro api --handler=api
```

### Run the Service

```
$ go build -o srv srv/*.go
$ ./srv --registry=etcd
```

### Run the API

```
$ go build -o api api/api/*.go
$ ./api --registry=etcd
```

### Curl the API

Test the index
```
curl http://localhost:8080/greeter/say/hello?name=John
{
  "message": "Hello John"
}
```

