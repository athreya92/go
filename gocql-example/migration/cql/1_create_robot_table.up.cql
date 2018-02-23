CREATE TABLE IF NOT EXISTS robot_basicInfo (
	robot_id text,
	name text,
	type text,
	host_ip text,
	vendor text,
	model text,
	PRIMARY KEY(type ,robot_id)
);

