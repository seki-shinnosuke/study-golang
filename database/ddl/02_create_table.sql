USE study_golang;
SET CHARSET utf8mb4;

-- task_management
DROP TABLE IF EXISTS task_management;
CREATE TABLE task_management
(
  task_id INT NOT NULL AUTO_INCREMENT COMMENT 'タスクID',
  person_name VARCHAR(128) NOT NULL COMMENT '担当者名',
  task_name VARCHAR(128) NOT NULL COMMENT 'タスク名',
  deadline_date DATE COMMENT '期限日',
  task_status ENUM('NO_PROCESSING', 'PROCESSING', 'DONE') NOT NULL DEFAULT 'NO_PROCESSING' COMMENT 'ステータス(未着手/着手中/完了)',
  register_datetime DATETIME COMMENT '登録日時',
  update_datetime DATETIME COMMENT '更新日時',
  PRIMARY KEY(task_id)
) COMMENT='タスク管理';