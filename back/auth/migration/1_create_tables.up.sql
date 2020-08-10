CREATE TABLE users (
        id serial PRIMARY KEY,
        nickname VARCHAR(30),
        password_hash VARCHAR(200),
);
