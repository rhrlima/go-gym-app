-- Idempotent SQL script for gym app database structure

-- Database
CREATE SCHEMA IF NOT EXISTS Gym;

-- -- Users Table
-- CREATE TABLE IF NOT EXISTS Users (
--     id SERIAL PRIMARY KEY,
--     username VARCHAR(50) UNIQUE NOT NULL,
--     email VARCHAR(100) UNIQUE NOT NULL,
--     password_hash VARCHAR(255) NOT NULL,
--     current_train_id INT,
--     FOREIGN KEY (current_train_id) REFERENCES Train(id) ON DELETE SET NULL
-- );

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

-- -- Train Table
-- CREATE TABLE IF NOT EXISTS Train (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(100) NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- -- TrainSections Table
-- CREATE TABLE IF NOT EXISTS TrainSections (
--     id SERIAL PRIMARY KEY,
--     train_id INT NOT NULL,
--     section_name CHAR(1) NOT NULL CHECK (section_name IN ('A', 'B', 'C', 'D')),
--     FOREIGN KEY (train_id) REFERENCES Train(id) ON DELETE CASCADE
-- );

-- -- TrainExercises Table
-- CREATE TABLE IF NOT EXISTS TrainExercises (
--     id SERIAL PRIMARY KEY,
--     section_id INT NOT NULL,
--     exercise_id INT NOT NULL,
--     sets INT NOT NULL,
--     reps INT NOT NULL,
--     comment TEXT,
--     FOREIGN KEY (section_id) REFERENCES TrainSections(id) ON DELETE CASCADE,
--     FOREIGN KEY (exercise_id) REFERENCES Exercises(id) ON DELETE CASCADE
-- );

-- -- TrainHistory Table
-- CREATE TABLE IF NOT EXISTS TrainHistory (
--     id SERIAL PRIMARY KEY,
--     user_id INT NOT NULL,
--     train_id INT NOT NULL,
--     completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
--     FOREIGN KEY (train_id) REFERENCES Train(id) ON DELETE CASCADE
-- );

-- -- SectionCompletions Table
-- CREATE TABLE IF NOT EXISTS SectionCompletions (
--     id SERIAL PRIMARY KEY,
--     user_id INT NOT NULL,
--     section_id INT NOT NULL,
--     completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
--     FOREIGN KEY (section_id) REFERENCES TrainSections(id) ON DELETE CASCADE
-- );
