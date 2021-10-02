
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id`         INT(10)  UNSIGNED    NOT NULL AUTO_INCREMENT,
  `name`       VARCHAR(20)          NOT NULL,
  `email`      VARCHAR(100)         NOT NULL,
  `age`        int(10)  UNSIGNED,
  `address`    VARCHAR(150)         NULL,
  `height`     FLOAT    UNSIGNED,
  `public`     TINYINT(1) DEFAULT 1 NOT NULL,
  `birthday`   DATETIME             NULL,
  `created_at` DATETIME             NOT NULL,
  `updated_at` DATETIME             NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT email_unique UNIQUE(email)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE `users`;