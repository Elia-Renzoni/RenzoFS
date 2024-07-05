-- RenzoFS Custom Database
-- This file contains a list of the
-- mains query used to setup the entire
-- database

-- I used a dockerized PostgreSQL database

CREATE DATABASE renzofsdb;

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL
);
