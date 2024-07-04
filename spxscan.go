// Package scan allows scanning data into Go structs and other composite types, when working with spanner library native interface.
package spxscan

import (
	"context"
	"fmt"
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

// UpdateAndGet is a package-level helper function that uses the DefaultAPI.Get() call inside a transaction.
// This should be used when data is being returned after an update using the THEN RETURN clause.
func UpdateAndGet(ctx context.Context, client TxnRunner, dst any, statement spanner.Statement) error {
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		if err := DefaultAPI.Get(ctx, txn, dst, statement); err != nil {
			return fmt.Errorf("spxscan.Get(): %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("ReadWriteTransaction(): %w", err)
	}

	return nil
}

// UpdateAndSelect is a package-level helper function that uses the DefaultAPI.Select() call inside a transaction.
// This should be used when data is being returned after an update using the THEN RETURN clause.
func UpdateAndSelect(ctx context.Context, client TxnRunner, dst any, statement spanner.Statement) error {
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		if err := DefaultAPI.Select(ctx, txn, dst, statement); err != nil {
			return fmt.Errorf("spxscan.Get(): %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("ReadWriteTransaction(): %w", err)
	}

	return nil
}
