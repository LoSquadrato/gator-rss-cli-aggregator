-- name: GetFeedFromUrl :one
SELECT * FROM feeds WHERE url = $1;