-- Create "todos" table
CREATE TABLE `todos` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `title` text NOT NULL, `priority` text NOT NULL DEFAULT ('NONE'), `status` text NOT NULL DEFAULT ('IN_PROGRESS'), `time_completed` datetime NULL, `todo_parent` text NULL, PRIMARY KEY (`id`), CONSTRAINT `todos_todos_parent` FOREIGN KEY (`todo_parent`) REFERENCES `todos` (`id`) ON DELETE SET NULL);
-- Create index "todo_title_priority_status" to table: "todos"
CREATE INDEX `todo_title_priority_status` ON `todos` (`title`, `priority`, `status`);
-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `username` text NOT NULL, `display_name` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX `users_username_key` ON `users` (`username`);
-- Create index "user_display_name" to table: "users"
CREATE INDEX `user_display_name` ON `users` (`display_name`);
-- Create index "user_username" to table: "users"
CREATE UNIQUE INDEX `user_username` ON `users` (`username`);
-- Create "user_todos" table
CREATE TABLE `user_todos` (`user_id` text NOT NULL, `todo_id` text NOT NULL, PRIMARY KEY (`user_id`, `todo_id`), CONSTRAINT `user_todos_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_todos_todo_id` FOREIGN KEY (`todo_id`) REFERENCES `todos` (`id`) ON DELETE CASCADE);
