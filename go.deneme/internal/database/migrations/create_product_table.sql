
DROP TABLE products;

CREATE TABLE IF NOT EXISTS products (
	ID BIGSERIAL PRIMARY KEY,
	Name VARCHAR ( 50 )  NOT NULL,
    Description VARCHAR ( 250 )  NOT NULL,
    BrandName VARCHAR ( 150 ),
	Price DECIMAL ( 8,2 ) NOT NULL,
    Stock INT
);