-- Create "task" table
CREATE TABLE `task` (`id` text NOT NULL, `title` text NOT NULL, `item_status` text NOT NULL DEFAULT ('IN_PROGRESS'), `todo_tasks` text NULL, PRIMARY KEY (`id`), CONSTRAINT `task_todo_tasks` FOREIGN KEY (`todo_tasks`) REFERENCES `todo` (`id`) ON DELETE SET NULL);
-- Create index "task_title_item_status" to table: "task"
CREATE INDEX `task_title_item_status` ON `task` (`title`, `item_status`);
-- Create "todo" table
CREATE TABLE `todo` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `title` text NOT NULL, `body` text NOT NULL, `item_priority` text NOT NULL DEFAULT ('NONE'), `item_status` text NOT NULL DEFAULT ('IN_PROGRESS'), `time_completed` datetime NULL, `user_todos` text NULL, PRIMARY KEY (`id`), CONSTRAINT `todo_user_todos` FOREIGN KEY (`user_todos`) REFERENCES `user` (`id`) ON DELETE SET NULL);
-- Create index "todo_title_item_priority_item_status" to table: "todo"
CREATE INDEX `todo_title_item_priority_item_status` ON `todo` (`title`, `item_priority`, `item_status`);
-- Create "user" table
CREATE TABLE `user` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `username` text NOT NULL, `display_name` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "user_username_key" to table: "user"
CREATE UNIQUE INDEX `user_username_key` ON `user` (`username`);
-- Create index "user_display_name" to table: "user"
CREATE INDEX `user_display_name` ON `user` (`display_name`);
-- Create index "user_username" to table: "user"
CREATE UNIQUE INDEX `user_username` ON `user` (`username`);
