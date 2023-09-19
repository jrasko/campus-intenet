**Campus-Internet** is an application to manage a whitelisting DHCP server. Since this application will be used by
non-technical users, the following german documentation is written as non-technical as possible.

I use arch btw.

---
<!-- TOC -->

* [Schnelleinstieg](#schnelleinstieg)
* [Übersicht](#übersicht)
    * [Aufgabe](#aufgabe)
    * [Die Anwendung](#die-anwendung)
        * [DHCPv4-Server](#dhcpv4-server)
        * [Frontend](#frontend)
        * [Backend](#backend)
        * [Datenbank](#datenbank)
    * [Docker](#docker)
* [Installation](#installation)
    * [Anforderungen an Hardware](#anforderungen-an-hardware)
    * [Anforderungen an Software](#anforderungen-an-software)
    * [Installation von Docker](#installation-von-docker)
    * [Installation der Anwendung](#installation-der-anwendung)
    * [Ordner](#ordner)
* [Konfiguration](#konfiguration)
    * [Schnellconfig](#schnellconfig)
    * [Konfiguration mit einer env-Datei](#konfiguration-mit-einer-env-datei)
    * [DB/Postgres](#dbpostgres)
    * [Authentifikation](#authentifikation)
        * [Erläuterung](#erläuterung)
        * [Konfiguration](#konfiguration-1)
    * [Passwort Hash](#passwort-hash)
        * [Erläuterung](#erläuterung-1)
        * [Hash aus Passwort erstellen](#hash-aus-passwort-erstellen)
    * [CIDR](#cidr)
    * [Sonstiges](#sonstiges)
* [Nutzung und Debugging](#nutzung-und-debugging)
    * [Starten und stoppen der Anwendung](#starten-und-stoppen-der-anwendung)
    * [Logs](#logs)
    * [Bearbeitung am DHCPv4-Server oder am Code](#bearbeitung-am-dhcpv4-server-oder-am-code)

<!-- TOC -->
---

# Schnelleinstieg

Du möchtest schnell loslegen ohne viel zu lesen? Los!

1. Installiere Docker
2. Lade den Ordner herunter
3. Navigiere in den Unterordner *infrastructure*
4. Erstelle eine Konfigurationsdatei `.env`, dazu kann diese [Vorlage](#schnellconfig) verwendet werden
5. Erstelle ein Docker Volume für die Datenbank mit `docker volume create dhcp-db`, falls nicht vorhanden
6. Starte die Anwendung mit `docker-compose up -d`
7. Prüfe mit `docker ps` ob alles funktioniert hat

Etwas funktioniert nicht oder diese Anleitung ist zu ungenau? Dann hier die ausführliche Dokumentation:

# Übersicht

## Aufgabe

Die Anwendung soll eine einfache und wartbare Lösung darstellen, um in einem Netzwerk nur bestimmten Geräten Zugriff auf
das Internet zu gewähren. Dazu soll mithilfe der Anwendung eine Liste an erlaubten Geräten (die sogenannte *whitelist*)
erstellt werden.

Im Programm sollen die Nutzer und die sogenannte *MAC-Addresse* ihrer Geräte hinterlegt werden. Erst wenn die
*MAC-Adresse* eines Gerätes auf der *Whitelist* steht, bekommt das Gerät eine *IPv4-Adresse* zugewiesen und auf das
Internet zugreifen. Ohne *IPv4-Adresse* ist ein Internetzugriff nicht möglich.

## Die Anwendung

Das Programm besteht aus mehreren einzelnen Services. Diese werden im Folgendem aufgelistet und ihre Funktion erläutert.

### DHCPv4-Server

Der DHCP-Server nutzt das *Dynamic Host Configuration Protocol v4 (DHCPv4)* um Geräten auf Anfrage eine *IPv4-Adresse*
zuzuteilen.
Wir nutzen in dieser Anwendung den [Kea DHCP Server](https://www.isc.org/kea/). Dieser wird in der kea-dhcp4.conf
konfiguriert.
Die Dokumentation zur Konfiguration der hier verwendet *Kea Version 2.4.0* findet sich hier:
https://kea.readthedocs.io/en/kea-2.4.0

### Frontend

Das Frontend ist die Website, auf der Nutzer angelegt, bearbeitet oder gelöscht werden können. Es kann ganz normal über
einen gängigen Webbrowser aufgerufen werden. Alle angezeigten Daten werden vom **Backend** abgefragt, alle Formulare
werden
an das **Backend** gesendet.

### Backend

Alle im **Frontend** eingegebenen Informationen werden an das Backend gesendet. Hier werden zum Beispiel freie
*IPv4-Adressen*überblick
ermittelt und Nutzerdaten mithilfe der **Datenbank** abgespeichert. Das Backend erstellt und verwaltet die *Whitelist*.
Wenn
sich diese ändert, etwa durch einen neuen oder gelöschten Nutzer, sendet das Backend ein Signal an den **DHCPv4-Server
**,
wodurch dieser seine Konfiguration neu lädt. Auch der Login im **Frontend** wird im Backend behandelt.

### Datenbank

Wir nutzen *PostgreSQL 15* als Datenbank. Sie dient zur Persistierung der Nutzerdaten auf der Festplatte des Servers.
Wie bei allen relationalen Datenbanken werden auch hier die Daten in Form von Tabellen gespeichert. Es werden
regelmäßige
Backups empfohlen, dazu später mehr.

## Docker

Um die Anwendung unabhängig vom Betriebssystem zu halten, die Installation zu vereinfachen und den Nutzern das
management
der einzelnen Services zu erleichtern, nutzen wir die Virtualisierungssoftware Docker.
Docker startet die einzelnen Services in sogenannten *Containern*. Ein Docker Container ist vergleichbar mit einer
[Virtuellen Maschine (VM)](https://de.wikipedia.org/wiki/Virtuelle_Maschine), nur sehr viel leichtgewichtiger bezüglich
zusätlicher Arbeisspeicher- und Prozessorbelastung.

Docker-Compose dient dazu, eine Anwendung mit mehreren Services und daher auch mehreren Containern zu managen. Wir
benötigen
also sowohl Docker als auch Docker-compose

# Installation

## Anforderungen an Hardware

Ich empfehle einen Rechner mit mindestens 4GB RAM und einem Prozessor, der mindestens dem an Leistung entspricht, was
2020 mal Mittelklasse war. Wenn du, lieber Netzer, diese Dokumentation liest und darüber nachdenkst einen neuen Rechner
anzuschaffen,
sind wahrscheinlich ausnahmslos alle Prozessoren auf dem Markt ausreichend.

## Anforderungen an Software

Es sollte ein modernes und verbreitetes Betriebssystem installiert werden, idealerweise eine aktuelle Linux-Distribution
mit großer Community und einfacher Bedienung. Es sollte eine Long-Term-Support (LTS) Variante verwendet werden. Wir
nutzen das
aktuellste Ubuntu LTS 22.04

## Installation von Docker

Es gibt sowohl Docker mit grafischer Oberfläche als auch Docker ausschließlich für die Kommandozeile. Da unser Server
keine grafische Oberfläche besitzt, sollte unbedingt Docker ohne grafische Oberfläche installiert werden. Bei Docker
heißt das [Docker Engine](https://docs.docker.com/engine). Die Dokumentation für die Installation der Docker Engine
findet
man hier: https://docs.docker.com/engine/install.

## Installation der Anwendung

Die Anwendung selbst besteht aus einem Ordner mit mehreren Unterordnern und befindet sich in einem GitHub Repository
unter https://github.com/jrasko/campus-intenet.
Es existieren 2 Möglichkeiten, die Anwendung zu installieren:

1. (empfohlen) Es wird *git* verwendet werden. Unter Ubuntu lässt sich *git* ganz einfach mit dem Paketmanager
   installieren: `sudo apt-get insall git`. Mithilfe von *git* kann man sich den Ordner mit folgendem Befehl auf den
   eigenen
   Rechner kopieren: `git clone https://github.com/jrasko/campus-intenet.git`.

2. Alternativ kann man sich die Dateien als ZIP-Ordner auf der GitHub Seite herunterladen und entpacken. Diese Methode
   wirkt einfacher, hat allerdings nicht die Vorteile, die eine Versionskontrolle wie *git* bietet.

## Ordner

Im Ordner befinden sich neben dieser README.md Datei 3 Unterordner.
Konfigurieren, Starten und Stoppen der Anwendung findet im Unterordner **infrastructure** statt.

Die anderen beiden Ordner sind nur dann relevant, wenn an der Programmierung der Services etwas geändert
werden soll.
**backend** enthält den in *go* oder *golang* geschriebenen backend-Service, dessen Aufgabe bereits [oben](#backend)
schon erläutert wurde. In **frontend** befindet sich die in Vue.js geschriebene Webanwendung, auch diese wurde bereits
in einem [oberen Abschnitt](#frontend) beschrieben.

# Konfiguration

## Schnellconfig

Hier eine Vorlage für eine Standardkonfiguration:

```
POSTGRES_PASSWORD= # >= 20 zufällige Zeichen
LOGIN_USER= # beliebiger nutzername
LOGIN_PASSWORD_HASH= # generiert mit argon2
HMAC_SECRET= # >150 zufällige Zeichen
CIDR= # Subnetzmaske mit erster vergebenen IP
```

## Konfiguration mit einer env-Datei

Die Konfiguration der Anwendung findet hauptsächlich mithilfe sogenannter *Umgebungsvariablen* statt. Diese können im
Service ausgelesen werden. Eine Umgebungsvariable hat immer einen Namen und einen Wert. Es sollte unbedingt eine
*.env-Datei* im **infrastructure** Ordner zum Setzen der Umgebungsvariablen verwendet werden.

Dabei sieht die Datei *.env* wie folgt aus:

```
NAME_1=ein-Wert
ANDERE_UMGEBUNGSVARIABLE=anderer-Wert
...
```

Es können die folgenden Konfigurationen in *.env* hinterlegt werden.

| Name                | Pflicht | Services   | Details                               |
|---------------------|---------|------------|---------------------------------------|
| POSTGRES_HOST       |         | Backend,DB | [DB/Postgres](#dbpostgres)            |
| POSTGRES_DB         |         | Backend,DB | [DB/Postgres](#dbpostgres)            |
| POSTGRES_USER       |         | Backend,DB | [DB/Postgres](#dbpostgres)            |
| POSTGRES_PASSWORD   | X       | Backend,DB | [DB/Postgres](#dbpostgres)            |
| LOGIN_USER          | X       | Backend    | [Authentifikation](#authentifikation) |
| LOGIN_PASSWORD_HASH | X       | Backend    | [Authentifikation](#authentifikation) |
| HMAC_SECRET         | X       | Backend    | [Authentifikation](#authentifikation) |
| CIDR                | X       | Backend    | [CIDR](#cidr)                         |
| SKIP_DHCP_RELOAD    |         | Backend    | [Sonstiges](#sonstiges)               |
| URL                 |         | Backend    | [Sonstiges](#sonstiges)               |
| OUTPUT_FILE         |         | Backend    | [Sonstiges](#sonstiges)               |

## DB/Postgres

Diese Konfigurationen betreffen die Datenbank. Außer *POSTGRES_HOST* werden diese Umgebungsvariablen sowohl von der
Datenbank als auch vom Backend eingelesen, damit das Backend eine Verbindung zur Datenbank herstellen kann.

- *POSTGRES_HOST* wird nur vom Backend eingelesen und gibt an, unter welcher IP oder welchem Hostnamen die Datenbank zu
  finden ist. Da innerhalb des Docker-Netzwerkes die Containernamen als Hostnamen verwendet werden, ist POSTGRES_HOST
  standardmäßig `dhcp_db`, also der Name des Datenbankcontainers.
- *POSTGRES_DB* ist der Name der Standard-Datenbank. Dieser ist defaultmäßig `postgres`.
- *POSTGRES_USER* ist der Standard-Benutzername der Datenbank. Defaultwert ist `postgres`.
- *POSTGRES_PASSWORD* ist das Password für den Standard-Benutzer. Dies ist das einzige verpflichtende Feld. Ich empfehle
  etwa 20 zufällige Zeichen.

## Authentifikation

### Erläuterung

Wenn sich ein Nutzer im Frontend anmeldet, prüft das Backend Benutzername und Passwort und erstellt dann ein sogenanntes
Token (JWT), das 2 Stunden gültig ist und bei jeder Anfrage an das Backend mitgesendet wird. Das Token wird mithilfe
einer
kryptografischen Funktion signiert, dazu wird ein *Secret* benötigt. Jeder, der Kenntnis über das Secret hat, kann
valide
Token ausstellen. Daher sollte das *Secret* niemals auf irgendeine Weise mit irgendjemanden geteilt und ausschließlich
in
der .env Datei vorhanden sein.

### Konfiguration

Damit ergeben sich folgende Konfigurationen:

- *LOGIN_USER* ist der verwendete Nutzername, der auch im frontend angegeben werden muss
- *LOGIN_PASSWORD_HASH* ist der Hash des Passwortes, näheres dazu im [nächsten Abschnitt](#passwort-hash)
- *HMAC_SECRET* ist das bereits erläuterte Secret, mit dem die Token signiert werden. Als Wert können 128 zufällige
  bytes mit einem kryptografischen Zufallsgenerator erzeugt und mit Base64 kodiert werden.

## Passwort Hash

### Erläuterung

Um zu verhindern, dass ein Angreifer im Falle einer schweren Sicherheitslücke das Passwort möglicherweise mehrfach
verwendete Passwort im Klartext auslesen kann, wird in der .env Datei ein Hash-Wert des Passwortes gespeichert. Ein
kryptografischer Hash ist eine Einwegfunktion, sodass sich nicht aus dem Hash auf das Passwort schließen lässt.

Als Hashfunktion nutzen wir Argon2ID, welche Schutzfunktionen gegen Angriffe mit spezialisierter Hardware oder mit
mehreren Threads anbietet. Argon2ID kann über verschiedene Parameter eingestellt werden:

- _Memory_ ist die Belastung des Arbeitsspeichers zur Berechnung des Hashes.
- _Parallelism_: Die Anzahl an aufzuwendenden Threads. Muss eine Zahl zwischen 1 und den zur Verfügung stehenden
  CPU-Kernen sein.
- _Salt_ sollte etwa aus 16 Bytes an Daten (entspricht >= 20 Zeichen) lang sein und aus zufälligen Zeichen bestehen
- _KeyLength_ 32 Bytes empfohlen, sollte nur aus gutem Grund geändert werden
- _Iterations_ gibt an, wie lange die Berechnung dauert. Am besten werden erst alle anderen Faktoren eingestellt und
  _Iterations_ auf 1 gesetzt. Erst danach sollte Iterations so eingestellt werden, dass Generierung eines Hashes etwa 1s
  dauert.

_Memory_ und _Parallelism_ sollten so eingestellt werden, ohne dass durch das Hashing so viele Ressourcen
beansprucht werden, dass andere Prozesse gestört werden. So hoch wie möglich, so niedrig wie nötig.
### Hash aus Passwort erstellen

Unter Ubuntu kann das Programm _argon2_ benutzt werden. Es kann mit `sudo apt-get install argon2` installiert werden.<br>
Es wird wie folgt benutzt:
```
echo -n <password> | argon2 <salt> -id -m <memory> -p <parallelism> -t <iterations>
```

KeyLength 32 ist bereits der Standardwert. Mit `argon2 -h` wird der Befehl erklärt und Einheiten angezeigt.

Das Programm gibt sowohl den Hash, der in die Konfiguration übernommen werden kann, als auch die benötigte Zeit an.

## CIDR

Mit dem Feld *CIDR* lässt sich Einstellen, in welchem Subnetz IP-Addressen vergeben werden. Dazu soll die erste
IP-Addresse angegeben werden, die verteilt werden sollte, sowie die Subnetzmaske.

Beispiel: <br>
`192.168.1.15/24` würde zwischen `192.168.1.15` und `192.168.1.254` IP-Adressen vergeben. <br>
`192.168.0.20/16` würde zwischen `192.168.0.20` und `192.168.255.254` Adressen vergeben.

Die letzte IP eines Subnetzes ist die Broadcast-Adresse und bleibt daher unbelegt.

## Sonstiges

- *SKIP_DHCP_RELOAD* verhindert das Senden eines config-reload Signals des Backends an den DHCP-Server. Dadurch kann das
  Backend funktionieren, auch wenn kein DHCP Server läuft.
- *URL* sollte nicht geändert werden, außer es gibt einen triftigen Grund. Ändert nur die URL **INNERHALB** des
  docker containers. Zum Ändern der von außen erreichbaren URL, docker umkonfigurieren (siehe unten).
- *OUTPUT_FILE* sollte nicht geändert werden, außer es gibt einen triftigen Grund. Ändert nur den Dateipfad
  **INNERHALB** des docker containers. Zum Ändern des Pfades auf dem Rechner, docker umkonfigurieren (siehe unten).

# Nutzung und Debugging

## Starten und stoppen der Anwendung

Zum Starten der Anwendung muss im Ordner **infrastructure** der Befehl `docker-compose up -d` ausgeführt werden. <br>
Mit `docker-compose down` wird die Anwendung gestoppt.

Mit `docker ps` lassen sich alle in Docker laufenden Services anzeigen. Es empfiehlt sich, nach dem Starten wenige
Sekunden zu warten und mit dem Befehl zu prüfen, ob alle container den Status *UP* haben. Falls nicht, sollte mithilfe
von [Logs](#logs) versucht werden, das Problem zu identifizieren

Einzelne Services (container) lassen sich mit `docker stop <containername>` stoppen und mit
`docker start <containername>` wieder starten. <br>

Sollte ein Programm abstürzen, wird es von Docker automatisch neu gestartet.

Alle wichtigen Daten werden mithilfe von sogenannten Docker volumes gespeichert, sodass sich alle container beliebig
starten, stoppen und sogar löschen lassen, ohne das ein Datenverlust eintritt.

## Logs

`docker logs <containername>` zeigt die Logs von einem container. Mit `-f` werden die Logs auch fortlaufend angezeigt.

Insbesondere die Logs vom Backend und vom DHCPv4-Server sind zur Fehlersuche relevant.

## Bearbeitung am DHCPv4-Server oder am Code

Sollte sich die Konfiguration des DHCPv4-Servers ändern, muss das container Image neu gebaut werden. Dazu muss
`docker-compose up -d --build dhcp4` ausgeführt werden. Dies kann etwas länger dauern.
