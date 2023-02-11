package repository

import (
	"context"
	"database/sql"

	"github.com/alitdarmaputra/belajar-golang-rest-api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int64)
	FindById(ctx context.Context, tx *sql.Tx, cateogryId int64) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
