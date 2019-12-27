//go:generate  go run github.com/99designs/gqlgen
package graphql

import "github.com/deni1688/bookingql/loaders"

type Resolver struct{ *loaders.Loaders }
