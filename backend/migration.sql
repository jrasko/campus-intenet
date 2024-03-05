INSERT INTO rooms(number, wg, block)
SELECT room_nr, wg, substr(room_nr, 0, 2)
FROM member_configs;

INSERT INTO net_configs(mac, ip, manufacturer, disabled, created_at, updated_at)
SELECT mac, ip, manufacturer, disabled, created_at, updated_at
FROM member_configs
WHERE firstname NOT ILIKE 'LEER';


INSERT INTO members(id, firstname, lastname, has_paid, email, phone, comment, last_editor, room_nr,
                    created_at, updated_at)
SELECT id,
       firstname,
       lastname,
       has_paid,
       email,
       phone,
       comment,
       last_editor,
       room_nr,
       created_at,
       updated_at
FROM member_configs
WHERE firstname NOT ILIKE 'LEER';

UPDATE members
SET net_config_id = (SELECT id
                     FROM net_configs n
                     WHERE n.mac =
                           (SELECT mac
                            FROM member_configs
                            WHERE member_configs.id = members.id))
WHERE TRUE
