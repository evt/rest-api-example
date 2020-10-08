CREATE TABLE files (
	id uuid NOT NULL,
	filename text NOT NULL,
	created_at timestamp default current_timestamp,
	CONSTRAINT "pk_file_id" PRIMARY KEY (id)
);

