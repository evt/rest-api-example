CREATE TABLE users (
	id uuid NOT NULL,
	firstname text NOT NULL,
	lastname text NOT NULL,
	created_at timestamp default current_timestamp,
	CONSTRAINT "pk_user_id" PRIMARY KEY (id)
);

