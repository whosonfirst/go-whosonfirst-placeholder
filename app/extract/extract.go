package extract

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"

	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
	"github.com/whosonfirst/go-whosonfirst-placeholder"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := RunOptionsFromFlagSet(fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

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

	iter, err := iterator.NewIterator(ctx, opts.IteratorURI, iter_cb)

	if err != nil {
		return fmt.Errorf("Failed to create new iterator, %w", err)
	}

	err = iter.IterateURIs(ctx, opts.IteratorSources...)

	if err != nil {
		return fmt.Errorf("Failed to iterate sources, %w", err)
	}

	return nil
}
