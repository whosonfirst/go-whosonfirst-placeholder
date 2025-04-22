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

```
$> docker run --platform=linux/amd64 placeholder-custom-db \
	/usr/local/bin/build-placeholder-db.sh \
	-T mem:///test.db \
	'sfomuseum-data://?prefix=sfomuseum-data-architecture' \
	'whosonfirst-data://?prefix=whosonfirst-data-admin-us'

/usr/local/bin/wof-extract-properties -iterator-uri org:///tmp?_dedupe=true&_exclude_alt=true&exclude=properties.edtf:deprecated=.* sfomuseum-data://?prefix=sfomuseum-data-architecture whosonfirst-data://?prefix=whosonfirst-data-admin-us > /code/pelias/placeholder/data/wof.extract
2025/04/22 22:16:46 INFO time to index paths (1) 5.11158821s

```