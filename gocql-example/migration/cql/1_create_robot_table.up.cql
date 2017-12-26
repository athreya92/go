CREATE TABLE IF NOT EXISTS robots (
	serialNo text,
	robotName text, 
	robotType text, 
	ipAddress text,
	PRIMARY KEY(serialNo, robotName)
);