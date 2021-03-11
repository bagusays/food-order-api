-- +migrate Up
INSERT INTO users (email, phone_number, pin, is_verified, verified_by) VALUES ("admin@gmail.com", "6283877309000", "hashedPIN", true, "email");

-- +migrate Down
DELETE FROM users where email = "admin@gmail.com";
