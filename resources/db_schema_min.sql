-- Idempotent SQL script for gym app database structure

-- Database
CREATE DATABASE gym;

\c gym

-- Exercises Table
CREATE TABLE IF NOT EXISTS Exercises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Tags Table
CREATE TABLE IF NOT EXISTS Tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- ExerciseTags Table
CREATE TABLE IF NOT EXISTS ExerciseTags (
    exercise_id INT NOT NULL,
    tag_id INT NOT NULL,
    PRIMARY KEY (exercise_id, tag_id),
    FOREIGN KEY (exercise_id) REFERENCES Exercises(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES Tags(id) ON DELETE CASCADE
);
