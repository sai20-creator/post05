-- DROP DATABASE IF EXISTS MSDS;
-- CREATE DATABASE MSDS;

-- DROP TABLE IF EXISTS Users;
-- DROP TABLE IF EXISTS Userdata;

-- \c go;

-- CREATE TABLE Users (
--     ID SERIAL,
--     Username VARCHAR(100) PRIMARY KEY
-- );

-- CREATE TABLE Userdata (
--     UserID Int NOT NULL,
--     Name VARCHAR(100),
--     Surname VARCHAR(100),
--     Description VARCHAR(200)
-- );

Drop the database if it exists
DROP DATABASE IF EXISTS MSDS;

-- Create the MSDS database
CREATE DATABASE MSDS;

-- Connect to the MSDS database
\c MSDS;

-- Create the MSDSCourseCatalog table
CREATE TABLE IF NOT EXISTS MSDSCourseCatalog (
    CID VARCHAR(100) PRIMARY KEY,
    CNAME VARCHAR(200),
    CPREREQ VARCHAR(200)
);
