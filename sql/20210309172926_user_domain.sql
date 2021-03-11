-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
	id INT auto_increment NOT NULL,
	email VARCHAR(60) NOT NULL,
	phone_number VARCHAR(15) NOT NULL,
	pin TEXT NOT NULL,
	is_verified BOOLEAN NOT NULL,
	verified_by VARCHAR(20) NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE users;
