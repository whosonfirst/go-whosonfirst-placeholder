package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst-placeholder/app/extract"
)

func main() {

	ctx := context.Background()
	err := extract.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
