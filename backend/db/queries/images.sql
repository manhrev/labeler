-- name: GetImageToLabel :one
SELECT * FROM images
WHERE url_selected IS NULL AND (labeler_id IS NULL OR labeler_id = $1) LIMIT 1;

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


-- name: UpdateImageLabelerID :exec
UPDATE images
  set 
    labeler_id = $3,
    updated_at = now()
WHERE id = $1 AND category = $2;
