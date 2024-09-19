-- init.sql

-- Создаем таблицу пользователей
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(50) UNIQUE NOT NULL,
                                     password VARCHAR(100) NOT NULL
);

-- Создаем таблицу задач
CREATE TABLE IF NOT EXISTS tasks (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(255) NOT NULL,
                                     description TEXT,
                                     done BOOLEAN DEFAULT FALSE,
                                     user_id INT,
                                     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
