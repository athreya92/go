CREATE TABLE IF NOT EXISTS payload_info (
    type text,
    properties tuple<text,text>,
    payload_id text,
    PRIMARY KEY((type, properties),payload_id)
);