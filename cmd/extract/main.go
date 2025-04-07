package main

import (
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"sync"

	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
	"github.com/whosonfirst/go-whosonfirst-placeholder"
)

func main() {

	var iterator_uri string

	flag.StringVar(&iterator_uri, "iterator-uri", "repo://?exclude=properties.edtf:deprecated=.*", "...")

	flag.Parse()

	iterator_sources := flag.Args()

	ctx := context.Background()

	wr := os.Stdout
	mu := new(sync.RWMutex)

	iter_cb := func(ctx context.Context, path string, r io.ReadSeeker, args ...interface{}) error {

		body, err := io.ReadAll(r)

		if err != nil {
			return err
		}

		doc := make(map[string]interface{})

		props_rsp := gjson.GetBytes(body, "properties")

		for k, v := range props_rsp.Map() {

			if !placeholder.IsExtractProperty(k) {
				continue
			}

			doc[k] = v.Value()
		}

		mu.Lock()
		defer mu.Unlock()

		enc := json.NewEncoder(wr)
		err = enc.Encode(doc)

		if err != nil {
			return err
		}

		return nil
	}

	iter, err := iterator.NewIterator(ctx, iterator_uri, iter_cb)

	if err != nil {
		log.Fatal(err)
	}

	err = iter.IterateURIs(ctx, iterator_sources...)

	if err != nil {
		log.Fatal(err)
	}
}
