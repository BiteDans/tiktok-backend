-- CREATE TABLE `user`
-- (
--     `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
--     `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
--     `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
--     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
--     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
--     `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
--     PRIMARY KEY (`id`),
--     KEY          `idx_username` (`username`) COMMENT 'Username index'
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

/*CREATE TABLE `video`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `author_id`   bigint unsigned NOT NULL COMMENT 'Author (user) id',
    `author_username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Author (user) username',
    `play_url`   varchar(256) NOT NULL DEFAULT '' COMMENT 'Url to where the video is stored',
    `cover_url`   varchar(256) NOT NULL DEFAULT '' COMMENT 'Url to where the cover of the video is stored',
    `favorite_count`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Number of likes one video has received',
    `comment_count`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Number of comments one video has received',
    `likes`     text NOT NULL COMMENT 'Array of user ids that like the video',
    `comments`      text NOT NULL COMMENT 'Array of comments under the video',
    `title`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Title of the video',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video info table';

CREATE TABLE `message`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `to_user_id`   bigint unsigned NOT NULL COMMENT 'Receiver id',
    `from_user_id`   bigint unsigned NOT NULL COMMENT 'Sender id',
    `content` varchar(256) NOT NULL DEFAULT '' COMMENT 'Message content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Message create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Message update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Message delete time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Message table';
CREATE TABLE `comment`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`    bigint unsigned NOT NULL COMMENT 'user id of the comment',
    `video_id`    bigint unsigned NOT NULL COMMENT 'video id of the comment',
    `content`    varchar(256) NOT NULL DEFAULT '' COMMENT 'the content of comment',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment info table';

 */
