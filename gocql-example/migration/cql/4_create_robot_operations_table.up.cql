CREATE TABLE IF NOT EXISTS robot_actionInfo (
    robot_id text,
    actions list<text>,
    PRIMARY KEY (robot_id)
);