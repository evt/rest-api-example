CREATE TABLE users (
	id binary(16) NOT NULL primary key,
	firstname text NOT NULL,
	lastname text NOT NULL,
	created timestamp default current_timestamp
);

