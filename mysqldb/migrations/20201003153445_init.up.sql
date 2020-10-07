CREATE TABLE files (
	id binary(16) NOT NULL PRIMARY KEY,
	filename text NOT NULL,
	created timestamp default current_timestamp
);

