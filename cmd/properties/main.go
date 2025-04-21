package main

/*

$> go run -mod readonly cmd/properties/main.go /usr/local/src/placeholder/data/wof.extract
abrv:eng_x_preferred
edtf:deprecated
geom:area
geom:bbox
geom:latitude
geom:longitude
gn:pop
gn:population
iso:country
lbl:bbox
lbl:latitude
lbl:longitude
meso:pop
mz:is_current
ne:iso_a2
ne:iso_a3
ne:pop_est
qs:gn_pop
qs:photo_sum
qs:pop
statoids:population
wk:population
wof:country_alpha3
wof:hierarchy
wof:id
wof:label
wof:megacity
wof:name
wof:parent_id
wof:placetype
wof:population
wof:shortcode
wof:superseded_by
zs:pop10

*/

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/aaronland/go-jsonl/walk"
	"github.com/tidwall/gjson"
)

func main() {

	flag.Parse()

	ctx := context.Background()

	lookup := new(sync.Map)

	for _, path := range flag.Args() {

		r, err := os.Open(path)

		if err != nil {
			log.Fatalf("Failed to open %s for reading, %v", path, err)
		}

		defer r.Close()

		iter_opts := &walk.IterateOptions{}

		for rec, err := range walk.IterateReader(ctx, iter_opts, r) {

			if err != nil {
				log.Fatal(err)
			}

			props, ok := gjson.ParseBytes(rec.Body).Value().(map[string]interface{})

			if !ok {
				log.Fatalf("Failed to parse JSON for record at line %d", rec.LineNumber)
			}

			for k, _ := range props {
				lookup.Store(k, true)
			}
		}
	}

	keys := make([]string, 0)

	lookup.Range(func(k interface{}, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	})

	sort.Strings(keys)

	for _, k := range keys {

		if !strings.HasPrefix(k, "name:") {
			fmt.Println(k)
		}
	}
}
