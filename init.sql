-- --DROP TABLE checks;

CREATE TABLE checks (
    id bigint PRIMARY KEY,
    solution_id bigint NOT NULL,
    test_id bigint NOT NULL,
    runner_id BIGINT NOT NULL,
    success BOOLEAN NOT NULL
);

CREATE SEQUENCE checks_id START 1;

ALTER TABLE checks ALTER COLUMN id SET DEFAULT nextval('checks_id');

INSERT INTO checks (solution_id, test_id, runner_id, success) VALUES (2, 3, 4, FALSE), (3,4,5,TRUE) RETURNING id;

CREATE OR REPLACE FUNCTION random_between(low INT ,high INT) 
   RETURNS INT AS
$$
BEGIN
   RETURN floor(random()* (high-low + 1) + low);
END;
$$ language 'plpgsql' STRICT;

INSERT INTO checks (solution_id, test_id, runner_id, success) 
SELECT random_between(1, 1000000), random_between(1, 1000000), random_between(1, 1000000), CASE WHEN RANDOM() < 0.5 THEN FALSE ELSE TRUE END FROM generate_series(1, 10000);
