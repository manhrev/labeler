-- name: GetImageToLabel :one
SELECT * FROM images
WHERE url_selected IS NULL LIMIT 1;

-- name: UpdateImageAfterLabeled :one
UPDATE images
  set 
    url_selected = $2,
    labeler_id = $3,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateImageUrlSelected :one
UPDATE images
  set 
    url_selected = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;


-- name: UpdateImageLabelerID :one
UPDATE images
  set 
    labeler_id = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;
