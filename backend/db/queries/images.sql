-- name: GetImageToLabel :one
SELECT * FROM images
WHERE url_selected IS NULL LIMIT 1;

-- name: GetImageByID :one
SELECT * FROM images
WHERE id = $1 AND category = $2 
LIMIT 1;

-- name: UpdateImageAfterLabeled :one
UPDATE images
  set 
    url_selected = $3,
    labeler_id = $4,
    background_type = $5,
    updated_at = now()
WHERE id = $1 AND category = $2
RETURNING *;

-- name: UpdateImageUrlSelectedByLabelerID :one
UPDATE images
  set 
    url_selected = $4,
    updated_at = now()
WHERE id = $1 AND category = $2 AND labeler_id = $3
RETURNING *;


-- name: UpdateImageLabelerID :one
UPDATE images
  set 
    labeler_id = $3,
    updated_at = now()
WHERE id = $1 AND category = $2
RETURNING *;
