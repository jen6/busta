CREATE TABLE USER (
	IDX INT AUTO_INCREMENT PRIMARY KEY,
	Created BIGINT,
	UserId VARCHAR(10) NOT NULL,
	UserPw VARCHAR(32) NOT NULL,
	UserName VARCHAR(5) NOT NULL,
	SUBJECT INT,
	GRADE INT,
	CLASS INT,
	NUM INT
);

CREATE TABLE BUSBOARD (
  	Id 	INT AUTO_INCREMENT PRIMARY KEY,
  	Created BIGINT,
  	Writer VARCHAR(10) NOT NULL,
  	Title  VARCHAR(20),
  	Content VARCHAR(50),
  	Want INT,
  	status INT
);

CREATE TABLE PROFILE (
	IDX INT AUTO_INCREMENT PRIMARY KEY,
	Written BIGINT,
	Id VARCHAR(10),
	Best VARCHAR(50),
	Can VARCHAR(50),
	intro VARCHAR(50),
);

CREATE TABLE PORTFOLIO (
	IDX INT AUTO_INCREMENT PRIMARY KEY,
	Created BIGINT,
	UserId VARCHAR(10)
);

CREATE TABLE AWARD(
);
