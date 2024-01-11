CREATE TABLE segments(
	id INT not null auto_increment,
	name varchar(15),
	primary key(id)
);

CREATE TABLE categories(
	id INT not null auto_increment,
	name varchar(30),
	primary key(id)
);

CREATE TABLE countries(
	id INT not null auto_increment,
	name varchar(50),
	primary key(id)
);


CREATE TABLE regions(
	id INT not null auto_increment,
	name varchar(30),
	primary key(id)
);

CREATE TABLE ship_modes(
	id INT not null auto_increment,
	name varchar(30),
	primary key(id)
);

CREATE TABLE customers(
	id VARCHAR(15) not null,
	name TEXT,
	segment_id int,
	primary key(id),
	foreign key (segment_id) references segments(id) on delete cascade on update cascade
);

CREATE TABLE sub_categories(
	id INT not null auto_increment,
	name varchar(30),
	category_id int,
	primary key(id),
	foreign key (category_id) references categories(id) on delete cascade on update cascade
);

CREATE TABLE states(
	id INT not null auto_increment,
	name varchar(30),
	country_id int,
	primary key(id),
	foreign key (country_id) references countries(id) on delete cascade on update cascade
);

CREATE TABLE cities(
	id INT not null auto_increment,
	name varchar(30),
	state_id int,
	primary key(id),
	foreign key (state_id) references states(id) on delete cascade on update cascade
);

CREATE TABLE products(
	id VARCHAR(30) not null,
	name longtext,
	sub_category_id int,
	numeric_id int not null auto_increment,
	primary key(id),
	key (numeric_id),
	foreign key (sub_category_id) references sub_categories(id) on delete cascade on update cascade
);

CREATE TABLE orders(
	id VARCHAR(20) not NULL,
	order_date date,
	ship_date date,
	postal_code int,
	customer_id VARCHAR(15),
	ship_mode_id int,
	city_id int,
	region_id int,
	primary key(id),
	foreign key (customer_id) references customers(id) on delete cascade on update cascade,
	foreign key (ship_mode_id) references ship_modes(id) on delete cascade on update cascade,
	foreign key (city_id) references cities(id) on delete cascade on update cascade,
	foreign key (region_id) references regions(id) on delete cascade on update cascade
);

CREATE TABLE order_details(
	id int not null auto_increment,
	sales float,
	quantity int,
	discount float,
	profit float,
	product_id VARCHAR(30),
	order_id VARCHAR(20) NOT NULL,
	primary key(id),
	foreign key (product_id) references products(id) on delete cascade on update cascade,
	foreign key (order_id) references orders(id) on delete cascade on update cascade
);

