# go-whosonfirst-placeholder

## Important

Work in progress. Documentation is incomplete. It may have bugs.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/extract cmd/extract/main.go
```

### Extract

Extract generates a Placeholder `wof.extract` (JSONL) file from one or more [whosonfirst/go-whosonfirst-iterate](https://github.com/whosonfirst/go-whosonfirst-iterate) sources.

For example, to generate source data for a Placeholder instance for geocoding administrative and venue records in the US:

```
$> ./bin/extract \
	/usr/local/data/whosonfirst-data-admin-us \
	/usr/local/data/whosonfirst-data-venue-us-ca \
	> /usr/local/src/placeholder/data/wof.extract
	
2025/04/07 16:13:34 INFO time to index paths (2) 2m13.55221475s
```

Now, go back to the [placeholder](https://github.com/pelias/placeholder) and run the `build` tool:

```
$> /usr/local/src/placeholder
$> npm run build
...time passes

populate fts...
optimize...
close...
Done!
```

_I guess this could be automated too but it hasn't been yet._

And then, query for venues and neighbourhoods:

```
$> npm run cli -- 'Latin American Club'

> pelias-placeholder@0.0.0-development cli
> node cmd/cli.js Latin American Club

Latin American Club

took: 1.21ms
 - 571986789	venue 	Latin American Club

> npm run cli -- 'Gowanus Heights'

> pelias-placeholder@0.0.0-development cli
> node cmd/cli.js Gowanus Heights

Gowanus Heights

took: 1.16ms
 - 102061079	neighbourhood 	Gowanus Heights
```

Or, to generate source data for a Placeholder instance for geocoding administrative and venue records in the San Francisco, San Mateo and Alameda counties:

```
$> ./bin/extract \
	-iterator-uri 'repo://?exclude=properties.edtf:deprecated=.*&include=properties.wof:belongsto=(102087579|102085387|102086959)' \
	/usr/local/data/whosonfirst-data-admin-us \
	/usr/local/data/whosonfirst-data-venue-us-ca \
	> /usr/local/src/placeholder/data/wof.extract
	
2025/04/07 17:10:05 INFO time to index paths (2) 2m5.602572958s
```

## See also

* https://github.com/pelias/placeholder/
* https://github.com/whosonfirst/go-whosonfirst-iterate