package database

import (
  "context"
  "gorm.io/gorm"
  "userprofile/application/apperror"
  "userprofile/infrastructure/log"
)

type contextDBType string

var ContextDBValue contextDBType = "DB"

type GormTransactionImpl struct {
  DB *gorm.DB
}

func (r *GormTransactionImpl) BeginTransaction(ctx context.Context) (context.Context, error) {
  log.Info(ctx, "Called")

  dbTrx := r.DB.Begin()

  trxCtx := context.WithValue(ctx, ContextDBValue, dbTrx)

  return trxCtx, nil
}

func (r *GormTransactionImpl) CommitTransaction(ctx context.Context) error {
  log.Info(ctx, "Called")

  db, err := ExtractDB(ctx)
  if err != nil {
    return err
  }

  return db.Commit().Error
}

func (r *GormTransactionImpl) RollbackTransaction(ctx context.Context) error {
  log.Info(ctx, "Called")

  db, err := ExtractDB(ctx)
  if err != nil {
    return err
  }

  return db.Rollback().Error
}

type GormReadOnlyImpl struct {
  DB *gorm.DB
}

func (r *GormReadOnlyImpl) GetDatabase(ctx context.Context) (context.Context, error) {
  log.Info(ctx, "Called")

  trxCtx := context.WithValue(ctx, ContextDBValue, r.DB)

  return trxCtx, nil
}

// ExtractDB is used by other repo to extract the database from context
func ExtractDB(ctx context.Context) (*gorm.DB, error) {

  db, ok := ctx.Value(ContextDBValue).(*gorm.DB)
  if !ok {
    return nil, apperror.DatabaseNotFoundInContextError
  }

  return db, nil
}
