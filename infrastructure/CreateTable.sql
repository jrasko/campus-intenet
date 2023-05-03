CREATE TABLE IF NOT EXISTS network_configs
(
    mac      TEXT PRIMARY KEY,
    name     TEXT,
    room_nr  TEXT,
    has_paid BOOLEAN,
    wg       TEXT,
    email    TEXT,
    phone    TEXT
);