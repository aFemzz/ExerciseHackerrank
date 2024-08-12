-- TASK 2 - SQL: Student Grades
-- Write a query to print the ID and StudentGrade for each record in the STUDENT table. Sort the
-- output by student ID, ascending, and use the following format.
-- Student STUDENT.ID has grade: StudentGrade
-- If Score < 20, StudentGrade = F
-- If 20 <= Score < 40, StudentGrade = D
-- If 40 <= Score < 60, StudentGrade = C
-- If 60 <= Score < 80, StudentGrade = B
-- If Score >= 80, StudentGrade = A

-- ASSUMING THE STRUCTURE OF THE DATABASE IS :
CREATE TABLE StudentScores (
    StudentID INT PRIMARY KEY,
    Score INT NOT NULL
);


SELECT CONCAT('Student ', StudentID, ' has grade: ', 
        CASE
            WHEN Score < 20 THEN 'F'
            WHEN Score >= 20 AND Score < 40 THEN 'D'
            WHEN Score >= 40 AND Score < 60 THEN 'C'
            WHEN Score >= 60 AND Score < 80 THEN 'B'
            WHEN Score >= 80 THEN 'A'
        END
    ) AS Report
    
FROM studentscores
ORDER BY StudentID ASC;