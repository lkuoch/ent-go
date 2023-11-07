-- Create "todo" table
CREATE TABLE `todo` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `title` text NOT NULL, `priority` text NOT NULL DEFAULT ('NONE'), `status` text NOT NULL DEFAULT ('IN_PROGRESS'), `time_completed` datetime NULL, `todo_parent` text NULL, PRIMARY KEY (`id`), CONSTRAINT `todo_todo_parent` FOREIGN KEY (`todo_parent`) REFERENCES `todo` (`id`) ON DELETE SET NULL);
-- Create index "todo_title_priority_status" to table: "todo"
CREATE INDEX `todo_title_priority_status` ON `todo` (`title`, `priority`, `status`);
-- Create "user" table
CREATE TABLE `user` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `username` text NOT NULL, `display_name` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "user_username_key" to table: "user"
CREATE UNIQUE INDEX `user_username_key` ON `user` (`username`);
-- Create index "user_display_name" to table: "user"
CREATE INDEX `user_display_name` ON `user` (`display_name`);
-- Create index "user_username" to table: "user"
CREATE UNIQUE INDEX `user_username` ON `user` (`username`);
-- Create "user_todos" table
CREATE TABLE `user_todos` (`user_id` text NOT NULL, `todo_id` text NOT NULL, PRIMARY KEY (`user_id`, `todo_id`), CONSTRAINT `user_todos_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE, CONSTRAINT `user_todos_todo_id` FOREIGN KEY (`todo_id`) REFERENCES `todo` (`id`) ON DELETE CASCADE);
