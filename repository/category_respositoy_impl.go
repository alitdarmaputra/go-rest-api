package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/alitdarmaputra/belajar-golang-rest-api/helper"
	"github.com/alitdarmaputra/belajar-golang-rest-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(
	ctx context.Context,
	tx *sql.Tx,
	category domain.Category,
) domain.Category {
	SQL := "INSERT INTO app_category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = id
	return category
}

func (repository *CategoryRepositoryImpl) Update(
	ctx context.Context,
	tx *sql.Tx,
	category domain.Category,
) domain.Category {
	SQL := "UPDATE app_category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(
	ctx context.Context,
	tx *sql.Tx,
	categoryId int64,
) {
	SQL := "DELETE FROM app_category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(
	ctx context.Context,
	tx *sql.Tx,
	categoryId int64,
) (domain.Category, error) {
	SQL := "SELECT id, name FROM app_category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(
	ctx context.Context,
	tx *sql.Tx,
) []domain.Category {
	SQL := "SELECT id, name FROM app_category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
