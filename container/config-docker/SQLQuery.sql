CREATE DATABASE store;

create table account(
    username varchar(255),
    password varchar(255)
)

create table product(
    id varchar(255) primary key,
    type varchar(255),
    name varchar(255),
    price int
)

INSERT INTO product(id, type, name, price)
VALUES
("123", "đồng hồ", "a", 100),
("124", "đồng hồ", "b", 100),
("125", "đồng hồ", "c", 100),
("126", "đồng hồ", "d", 100),
("127", "đồng hồ", "e", 100),
("128", "đồng hồ", "f", 100),
("129", "túi xách", "g", 100),
("130", "túi xách", "h", 100),
("131", "túi xách", "k", 100),
("132", "túi xách", "t", 100),
("133", "túi xách", "i", 100),
("134", "túi xách", "o", 100)

SELECT * from product

SELECT COUNT(*) FROM (
    SELECT * FROM product
    LIMIT 5 OFFSET 2
) as product

SELECT name, type FROM product LIMIT 5 OFFSET 2 

CREATE INDEX idx_name ON product(name);


select * from account

delete  from product

drop TABLE account