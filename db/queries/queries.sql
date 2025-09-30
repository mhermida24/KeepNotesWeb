-- name: GetFolder :one
SELECT id, name, description, parent_folder_id, created_at
FROM folders
WHERE id = $1;

-- name: GetNote :one
SELECT id, title, body, created_at
FROM notes
WHERE id = $1;

-- name: ListFolders :many
SELECT id, name, description, parent_folder_id, created_at
FROM folders
ORDER BY name;

-- name: ListNotes :many
SELECT id, title, body, created_at
FROM notes
ORDER BY title;

-- name: CreateFolder :one
INSERT INTO folders (name, description, parent_folder_id)
VALUES ($1, $2, $3)
RETURNING id, name, description, parent_folder_id, created_at;

-- name: CreateNote :one
INSERT INTO notes (title, body, folder_id)
VALUES ($1, $2, $3)
RETURNING id, title, body, folder_id, created_at;

-- name: UpdateFolder :exec
UPDATE folders
SET name = $2, description = $3, parent_folder_id = $4
WHERE id = $1;

-- name: UpdateNote :exec
UPDATE notes
SET title = $2, body = $3, folder_id = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteFolder :exec
DELETE FROM folders
WHERE id = $1;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;