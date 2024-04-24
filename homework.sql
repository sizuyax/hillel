CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(15),
    email VARCHAR(100)
);

CREATE TABLE drivers (
    driver_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(15),
    email VARCHAR(100)
);

CREATE TABLE cars (
    car_id SERIAL PRIMARY KEY,
    driver_id INT,
    car_make VARCHAR(50),
    car_model VARCHAR(50),
    car_year time,
    license_plate VARCHAR(20),

    CONSTRAINT fk_cars
        FOREIGN KEY (driver_id)
            REFERENCES drivers(driver_id)
);

CREATE TABLE addresses (
    address_id SERIAL PRIMARY KEY,
    street VARCHAR(100),
    city VARCHAR(50),
    state VARCHAR(50),
    zip_code VARCHAR(10)
);


CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    customer_id INT,
    driver_id INT,
    order_date time,
    pickup_address_id INT,
    delivery_address_id INT,

    CONSTRAINT fk_customers
	    FOREIGN KEY (customer_id)
    		REFERENCES customers(customer_id),

    CONSTRAINT fk_drivers
    	FOREIGN KEY (driver_id)
    		REFERENCES drivers(driver_id),

    CONSTRAINT fk_pick_addresses
    	FOREIGN KEY (pickup_address_id)
    		REFERENCES addresses(address_id),

    CONSTRAINT fk_delivery_addresses
	    FOREIGN KEY (delivery_address_id)
    		REFERENCES addresses(address_id)
);


CREATE TABLE order_addresses (
    order_id INT,
    address_id INT,
    PRIMARY KEY (order_id, address_id),

    CONSTRAINT fk_orders
    	FOREIGN KEY (order_id)
    		REFERENCES orders(order_id),

    CONSTRAINT fk_addresses
    	FOREIGN KEY (address_id)
    		REFERENCES addresses(address_id)
);



INSERT INTO drivers (first_name, last_name, phone_number, email)
	VALUES ('man', 'bro', '+48791013993', 'bro@gmail.com');


INSERT INTO customers (first_name, last_name, phone_number, email)
	VALUES ('customer', 'one', '+3808319411', 'cust@cust.com');

INSERT INTO cars (driver_id, car_make, car_model, car_year, license_plate)
	VALUES (2, 'BMW', 'Sedan', now(), 'AA1234AO');

INSERT INTO addresses (street, city, state, zip_code)
	VALUES ('Balladyny', 'Lublin', 'Lublin', '20-601');

INSERT INTO addresses (street, city, state, zip_code)
	VALUES ('Skierki', 'Lublin', 'Lublin', '20-601');


INSERT INTO orders (customer_id, driver_id, order_date, pickup_address_id, delivery_address_id)
	VALUES (1, 1, now(), 1, 2);

INSERT INTO order_addresses (order_id, address_id)
	VALUES (1, 2);
