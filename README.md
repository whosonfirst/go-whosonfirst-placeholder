# go-whosonfirst-placeholder

Go package for building custom [pelias/placeholder](https://github.com/pelias/placeholder/) SQLite databases from one or more Who's On First (style) data sources.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/extract cmd/extract/main.go
```

### Extract

Extract generates a Placeholder `wof.extract` (JSONL) file from one or more [whosonfirst/go-whosonfirst-iterate](https://github.com/whosonfirst/go-whosonfirst-iterate) sources.

```
$> ./bin/extract -h
Extract generates a Placeholder `wof.extract` (JSONL) file from one or more whosonfirst/go-whosonfirst-iterate sources.
Usage:
	./bin/extract [options] uri(N) uri(N)
  -access-token string
    	An optional GitHub API access token. If not-empty it will be used to replace any instances of the string "{access_token}" in any iteraror source URIs
  -iterator-uri string
    	A registered whosonfirst/go-whosonfirst-iterate emitter URI. (default "repo://?exclude=properties.edtf:deprecated=.*")
  -verbose
    	Enable verbose (debug) logging.
```

#### Example

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

#### Iterator sources

The following iterator (emitter, actually) sources are enabled by default:

* All of those provided by the [whosonfirst/go-whosonfirst-iterate](https://github.com/whosonfirst/go-whosonfirst-iterate?tab=readme-ov-file#uris-and-schemes-for-emitters) package.
* [whosonfirst/go-whosonfirst-iterate-git](https://github.com/whosonfirst/go-whosonfirst-iterate-git), for iterating documents in one or more Git repositories
* [whosonfirst/go-whosonfirst-iterate-organization](https://github.com/whosonfirst/go-whosonfirst-iterate-organization), for iterating documents in a set of (Git) repositories in one or more GitHub organizations.

In order to enable other iterators (emitters) you will need to clone the [cmd/extract/main.go](cmd/extract/main.go) tool and add the relevant import statement. For example this is how you would enable support for the [whosonfirst/go-whosonfirst-iterate-bucket](https://github.com/whosonfirst/go-whosonfirst-iterate-bucket) package:

```
package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst-iterate-bucket"
	
	"github.com/whosonfirst/go-whosonfirst-placeholder/app/extract"
)

func main() {

	ctx := context.Background()
	err := extract.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
```

## Docker

Consult [docker/README.md](docker/README.md) for details.

## See also

* https://github.com/pelias/placeholder/
* https://github.com/whosonfirst/go-whosonfirst-iterate
* https://github.com/whosonfirst/go-whosonfirst-iterate-git
* https://github.com/whosonfirst/go-whosonfirst-iterate-organization