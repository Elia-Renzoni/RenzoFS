-- RenzoFS Custom Database
-- This file contains a list of the
-- mains query used to setup the entire
-- database

-- I used a dockerized PostgreSQL database

CREATE DATABASE renzofsdb;

CREATE TABLE users (
    username VARCHAR (255) PRIMARY KEY,
    password VARCHAR (255) NOT NULL
);

CREATE TABLE folders (
    folder_id BIGSERIAL PRIMARY KEY,
    folder_name VARCHAR (255) UNIQUE NOT NULL,
    username VARCHAR(255) NOT NULL,
    FOREIGN KEY(username) REFERENCES users(username)
);

CREATE TABLE friends (
    friendship_id BIGSERIAL PRIMARY KEY,
    user1 VARCHAR(255) NOT NULL,
    user2 VARCHAR(255) NOT NULL,
    FOREIGN KEY (user1) REFERENCES users(username),
    FOREIGN KEY (user2) REFERENCES users(username)
);

-- docker exec -it postgresql bash
-- psql -U elia