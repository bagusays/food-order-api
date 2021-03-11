-- +migrate Up
INSERT INTO payment_channels (id) VALUES ("LINKAJA");

-- +migrate Down
DELETE FROM payment_channels WHERE id = "LINKAJA";
