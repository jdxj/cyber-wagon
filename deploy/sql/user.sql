-- im.`user` definition

CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL COMMENT 'userID',
    `created_at` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `nickname`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `email`      varchar(100)                                                  NOT NULL,
    `salt`       varchar(100)                                                  NOT NULL,
    `password`   varchar(100)                                                  NOT NULL,
    PRIMARY KEY (`id`),
    KEY          `user_email_IDX` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
