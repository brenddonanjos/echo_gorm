DROP DATABASE echo_gorm;

CREATE DATABASE IF NOT EXISTS echo_gorm;

CREATE TABLE IF NOT EXISTS echo_gorm.products (
    id INT NOT NULL AUTO_INCREMENT,
    bar_code VARCHAR(255),
    name VARCHAR(255),
    value DECIMAL(13,4),
    quantity INT,
    description TEXT,
    active TINYINT(1),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id)
) ENGINE = innodb;

INSERT INTO echo_gorm.products (id,bar_code, name, value, quantity, description, created_at, updated_at, deleted_at) VALUES (1,"010102544544", "Antitranspirante Ativa Oboticário", 12.33, 850, "Antitranspirante em creme, linha cuide-se bem", "2022-03-12 11:21:00", "2022-03-12 11:21:00", NULL);