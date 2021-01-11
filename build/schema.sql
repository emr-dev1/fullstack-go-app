CREATE DATABASE fullstack_db;

\c fullstack_db;

CREATE TABLE IF NOT EXISTS user (
	id SERIAL PRIMARY KEY,
	nickname VARCHAR(255) NOT NULL UNIQUE,
	email VARCHAR(100) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL,
	created_at timestamp,
	updated_at timestamp
);

CREATE TABLE IF NOT EXISTS post (
	id SERIAL PRIMARY KEY,
	title VARCHAR,
	content TEXT,
	author_id INTEGER NOT NULL,
	created_at timestamp,
	updated_at timestamp
);
