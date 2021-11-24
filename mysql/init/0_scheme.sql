CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) NOT NULL PRIMARY KEY,
  `name` varchar(36) NOT NULL,
  `hashed_password` varchar(255) NOT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `dashboards` (
  `id` char(36) NOT NULL PRIMARY KEY,
  `user_id` char(36) NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  FOREIGN KEY fk_user_id (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `task_status` (
  `id` char(36) NOT NULL PRIMARY KEY,
  `dashboard_id` char(36) NOT NULL,
  `name` varchar(50) NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  FOREIGN KEY fk_dashboard_id (`dashboard_id`) REFERENCES `dashboards` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `tasks` (
  `id` char(36) NOT NULL PRIMARY KEY,
  `task_status_id` char(36) NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  FOREIGN KEY fk_task_status_id (`task_status_id`) REFERENCES `task_status` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;
