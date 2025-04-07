# go-whosonfirst-aceholder

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/extract cmd/extract/main.go
```

### Extract

```
$> ./bin/extract \
	/usr/local/data/whosonfirst-data-admin-us \
	/usr/local/data/whosonfirst-data-venue-us-ca \
	> /usr/local/src/placeholder/data/wof.extract
	
2025/04/07 16:13:34 INFO time to index paths (2) 2m13.55221475s
```

## See also

* https://github.com/pelias/placeholder/