CREATE TABLE `users`
(
    `id`            INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uuid`          VARCHAR(36) NOT NULL UNIQUE ,
    `username`      VARCHAR(36) NOT NULL UNIQUE ,
    `first_name`    VARCHAR(100),
    `last_name`     VARCHAR(100),
    `email`         VARCHAR(100) UNIQUE ,
    `password`      VARCHAR(100),
    `password_seed` VARCHAR(100),
    `created_date`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `modified_date` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;



CREATE TABLE `products`
(
    `id`            INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uuid`          VARCHAR(36) NOT NULL,
    `name`          VARCHAR(100),
    `brand`         VARCHAR(100),
    `created_date`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `modified_date` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uuid` (`uuid`),
    KEY             `name` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


CREATE TABLE `retailers`
(
    `id`            INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uuid`          VARCHAR(36) NOT NULL,
    `name`          VARCHAR(100),
    `email`         VARCHAR(100),
    `created_date`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `modified_date` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uuid` (`uuid`),
    KEY             `email` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


CREATE TABLE `promotions`
(
    `id`            INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uuid`          VARCHAR(36) NOT NULL,
    `product_id`    INT(11),
    `retailer_id`   INT(11),
    `discount`      FLOAT,
    `start_time`    DATETIME DEFAULT NULL,
    `end_time`      DATETIME DEFAULT NULL,
    `is_active`     tinyint(1) DEFAULT 0,
    `created_date`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `modified_date` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uuid` (`uuid`),
    CONSTRAINT `promotions_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `promotions_ibfk_2` FOREIGN KEY (`retailer_id`) REFERENCES `retailers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


CREATE TABLE `products_retailers`
(
    `id`            INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `uuid`          VARCHAR(36) NOT NULL,
    `product_id`    INT(11),
    `retailer_id`   INT(11),
    `price`         FLOAT    DEFAULT '0.0',
    `quantity`      INT(11) DEFAULT 0,
    `created_date`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `modified_date` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uuid` (`uuid`),
    UNIQUE KEY `i_product_retailer` (`product_id`,`retailer_id`),
    CONSTRAINT `products_retailers_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `products_retailers_ibfk_2` FOREIGN KEY (`retailer_id`) REFERENCES `retailers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


