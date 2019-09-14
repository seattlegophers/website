CREATE DATABASE seattleGophers CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE seattleGophers;

CREATE TABLE user_account (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(190) NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created DATETIME NOT NULL,
  sponsor BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE thread (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  subject VARCHAR(255) NOT NULL,
  created DATETIME NOT NULL,
  user_account_id INTEGER NOT NULL,
  FOREIGN KEY (user_account_id) REFERENCES user_account(id)
  );

CREATE TABLE post (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  content VARCHAR(1000) NOT NULL,
  created DATETIME NOT NULL,
  thread_id INTEGER NOT NULL,
  user_account_id INTEGER NOT NULL,
  FOREIGN KEY (user_account_id) REFERENCES user_account(id),
  FOREIGN KEY (thread_id) REFERENCES thread(id)
);

ALTER TABLE user_account ADD UNIQUE (email);
GRANT SELECT, INSERT, UPDATE ON seattleGophers.* TO 'mysqlUser'@'localhost';
ALTER USER 'mysqlUser'@'localhost' IDENTIFIED BY 'pwd';
