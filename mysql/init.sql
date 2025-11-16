CREATE DATABASE IF NOT EXISTS blogdb;
USE blogdb;

CREATE TABLE IF NOT EXISTS blogs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255)
);

INSERT INTO blogs (title) VALUES ('First Blog'), ('Hello World'), ('3-Tier App Ready');
