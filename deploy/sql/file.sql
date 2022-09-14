CREATE TABLE `file`
(
    `id`         bigint unsigned NOT NULL,
    `created_at` timestamp                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `user_id`    bigint unsigned NOT NULL,
    `filename`   varchar(100)                                              NOT NULL,
    `md5`        char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
