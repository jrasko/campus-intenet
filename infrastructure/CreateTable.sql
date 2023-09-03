CREATE TABLE IF NOT EXISTS member_configs
(
    id        SERIAL PRIMARY KEY,
    mac       macaddr UNIQUE,
    firstname TEXT,
    lastname  TEXT,
    room_nr   TEXT UNIQUE,
    has_paid  BOOLEAN,
    wg        TEXT,
    email     TEXT,
    phone     TEXT,
    ip        TEXT UNIQUE
);
