CREATE TABLE users (
	id uuid NOT NULL,
	firstname text NOT NULL,
	lastname text NOT NULL,
	created timestamp default current_timestamp,
	CONSTRAINT "pk_test_id" PRIMARY KEY (id)
);

