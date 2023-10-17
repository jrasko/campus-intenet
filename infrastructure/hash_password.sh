#! /bin/bash
echo "enter a password"
read -srp "> " PASSWORD
SALT=$(head -c 16 /dev/random | base64 -w 0)
echo
echo -n "$PASSWORD" | argon2 "$SALT" -id -m 20 -p 4 -t 1

