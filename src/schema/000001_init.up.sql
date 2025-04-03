CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE address (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    country VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    street VARCHAR(255) NOT NULL
);

CREATE TABLE images (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    image BYTEA NOT NULL
);

CREATE TABLE supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    address_id UUID REFERENCES address(id) ON DELETE SET NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE
);

CREATE TABLE client (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    client_name VARCHAR(100) NOT NULL,
    client_surname VARCHAR(100) NOT NULL,
    birthday DATE NOT NULL,
    gender VARCHAR(10) CHECK (gender IN ('Male', 'Female', 'Other')) NOT NULL,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    address_id UUID REFERENCES address(id) ON DELETE SET NULL
);

CREATE TABLE product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    category VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    available_stock INT NOT NULL CHECK (available_stock >= 0),
    last_update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    supplier_id UUID REFERENCES supplier(id) ON DELETE SET NULL,
    image_id UUID REFERENCES images(id) ON DELETE SET NULL
);
