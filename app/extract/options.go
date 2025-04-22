package extract

import (
	"flag"
	"fmt"
	"strings"
	
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

	sources := fs.Args()

	if access_token != "" {

		for i, src := range sources {

			if ! strings.Contains(src, "{access_token}"){
				continue
			}

			sources[i] = strings.Replace(src, "{access_token}", access_token, 1)
		}
	}
	
	opts := &RunOptions{
		IteratorURI:     iterator_uri,
		IteratorSources: fs.Args(),
		Verbose:         verbose,
	}

	return opts, nil
}
