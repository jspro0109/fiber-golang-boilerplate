-- name: GetSystemStats :one
SELECT
    (SELECT count(*) FROM users WHERE deleted_at IS NULL) AS active_users,
    (SELECT count(*) FROM users WHERE deleted_at IS NOT NULL) AS deleted_users,
    (SELECT count(*) FROM files WHERE deleted_at IS NULL) AS total_files,
    (SELECT COALESCE(SUM(size), 0)::BIGINT FROM files WHERE deleted_at IS NULL) AS total_file_size;
