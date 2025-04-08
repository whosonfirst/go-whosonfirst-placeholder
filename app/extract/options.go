package extract

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	IteratorURI     string
	IteratorSources []string
	Verbose         bool
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "WHOSONFIRST")

	if err != nil {
		return nil, fmt.Errorf("Failed to assign flags from environment variables, %w", err)
	}

	opts := &RunOptions{
		IteratorURI:     iterator_uri,
		IteratorSources: fs.Args(),
		Verbose:         verbose,
	}

	return opts, nil
}
