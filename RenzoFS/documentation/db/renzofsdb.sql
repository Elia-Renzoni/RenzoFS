-- RenzoFS Custom Database
-- This file contains a list of the
-- mains query used to setup the entire
-- database

-- I used a dockerized PostgreSQL database

CREATE DATABASE renzofsdb;

CREATE TABLE users (
    username VARCHAR (50) PRIMARY KEY,
    password VARCHAR (50) NOT NULL
);

CREATE TABLE folders (
    folder_id BIGSERIAL PRIMARY KEY,
    folder_name VARCHAR (50) UNIQUE NOT NULL,
    username VARCHAR(50) NOT NULL,
    FOREIGN KEY(username) REFERENCES users(username)
);

CREATE TABLE friends (
    friendship_id BIGSERIAL PRIMARY KEY,
    user1 VARCHAR(50) NOT NULL,
    user2 VARCHAR(50) NOT NULL,
    FOREIGN KEY (user1) REFERENCES users(username),
    FOREIGN KEY (user2) REFERENCES users(username)
);