# docker

## Building

The easiest way to build the `placeholder-custom-db` container is to run the handy `docker` Makefile target. For example:

```
$> make docker
docker buildx build --platform=linux/amd64 --no-cache=true -f Dockerfile -t placeholder-custom-db .

...docker stuff happens                                                                                                                                                                         
View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/z4gvob24zf7kse7ufadivx67d
```

## Running

Run the `/usr/local/bin/build-placeholder-db.sh` script passing in a `-T` (target) flag and one or more [whosonfirst/go-whosonfirst-iterate-organization](https://github.com/whosonfirst/go-whosonfirst-iterate-organization) style iterator sources to use as input for your custom Placeholder database.

For example, to build a custom Placeholder database containing records source from the [sfomuseum-data/sfomuseum-data-architecture](https://github.com/sfomuseum-data/sfomuseum-data-architecture) and [whosonfirst-data/whosonfirst-data-admin-us](https://github.com/whosonfirst-data/whosonfirst-data-admin-us) repositories you would do this:

```
$> docker run --platform=linux/amd64 placeholder-custom-db \
	/usr/local/bin/build-placeholder-db.sh \
	-T mem:///test.db \
	'sfomuseum-data://?prefix=sfomuseum-data-architecture' \
	'whosonfirst-data://?prefix=whosonfirst-data-admin-us'

/usr/local/bin/wof-extract-properties -iterator-uri org:///tmp?dedupe=true&_exclude_alt=true&exclude=properties.edtf:deprecated=.* sfomuseum-data://?prefix=sfomuseum-data-architecture whosonfirst-data://?prefix=whosonfirst-data-admin-us > /code/pelias/placeholder/data/wof.extract
2025/04/22 22:16:46 INFO time to index paths (1) 5.11158821s
2025/04/22 22:26:28 INFO time to index paths (1) 9m40.539258724s
2025/04/22 22:26:28 INFO time to index paths (2) 9m48.501629602s

build database

> pelias-placeholder@0.0.0-development build
> bash ./cmd/build.sh

import...
populate fts...
optimize...
close...
Done!

/usr/local/bin/copy -source file:///code/pelias/placeholder/data/store.sqlite3 -target mem:///test.db
```

### Targets

Target URIs (the place where the custom database will be copied to) take the form of a fully-qualified [gocloud.dev/blob](https://pkg.go.dev/gocloud.dev/blob) URI. 

Under the hood the [aaronland/gocloud-blob/cmd/copy](https://github.com/aaronland/gocloud-blob) tool is being used to copy the custom database created in the container to "somewhere else". In the example above the custom database is copied from disk to memory.

The details of "somewhere else" are defined by the target URI. The following `gocloud.dev/blob` providers are supported by default:

* [mem://](https://gocloud.dev/howto/blob/#local)
* [file://](https://gocloud.dev/howto/blob/#local)
* [s3://](https://gocloud.dev/howto/blob/#s3)
* [s3blob://](https://github.com/aaronland/gocloud-blob/tree/main/s3)

## Dockerfile.sfomuseum

There is also a [Dockerfile.sfomuseum](Dockerfile.sfomuseum) build file (and a corresponding `docker-sfomuseum` Makefile target) which can be used to create a container which uses the `sfomuseum-all` branch of the [sfomuseum/placeholder](https://github.com/sfomuseum/placeholder/tree/sfomuseum-all) package. This package introduces two important changes from the main `pelias/placeholder` codebase:

1. It enables indexing of Who's On First (WOF) records regardless of their `mz:is_current` flag (but which still have NOT been marked as deprecated).
2. It enables indexing of WOF records which have been superseded.

These changes are necessary in order to build a Placeholder (SQLite) database with support for historical records.

## See also

* https://github.com/whosonfirst/go-whosonfirst-iterate-organization
* https://github.com/aaronland/gocloud-blob
* https://pkg.go.dev/gocloud.dev/blob