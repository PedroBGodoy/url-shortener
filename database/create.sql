CREATE TABLE bitlink (
    id char(7) primary key,
    long_url varchar(2083),
    domain varchar(255),
    created_at timestamp
);