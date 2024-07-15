// Package scan allows scanning data into Go structs and other composite types, when working with spanner library native interface.
package spxscan

import (
	"context"
	"time"

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

type TxnRunner interface {
	ReadWriteTransaction(ctx context.Context, f func(context.Context, *spanner.ReadWriteTransaction) error) (commitTimestamp time.Time, err error)
}

// UpdateAndGet is a package-level helper function that uses the DefaultAPI object.
// See API.UpdateAndGet for details.
func UpdateAndGet(ctx context.Context, client TxnRunner, dst any, statement spanner.Statement) error {
	return DefaultAPI.UpdateAndGet(ctx, client, dst, statement)
}

// UpdateAndSelect is a package-level helper function that uses the DefaultAPI object.
// See API.UpdateAndSelect for details.
func UpdateAndSelect(ctx context.Context, client TxnRunner, dst any, statement spanner.Statement) error {
	return DefaultAPI.UpdateAndSelect(ctx, client, dst, statement)
}
