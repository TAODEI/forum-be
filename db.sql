-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`            int(11) AUTO_INCREMENT PRIMARY KEY,
    `name`          varchar(20) NOT NULL,
    `email`         varchar(35) UNIQUE DEFAULT NULL,
    `avatar`        varchar(100)       DEFAULT NULL,
    `student_id`    char(10) UNIQUE    DEFAULT NULL,
    `hash_password` varchar(100)       DEFAULT NULL,
    `role`          int(11)     NOT NULL COMMENT '权限 0-无权限用户 1-普通学生用户 2（闲置，可能学生管理员） 3-团队成员 4-团队管理员',
    `signature`     varchar(200)       DEFAULT NULL,
    `re`            tinyint(1)         DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
    KEY (`email`),
    UNIQUE KEY (`student_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = DYNAMIC;

-- --------------------------------------------
-- Table structure for posts
-- --------------------------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts`
(
    `id`             int(11) AUTO_INCREMENT PRIMARY KEY,
    `type_name`      varchar(30)  NOT NULL,
    `content`        text         NOT NULL,
    `title`          varchar(150) NOT NULL,
    `create_time`    varchar(30)  NOT NULL,
    `category_id`    int(11)      NOT NULL,
    `re`             tinyint(1)   NOT NULL,
    `creator_id`     int(11)      NOT NULL,
    `last_edit_time` varchar(30)  NOT NULL,
    `main_post_id`   int(11)      NOT NULL,
    `like_num`       int(11) DEFAULT 0,
    KEY (`category_id`),
    FOREIGN KEY (`main_post_id`) REFERENCES `posts` (`id`),
    FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`
(
    `id`          int(11) AUTO_INCREMENT PRIMARY KEY,
    `type_name`   varchar(30) NOT NULL,
    `content`     text        NOT NULL,
    `father_id`   int(11)     DEFAULT NULL,
    `create_time` varchar(30) DEFAULT NULL,
    `re`          tinyint(1)  DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
    `creator_id`  int(11)     DEFAULT NULL,
    `post_id`     int(11)     DEFAULT NULL,
    `like_num`    int(11)     DEFAULT 0,
    FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = DYNAMIC;

-- --------------------------------------------
-- Table structure for tags
-- --------------------------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`
(
    `id`      int(11) AUTO_INCREMENT PRIMARY KEY,
    `content` varchar(30) NOT NULL,
    KEY (`content`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC;

-- --------------------------------------------
-- Table structure for post2tags
-- --------------------------------------------
DROP TABLE IF EXISTS `post2tags`;
CREATE TABLE `post2tags`
(
    `id`      int(11) AUTO_INCREMENT PRIMARY KEY,
    `post_id` int(11) NOT NULL,
    `tag_id`  int(11) NOT NULL,
    KEY (`post_id`, `tag_id`),
    FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
    FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC;