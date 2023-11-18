-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_task" table
CREATE TABLE `new_task` (`id` text NOT NULL, `title` text NOT NULL, `item_status` text NOT NULL DEFAULT ('IN_PROGRESS'), `todo_tasks` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `task_todo_tasks` FOREIGN KEY (`todo_tasks`) REFERENCES `todo` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "task" to new temporary table "new_task"
INSERT INTO `new_task` (`id`, `title`, `item_status`, `todo_tasks`) SELECT `id`, `title`, `item_status`, `todo_tasks` FROM `task`;
-- Drop "task" table after copying rows
DROP TABLE `task`;
-- Rename temporary table "new_task" to "task"
ALTER TABLE `new_task` RENAME TO `task`;
-- Create index "task_id" to table: "task"
CREATE UNIQUE INDEX `task_id` ON `task` (`id`);
-- Create index "task_title_item_status" to table: "task"
CREATE INDEX `task_title_item_status` ON `task` (`title`, `item_status`);
-- Create "new_todo" table
CREATE TABLE `new_todo` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `title` text NOT NULL, `body` text NOT NULL, `item_priority` text NOT NULL DEFAULT ('NONE'), `item_status` text NOT NULL DEFAULT ('IN_PROGRESS'), `time_completed` datetime NULL, `user_todos` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `todo_user_todos` FOREIGN KEY (`user_todos`) REFERENCES `user` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "todo" to new temporary table "new_todo"
INSERT INTO `new_todo` (`id`, `created_at`, `updated_at`, `title`, `body`, `item_priority`, `item_status`, `time_completed`, `user_todos`) SELECT `id`, `created_at`, `updated_at`, `title`, `body`, `item_priority`, `item_status`, `time_completed`, `user_todos` FROM `todo`;
-- Drop "todo" table after copying rows
DROP TABLE `todo`;
-- Rename temporary table "new_todo" to "todo"
ALTER TABLE `new_todo` RENAME TO `todo`;
-- Create index "todo_id" to table: "todo"
CREATE UNIQUE INDEX `todo_id` ON `todo` (`id`);
-- Create index "todo_title_item_priority_item_status" to table: "todo"
CREATE INDEX `todo_title_item_priority_item_status` ON `todo` (`title`, `item_priority`, `item_status`);
-- Create index "user_id" to table: "user"
CREATE UNIQUE INDEX `user_id` ON `user` (`id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
