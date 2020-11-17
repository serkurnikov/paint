-- +goose Up
CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXISTS histories
(
    id SERIAL primary key,
    value double precision default 0.0,
    created_at time with time zone default now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.