CREATE TABLE files (
	id char(36) NOT NULL PRIMARY KEY,
	filename text NOT NULL,
	created_at timestamp default current_timestamp
);

