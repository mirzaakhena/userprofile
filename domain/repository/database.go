package repository

import "context"

// ReadOnlyDB used to get database object from any database implementation.
// For consistency reason both TransactionDB and ReadOnlyDB will seek database object under the context params
type ReadOnlyDB interface {
  GetDatabase(ctx context.Context) (context.Context, error)
}

// TransactionDB used for common transaction handling
// all the context must use the same database session.
type TransactionDB interface {
  BeginTransaction(ctx context.Context) (context.Context, error)
  CommitTransaction(ctx context.Context) error
  RollbackTransaction(ctx context.Context) error
}

// ReadOnly is helper function that simplify the readonly db
func ReadOnly(ctx context.Context, trx ReadOnlyDB, trxFunc func(dbCtx context.Context) error) error {
  dbCtx, err := trx.GetDatabase(ctx)
  if err != nil {
    return err
  }
  return trxFunc(dbCtx)
}

// WithTransaction is helper function that simplify the transaction execution handling
func WithTransaction(ctx context.Context, trx TransactionDB, trxFunc func(dbCtx context.Context) error) error {
  dbCtx, err := trx.BeginTransaction(ctx)
  if err != nil {
    return err
  }

  defer func() {
    if p := recover(); p != nil {
      err = trx.RollbackTransaction(dbCtx)
      panic(p)

    } else if err != nil {
      err = trx.RollbackTransaction(dbCtx)

    } else {
      err = trx.CommitTransaction(dbCtx)

    }
  }()

  err = trxFunc(dbCtx)
  return err
}
