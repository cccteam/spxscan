// Package scan allows scanning data into Go structs and other composite types, when working with spanner library native interface.
package spxscan

import (
	"context"

	"cloud.google.com/go/spanner"
)

type Querier interface {
	Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
}

// Get is a package-level helper function that uses the DefaultAPI object.
// See API.Get for details.
func Get(ctx context.Context, db Querier, dst any, statement spanner.Statement) error {
	return DefaultAPI.Get(ctx, db, dst, statement)
}

// Select is a package-level helper function that uses the DefaultAPI object.
// See API.Select for details.
func Select(ctx context.Context, db Querier, dst any, statement spanner.Statement) error {
	return DefaultAPI.Select(ctx, db, dst, statement)
}
