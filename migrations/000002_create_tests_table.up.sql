CREATE TABLE tests (
    id bigint PRIMARY KEY,
    task_id bigint NOT NULL,
    input TEXT NULL,
    output TEXT NULL
);

CREATE SEQUENCE tests_id START 1;

ALTER TABLE tests ALTER COLUMN id SET DEFAULT nextval('tests_id');

INSERT INTO tests (task_id, input, output) VALUES (2, "echo off", "error"), (3,"exec", "true") RETURNING id;
