-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    username,
    email,
    secret_code
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE verify_emails
SET
    is_used = TRUE
WHERE
    id = @id
    AND secret_code = LEFT(@secret_code, LENGTH(@secret_code) - 1)
    AND is_used = FALSE
    AND expired_at > now()
RETURNING *;