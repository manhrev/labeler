// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: images.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getImageByID = `-- name: GetImageByID :one
SELECT id, category, background_type, labeler_id, name, display_name, url1, url2, url3, url_selected, created_at, updated_at FROM images
WHERE id = $1 AND category = $2 
LIMIT 1
`

type GetImageByIDParams struct {
	ID       int64
	Category Category
}

func (q *Queries) GetImageByID(ctx context.Context, arg GetImageByIDParams) (Image, error) {
	row := q.db.QueryRow(ctx, getImageByID, arg.ID, arg.Category)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.Category,
		&i.BackgroundType,
		&i.LabelerID,
		&i.Name,
		&i.DisplayName,
		&i.Url1,
		&i.Url2,
		&i.Url3,
		&i.UrlSelected,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getImageToLabel = `-- name: GetImageToLabel :one
SELECT id, category, background_type, labeler_id, name, display_name, url1, url2, url3, url_selected, created_at, updated_at FROM images
WHERE url_selected IS NULL LIMIT 1
`

func (q *Queries) GetImageToLabel(ctx context.Context) (Image, error) {
	row := q.db.QueryRow(ctx, getImageToLabel)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.Category,
		&i.BackgroundType,
		&i.LabelerID,
		&i.Name,
		&i.DisplayName,
		&i.Url1,
		&i.Url2,
		&i.Url3,
		&i.UrlSelected,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const rollbackImageLabeled = `-- name: RollbackImageLabeled :exec
UPDATE images
  set 
    url_selected = NULL,
    updated_at = now(),
    background_type = NULL
WHERE id = $1 AND category = $2 AND labeler_id = $3
`

type RollbackImageLabeledParams struct {
	ID        int64
	Category  Category
	LabelerID pgtype.Int8
}

func (q *Queries) RollbackImageLabeled(ctx context.Context, arg RollbackImageLabeledParams) error {
	_, err := q.db.Exec(ctx, rollbackImageLabeled, arg.ID, arg.Category, arg.LabelerID)
	return err
}

const updateImageAfterLabeled = `-- name: UpdateImageAfterLabeled :exec
UPDATE images
  set 
    url_selected = $3,
    labeler_id = $4,
    background_type = $5,
    updated_at = now()
WHERE id = $1 AND category = $2
`

type UpdateImageAfterLabeledParams struct {
	ID             int64
	Category       Category
	UrlSelected    pgtype.Int2
	LabelerID      pgtype.Int8
	BackgroundType NullBackgroundType
}

func (q *Queries) UpdateImageAfterLabeled(ctx context.Context, arg UpdateImageAfterLabeledParams) error {
	_, err := q.db.Exec(ctx, updateImageAfterLabeled,
		arg.ID,
		arg.Category,
		arg.UrlSelected,
		arg.LabelerID,
		arg.BackgroundType,
	)
	return err
}

const updateImageLabelerID = `-- name: UpdateImageLabelerID :exec
UPDATE images
  set 
    labeler_id = $3,
    updated_at = now()
WHERE id = $1 AND category = $2
`

type UpdateImageLabelerIDParams struct {
	ID        int64
	Category  Category
	LabelerID pgtype.Int8
}

func (q *Queries) UpdateImageLabelerID(ctx context.Context, arg UpdateImageLabelerIDParams) error {
	_, err := q.db.Exec(ctx, updateImageLabelerID, arg.ID, arg.Category, arg.LabelerID)
	return err
}

const updateImageUrlSelectedByLabelerID = `-- name: UpdateImageUrlSelectedByLabelerID :exec
UPDATE images
  set 
    url_selected = $4,
    updated_at = now()
WHERE id = $1 AND category = $2 AND labeler_id = $3
`

type UpdateImageUrlSelectedByLabelerIDParams struct {
	ID          int64
	Category    Category
	LabelerID   pgtype.Int8
	UrlSelected pgtype.Int2
}

func (q *Queries) UpdateImageUrlSelectedByLabelerID(ctx context.Context, arg UpdateImageUrlSelectedByLabelerIDParams) error {
	_, err := q.db.Exec(ctx, updateImageUrlSelectedByLabelerID,
		arg.ID,
		arg.Category,
		arg.LabelerID,
		arg.UrlSelected,
	)
	return err
}
