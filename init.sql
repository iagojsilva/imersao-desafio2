CREATE TABLE routes (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    source JSON NOT NULL,
    destination JSON NOT NULL
);