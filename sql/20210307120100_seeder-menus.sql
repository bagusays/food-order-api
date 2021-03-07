-- +migrate Up
INSERT INTO additionals (name, price) VALUES ("espresso +1", 5000);
INSERT INTO additionals (name, price) VALUES ("soy milk", 3000);
INSERT INTO menu_categories (name) VALUES ("Signature");
INSERT INTO menus (name, description, price, menu_category_id) VALUES ("Latte", "kopi latte", 20000, 1);
INSERT INTO eligible_additional_menu (menu_id, additional_id) VALUES (1, 1);
INSERT INTO eligible_additional_menu (menu_id, additional_id) VALUES (1, 2);

-- +migrate Down
DELETE FROM additionals WHERE name = "espresso +1";
DELETE FROM menu_categories WHERE name = "Signature";
DELETE FROM menus WHERE name = "Latte";
DELETE FROM eligible_additional_menu WHERE menu_id = 1;
