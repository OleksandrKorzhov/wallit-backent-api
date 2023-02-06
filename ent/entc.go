//go:build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaPath("./api/graphql/ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("create entgql extension: %v", err)
	}

	options := []entc.Option{
		entc.Extensions(ex),
	}
	if err := entc.Generate("./ent/schema", &gen.Config{}, options...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
