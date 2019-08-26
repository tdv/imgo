BEGIN;

CREATE TABLE images
(
	id varchar unique,
	data bytea
);

COMMIT;
