-- +migrate Up
CREATE TABLE IF NOT EXISTS additionals (
	id INT auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	price DECIMAL(10, 2) NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NULL,
	CONSTRAINT additionals_PK PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS menu_categories (
	id INT auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NULL,
	CONSTRAINT menu_categories_PK PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS menus (
	id INT auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	description text NOT NULL,
	price DECIMAL(10, 2) NOT NULL,
	menu_category_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NULL,
	CONSTRAINT menus_PK PRIMARY KEY (id),
	CONSTRAINT menus_and_menu_categories_id_FK FOREIGN KEY (menu_category_id) REFERENCES menu_categories(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS eligible_additional_menu (
	id INT auto_increment NOT NULL,
	menu_id INT NOT NULL,
	additional_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NULL,
	CONSTRAINT eligible_additional_menu_PK PRIMARY KEY (id),
	CONSTRAINT eligible_additional_menu_and_menus_FK FOREIGN KEY (menu_id) REFERENCES menus(id) ON DELETE CASCADE,
	CONSTRAINT eligible_additional_menu_and_additionals_FK FOREIGN KEY (additional_id) REFERENCES additionals(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE eligible_additional_menu;
DROP TABLE menus;
DROP TABLE menu_categories;
DROP TABLE additionals;
