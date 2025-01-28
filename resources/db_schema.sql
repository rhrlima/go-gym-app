-- Idempotent SQL script for gym app database structure

-- Database
SELECT 'CREATE DATABASE gym'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'gym')

-- Connect to the correct database
\c gym

-- exercises Table
CREATE TABLE IF NOT EXISTS exercises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- tags Table
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- Exercisetags Table
CREATE TABLE IF NOT EXISTS exercise_tags (
    exercise_id INT NOT NULL,
    tag_id INT NOT NULL,
    PRIMARY KEY (exercise_id, tag_id),
    FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- -- train Table
CREATE TABLE IF NOT EXISTS train (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- -- train_sections Table
CREATE TABLE IF NOT EXISTS train_sections (
    id SERIAL PRIMARY KEY,
    train_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    FOREIGN KEY (train_id) REFERENCES train(id) ON DELETE CASCADE
);

-- -- train_exercises Table
CREATE TABLE IF NOT EXISTS train_exercises (
    id SERIAL PRIMARY KEY,
    section_id INT NOT NULL,
    exercise_id INT NOT NULL,
    sets TEXT,
    comment TEXT,
    FOREIGN KEY (section_id) REFERENCES train_sections(id) ON DELETE CASCADE,
    FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE
);
