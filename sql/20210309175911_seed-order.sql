-- +migrate Up
INSERT INTO orders (user_id, payment_status, paid_by, paid_at, total_price, order_status) VALUES (1, "PAID", "LINKAJA", NOW(), 25000, "COMPLETED");

INSERT INTO order_details (order_id, menu_id, price_menu) VALUES (1, 1, 20000);
INSERT INTO order_details (order_id, menu_id, price_menu) VALUES (1, 1, 20000);

INSERT INTO item_details (order_detail_id, additional_id, additional_price) VALUES (1, 1, 5000);
INSERT INTO item_details (order_detail_id, additional_id, additional_price) VALUES (2, 1, 5000);
INSERT INTO item_details (order_detail_id, additional_id, additional_price) VALUES (2, 1, 5000);

-- +migrate Down
DELETE FROM item_details WHERE id = 1;
DELETE FROM item_details WHERE id = 2;
DELETE FROM item_details WHERE id = 3;

DELETE FROM order_details WHERE id = 1;
DELETE FROM order_details WHERE id = 2;

DELETE FROM orders WHERE id = 1;