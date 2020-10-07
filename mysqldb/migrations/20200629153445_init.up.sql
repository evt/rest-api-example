CREATE TABLE users (
	id char(36) NOT NULL primary key,
	firstname text NOT NULL,
	lastname text NOT NULL,
	created_at timestamp default current_timestamp
);

