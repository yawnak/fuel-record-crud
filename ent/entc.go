//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	features := entc.FeatureNames("privacy")
	if err := entc.Generate("./schema", &gen.Config{}, features); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
