
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS `todos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `finished` boolean NOT NULL default 0,
  `created_at` datetime NOT NULL default current_timestamp,
  `updated_at` timestamp NOT NULL default current_timestamp on update current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE `todos`;