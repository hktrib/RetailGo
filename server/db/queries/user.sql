-- Checkout the Link for specific on sqlc query commenting for generation!
-- https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#:~:text=Next%2C%20create%20a%20query.sql%20file%20with%20the%20following


-- name: CreateUser :one
INSERT into users (
    username, 
    hashed_password, 
    email
) VALUES (@username, @hashed_password, @email) RETURNING *;
