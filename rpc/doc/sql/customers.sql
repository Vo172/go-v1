CREATE TABLE `customers`
(
    `customer_id`               bigint NOT NULL AUTO_INCREMENT,
    `customer_name`             varchar(255) NULL DEFAULT '' COMMENT 'Ten khach hang',
    `email`                     varchar(255) NULL DEFAULT '' COMMENT 'Email',
    `create_date`               timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`customer_id`),
    UNIQUE KEY `customer_id_unique` (`customer_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;