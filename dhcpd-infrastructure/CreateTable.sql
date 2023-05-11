CREATE TABLE IF NOT EXISTS network_configs
(
    mac       macaddr PRIMARY KEY,
    firstname TEXT,
    lastname  TEXT,
    room_nr   TEXT,
    has_paid  BOOLEAN,
    wg        TEXT,
    email     TEXT,
    phone     TEXT,
    ip        inet
);

CREATE TABLE IF NOT EXISTS allocated_ips
(
    ip inet PRIMARY KEY
);