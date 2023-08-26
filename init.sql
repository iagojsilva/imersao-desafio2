CREATE TABLE routes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    source JSON NOT NULL,
    destination JSON NOT NULL
);
