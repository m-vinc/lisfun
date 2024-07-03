package main

import (
	"log"
	"os"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/m-vinc/entx"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: <schema_dir> <target> <package>")
	}

	err := entc.Generate(os.Args[1], &gen.Config{
		Features: []gen.Feature{gen.FeatureLock, gen.FeatureExecQuery, gen.FeatureUpsert},
		Target:   os.Args[2],
		Package:  os.Args[3],
	}, entc.Extensions(
		// entviz.Extension{},
		entx.New(nil),
	))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
