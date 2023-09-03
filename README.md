# Passwort ändern

Zum Ändern des Passworts sind folgende Schritte zu befolgen

1. Generierung eines Salts (zufälliger String, 20+ Zeichen)
2. Hashen des Passwortes (z.B. mit Funktion argon2.IDKey des go-paketes "crypto" mit Parametern
    - passwort=passwort
    - Salt=Salt
    - time=4
    - memory=512 * 1024 KB
    - 8 threads
    - 64 keylen

3. Ersetze in der .env-Datei Salt und Hash in den Zeilen **SALT=<>** und **LOGIN_PASSWORD_HASH=<>**


# Nutzer ändern
Einfach in der .env-Datei LOGIN_USER=*username* ersetzen.
