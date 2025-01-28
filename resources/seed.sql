-- Exercises Inserts
INSERT INTO Exercises (name) VALUES
('Squat'),
('Bench Press'),
('Deadlift');

-- Tags Inserts
INSERT INTO Tags (name) VALUES
('Strength'),
('Hypertrophy'),
('Endurance');

-- ExerciseTags Inserts
INSERT INTO ExerciseTags (exercise_id, tag_id) VALUES
(1, 1), -- Squat is a Strength exercise
(1, 2), -- Squat is also for Hypertrophy
(2, 1), -- Bench Press is a Strength exercise
(2, 2), -- Bench Press is also for Hypertrophy
(3, 1), -- Deadlift is a Strength exercise
(3, 3); -- Deadlift is also for Endurance
