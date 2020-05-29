### Run the Micro API

```
$  micro --registry=etcd web
```

### Run the Service

```
$ go build -o web srv/*.go
$ ./srv --registry=etcd
```

### Run the API

```
$ go build -o web web/*.go
$ ./web --registry=etcd
```

### Curl the API

Test the index
```
http://localhost:8082/greeter
```

