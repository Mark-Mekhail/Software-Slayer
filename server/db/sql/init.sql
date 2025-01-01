CREATE TABLE users (
  id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE CHECK (`username` regexp '^[a-zA-Z0-9_ -]{1,30}$'),
  email VARCHAR(255) NOT NULL UNIQUE CHECK (`email` regexp '^[^@]+@[^@]+\.[^@]{2,}$'),
  password_hash VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL
);

CREATE TABLE user_learning_list (
  id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT UNSIGNED NOT NULL,
  title VARCHAR(255) NOT NULL,
  category ENUM('Languages', 'Technologies', 'Concepts', 'Projects', 'Other') NOT NULL,
  UNIQUE (user_id, title, category),
  FOREIGN KEY (user_id) REFERENCES users(id)
);