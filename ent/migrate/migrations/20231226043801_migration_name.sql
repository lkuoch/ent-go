-- Create "remotes" table
CREATE TABLE "remotes" ("id" character varying NOT NULL, "data" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "user" table
CREATE TABLE "user" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "username" character varying NOT NULL, "display_name" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "user_display_name" to table: "user"
CREATE INDEX "user_display_name" ON "user" ("display_name");
-- Create index "user_id" to table: "user"
CREATE UNIQUE INDEX "user_id" ON "user" ("id");
-- Create index "user_username" to table: "user"
CREATE UNIQUE INDEX "user_username" ON "user" ("username");
-- Create index "user_username_key" to table: "user"
CREATE UNIQUE INDEX "user_username_key" ON "user" ("username");
-- Create "todo" table
CREATE TABLE "todo" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "title" character varying NOT NULL, "body" character varying NOT NULL, "item_priority" character varying NOT NULL DEFAULT 'NONE', "item_status" character varying NOT NULL DEFAULT 'IN_PROGRESS', "time_completed" timestamptz NULL, "user_todos" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "todo_user_todos" FOREIGN KEY ("user_todos") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "todo_id" to table: "todo"
CREATE UNIQUE INDEX "todo_id" ON "todo" ("id");
-- Create index "todo_title_item_priority_item_status" to table: "todo"
CREATE INDEX "todo_title_item_priority_item_status" ON "todo" ("title", "item_priority", "item_status");
-- Create "task" table
CREATE TABLE "task" ("id" character varying NOT NULL, "title" character varying NOT NULL, "item_status" character varying NOT NULL DEFAULT 'IN_PROGRESS', "todo_tasks" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "task_todo_tasks" FOREIGN KEY ("todo_tasks") REFERENCES "todo" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "task_id" to table: "task"
CREATE UNIQUE INDEX "task_id" ON "task" ("id");
-- Create index "task_title_item_status" to table: "task"
CREATE INDEX "task_title_item_status" ON "task" ("title", "item_status");
