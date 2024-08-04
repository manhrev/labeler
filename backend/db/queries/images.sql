-- name: GetImageToLabel :one
SELECT * FROM images
WHERE url_selected IS NULL AND (labeler_id IS NULL OR labeler_id = $1) LIMIT 1;

-- name: UpdateImageLabelerID :exec
UPDATE images
  set 
    labeler_id = $3,
    updated_at = now()
WHERE id = $1 AND category = $2;

-- name: GetImageByID :one
SELECT * FROM images
WHERE id = $1 AND category = $2 
LIMIT 1;

-- name: UpdateImageAfterLabeled :exec
UPDATE images
  set 
    url_selected = $3,
    labeler_id = $4,
    background_type = $5,
    updated_at = now()
WHERE id = $1 AND category = $2;

-- name: UpdateImageUrlSelectedByLabelerID :exec
UPDATE images
  set 
    url_selected = $4,
    updated_at = now()
WHERE id = $1 AND category = $2 AND labeler_id = $3;

-- name: RollbackImageLabeled :exec
UPDATE images
  set 
    url_selected = NULL,
    updated_at = now(),
    background_type = NULL
WHERE id = $1 AND category = $2 AND labeler_id = $3;

-- name: CountImagesByLabelerID :one
SELECT COUNT(*) FROM images
WHERE labeler_id = $1;

-- name: FindImagesByFilters :many
SELECT * FROM images
WHERE 
  (@category::category IS NULL OR category = @category) AND
  (@labeler_id::BIGINT = 0 OR labeler_id = @labeler_id)
ORDER BY updated_at DESC
LIMIT @lim::BIGINT OFFSET @off::BIGINT;

-- name: CountImagesByFilters :one
SELECT COUNT(*) FROM images
WHERE 
  (@category::category IS NULL OR category = @category) AND
  (@labeler_id::BIGINT = 0 OR labeler_id = @labeler_id);
