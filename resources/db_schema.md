# Database Schema

```
Users
+----------------+       TrainHistory
| id (PK)        |<-------+---------+      
| username       |        | id (PK) |      
| email          |        | user_id |----+ 
| password_hash  |        | train_id|    | 
| current_train_id|------+ | completed_at| 
+----------------+      |  +---------+    | 
                        |                 | 
                    Train                 | 
       +-----------+------+              | 
       | id (PK)          |              | 
       | name             |              | 
       | created_at       |<-------------+ 
       +------------------+              
                 |                        
                 |                        
        TrainSections                     
       +-----------+--------+           
       | id (PK)           |           
       | train_id (FK)     |           
       | section_name (A/B)|           
       +-------------------+           
                 |                      
                 |                      
        TrainExercises                  SectionCompletions
       +-----------+---------+          +-------------------+
       | id (PK)             |          | id (PK)          |
       | section_id (FK)     |<---------| section_id (FK)  |
       | exercise_id (FK)    |          | user_id (FK)     |
       | sets               |           | completed_at     |
       | reps               |           +-------------------+
       | comment            |                        
       +--------------------+           
                 |                      
                 |                      
              Exercises                  
       +-------------------+            
       | id (PK)           |            
       | name              |            
       | description       |            
       | is_machine        |            
       +-------------------+            
```
