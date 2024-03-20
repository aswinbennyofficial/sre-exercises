-- Create Student table to store information about students
CREATE TABLE IF NOT EXISTS Student (
    id SERIAL PRIMARY KEY, -- Unique log ID, Primary Key, Auto-incrementing number
    Name VARCHAR(255) NOT NULL, -- Name of the student
    Phone VARCHAR(15) NOT NULL, -- Phone number of the student
    Address VARCHAR(255) NOT NULL -- Address of the student
);


-- Add comments to describe the table and its columns
COMMENT ON TABLE Student IS 'Table to store information about students';

COMMENT ON COLUMN Student.id IS 'Unique log ID, Primary Key, Auto-incrementing number';
COMMENT ON COLUMN Student.Name IS 'Name of the student';
COMMENT ON COLUMN Student.Phone IS 'Phone number of the student';
COMMENT ON COLUMN Student.Address IS 'Address of the student';
