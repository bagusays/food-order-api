-- +migrate Up
CREATE TABLE IF NOT EXISTS payment_channels (
	id VARCHAR(60) NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT payments_PK PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS orders (
	id INT auto_increment NOT NULL,
	user_id INT NOT NULL,
	payment_status VARCHAR(20) NOT NULL,
	paid_by VARCHAR(60) NULL,
	paid_at TIMESTAMP NULL,
	total_price BIGINT NOT NULL,
	order_status VARCHAR(20) NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT orders_PK PRIMARY KEY (id),
	CONSTRAINT orders_and_users_FK FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT orders_and_payment_channels_FK FOREIGN KEY (paid_by) REFERENCES payment_channels(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_details (
	id INT auto_increment NOT NULL,
	order_id INT NOT NULL,
	menu_id INT NOT NULL,
	price_menu BIGINT NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT order_details_PK PRIMARY KEY (id),
	CONSTRAINT order_details_and_orders_FK FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
	CONSTRAINT order_details_and_menus_FK FOREIGN KEY (menu_id) REFERENCES menus(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS item_details (
	id INT auto_increment NOT NULL,
	order_detail_id INT NOT NULL,
	additional_id INT NOT NULL,
	additional_price BIGINT NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT item_details_PK PRIMARY KEY (id),
	CONSTRAINT item_details_and_order_details_FK FOREIGN KEY (order_detail_id) REFERENCES order_details(id) ON DELETE CASCADE,
	CONSTRAINT item_details_and_additionals_FK FOREIGN KEY (additional_id) REFERENCES additionals(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE item_details;
DROP TABLE order_details;
DROP TABLE orders;
DROP TABLE payment_channels;
