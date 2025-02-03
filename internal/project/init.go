package project

import (
	"context"
	"fmt"
	"os"
)

func (p *Project) Init(ctx context.Context, forceCreate bool) error {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory")
		panic(err)
	}

	if len(entries) > 0 && !forceCreate {
		fmt.Println("Directory is not empty")
		panic(err)
	}

	return nil
}
