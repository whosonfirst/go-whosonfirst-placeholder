package extract

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
)

var iterator_uri string
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("extract")

	fs.StringVar(&iterator_uri, "iterator-uri", "repo://?exclude=properties.edtf:deprecated=.*", "A registered whosonfirst/go-whosonfirst-iterate emitter URI.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Extract generates a Placeholder `wof.extract` (JSONL) file from one or more whosonfirst/go-whosonfirst-iterate sources.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options] uri(N) uri(N)\n", os.Args[0])
		fs.PrintDefaults()
	}

	return fs
}
