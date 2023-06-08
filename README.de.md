# radepa-x-ui
![4](./web/assets/svg/footprint.svg)


# Languages

[![](https://custom-icon-badges.herokuapp.com/badge/Deutch_Sprache-FF6A00?style=for-the-badge&logo=germany-logo&logoColor=white&logoWidth=50)](README.de.md)


Verstecken Sie Ihren Fußabdruck mit Radepa

> **Haftungsausschluss: Dieses Projekt dient nur dem persönlichen Lernen und der Kommunikation, bitte verwenden Sie es nicht für illegale Zwecke, bitte verwenden Sie es nicht in einer Produktionsumgebung**


[![Lizenz](https://img.shields.io/badge/license-GPL%20V3-blue.svg?longCache=true)](https://www.gnu.org/licenses/gpl-3.0.en.html)

radepa-x-ui-Panel unterstützt Multiprotokoll, **Mehrsprachig (Englisch, Farsi, Chinesisch, Russisch)**

**Wenn Sie der Meinung sind, dass dieses Projekt für Sie hilfreich ist, möchten Sie vielleicht eine geben** :star2:


# Installieren und aktualisieren

```
bash <(curl -Ls https://raw.githubusercontent.com/4xmen/radepa-x-ui/master/install.sh)
```

[//]: # (## Benutzerdefinierte Version installieren)

[//]: # ()
[//]: # (Um die gewünschte Version zu installieren, können Sie die Version am Ende des Installationsbefehls hinzufügen. Beispiel für Version „v1.6.0“:)

[//]: # ()
[//]: # (```)

[//]: # (bash <&#40;curl -Ls https://raw.githubusercontent.com/4xmen/radepa-x-ui/master/install.sh&#41; v1.6.0)

[//]: # (```)

# SSL

```
apt-get install certbot -y
certbot certonly --standalone --agree-tos --register-unsafely-without-email -d yourdomain.com
certbot renew --dry-run
```

oder Sie können das x-ui-Menü und dann die Nummer „16“ verwenden (Ein SSL-Zertifikat beantragen).

# Default settings

- Port: 2086
- username and password will be generated randomly if you skip to modify your own security(x-ui "7")
- database path: /etc/x-ui/x-ui.db
- xray config path: /usr/local/x-ui/bin/config.json

- Port: 2086
- Benutzername und Passwort werden zufällig generiert, wenn Sie die Änderung Ihrer eigenen Sicherheit überspringen (x-ui „7“)
- Datenbankpfad: /etc/x-ui/x-ui.db
- xray-Konfigurationspfad: /usr/local/x-ui/bin/config.json

Bevor Sie SSL in den Einstellungen festlegen

- http://ip:2086/
- http://domain:2086/

Nachdem Sie SSL in den Einstellungen festgelegt haben

- https://yourdomain:2086/

# Environment Variables

| Variable       |                      Type                      | Default       |
| -------------- | :--------------------------------------------: | :------------ |
| XUI_LOG_LEVEL  | `"debug"` \| `"info"` \| `"warn"` \| `"error"` | `"info"`      |
| XUI_DEBUG      |                   `boolean`                    | `false`       |
| XUI_BIN_FOLDER |                    `string`                    | `"bin"`       |
| XUI_DB_FOLDER  |                    `string`                    | `"/etc/x-ui"` |

Beispiel:

```sh
XUI_BIN_FOLDER="bin" XUI_DB_FOLDER="/etc/x-ui" go build main.go
```

# Röntgenkonfigurationen:

**Kopieren und in die Xray-Konfiguration einfügen:** (Sie müssen dies nicht tun, wenn Sie eine Neuinstallation haben)

- [traffic](./media/configs/traffic.json)
- [traffic + Block all Iran IP address](./media/configs/traffic+block-iran-ip.json)
- [traffic + Block all Iran Domains](./media/configs/traffic+block-iran-domains.json)
- [traffic + Block Ads + Use IPv4 for Google](./media/configs/traffic+block-ads+ipv4-google.json)
- [traffic + Block Ads + Route Google + Netflix + Spotify + OpenAI (ChatGPT) to WARP](./media/configs/traffic+block-ads+warp.json)

# [WARP-Konfiguration](https://github.com/fscarmen/warp) (Optional)

Wenn Sie Routing zu WARP verwenden möchten, führen Sie die folgenden Schritte aus:

1. Wenn Sie Warp bereits installiert haben, können Sie es mit dem folgenden Befehl deinstallieren:

   ```sh
   warp u
   ```

2. WARP im **Socks-Proxy-Modus** installieren:

   ```sh
   curl -fsSL https://gist.githubusercontent.com/hamid-gh98/dc5dd9b0cc5b0412af927b1ccdb294c7/raw/install_warp_proxy.sh | bash
   ```

3. Aktivieren Sie die Konfiguration, die Sie benötigen, im Panel oder [Kopieren Sie diese Datei und fügen Sie sie in die Xray-Konfiguration ein](./media/configs/traffic+block-ads+warp.json)

Konfigurationsfunktionen:

   - Werbung blockieren
   - Leiten Sie Google + Netflix + Spotify + OpenAI (ChatGPT) an WARP weiter
   - Beheben Sie den Google 403-Fehler

# Merkmale

- Systemstatusüberwachung
- Suche in allen Inbounds und Clients
- Unterstützt die Benutzeroberfläche des Dunkel-/Hell-Designs
- Unterstützt Mehrbenutzer-Mehrprotokoll- und Webseitenvisualisierungsvorgänge
- Unterstützte Protokolle: vmess, vless, Trojaner, Shadowsocks, Dokodemo-Door, Socks, http
– Unterstützung für die Konfiguration weiterer Transportkonfigurationen
- Verkehrsstatistiken, Verkehr begrenzen, Ablaufzeit begrenzen
- Anpassbare Röntgenkonfigurationsvorlagen
- Unterstützung des https-Zugriffspanels (selbst bereitgestellter Domänenname + SSL-Zertifikat)
- Unterstützt die Ein-Klick-SSL-Zertifikatsanwendung und die automatische Verlängerung
- Weitere erweiterte Konfigurationselemente finden Sie im Panel
- Unterstützung zum Ändern von Konfigurationen durch verschiedene im Panel bereitgestellte Elemente
- Unterstützt den Export/Import von Datenbanken aus dem Panel
- Stabile APIs
- Verstecken Sie Ihr Panel-Formular GFW

# Tg-Robotereinsatz

X-UI unterstützt tägliche Verkehrsbenachrichtigungen, Panel-Anmeldeerinnerungen und andere Funktionen über den Tg-Roboter. Um den Tg-Roboter nutzen zu können, müssen Sie sich für das spezifische Anwendungs-Tutorial bewerben. Sie können sich auf den [Blog] (https://coderfan.net/how-to-use-telegram-bot-to-alarm-you-when-someone-login-into-your-vps.html) beziehen.
Legen Sie die roboterbezogenen Parameter im Panel-Hintergrund fest, einschließlich:

- Tg-Roboter-Token
- Tg-Roboter-ChatId
- Laufzeit des Tg-Roboterzyklus, in Crontab-Syntax
- Tg-Roboter-Ablaufschwelle
- Tg-Roboter-Verkehrsschwelle
- Tg-Roboter Aktivieren Sie das Senden von Backups zur Zykluslaufzeit
- Tg robot Aktivieren Sie den Alarmschwellenwert für die CPU-Auslastung

Referenzsyntax:

- 30 \* \* \* \* \* //Benachrichtigen Sie bei den 30ern jedes Punktes
- 0 \*/10 \* \* \* \* //Benachrichtigen Sie alle 10 Minuten in der ersten Sekunde
- @hourly // stündliche Benachrichtigung
- @daily // Tägliche Benachrichtigung (00:00 Uhr morgens)
- @weekly // wöchentliche Benachrichtigung
- @every 8h // alle 8 Stunden benachrichtigen

# Telegram-Bot-Funktionen

- Regelmäßige Berichte
- Anmeldebenachrichtigung
- CPU-Schwellenwertbenachrichtigung
- Schwellenwert für Ablaufzeit und Traffic, der im Voraus gemeldet werden muss
- Unterstützt das Client-Berichtsmenü, wenn der Telegram-Benutzername des Clients zu den Konfigurationen des Benutzers hinzugefügt wird
- Support-Telegramm-Verkehrsbericht, der mit UID (VMESS/VLESS) oder Passwort (TROJAN) durchsucht wird – anonym
- Menübasierter Bot
- Kunden per E-Mail durchsuchen (nur Administrator)
- Überprüfen Sie alle Eingänge
- Überprüfen Sie den Serverstatus
- Überprüfen Sie erschöpfte Benutzer
- Erhalten Sie Backups auf Anfrage und in regelmäßigen Berichten

## API v1-Routen

- `cpsess7945419007/frontend/jupiter/api/v1/login` mit `PUSH` Benutzerdaten: `{username: '', password: ''}` für die Anmeldung
- `/cpless7945419007/frontend/jupiter/api/v1` Basis für folgende Aktionen:


| Method | Path                               | Action                                                                               |
|:------:|------------------------------------|--------------------------------------------------------------------------------------|
| `GET`  | `"/list"`                          | Get all inbounds                                                                     |
| `GET`  | `"/get/:id"`                       | Get inbound with inbound.id                                                          |
| `GET`  | `"/getClientTraffics/:email"`      | Get Client Traffics with email                                                       |
| `POST` | `"/getData"`                       | Get single client usage, not need login ,required paramaters: `email` and `password` |
| `GET`  | `"/createbackup"`                  | Telegram bot sends backup to admins                                                  |
| `POST` | `"/add"`                           | Add inbound                                                                          |
| `POST` | `"/del/:id"`                       | Delete Inbound                                                                       |
| `POST` | `"/update/:id"`                    | Update Inbound                                                                       |
| `POST` | `"/clientIps/:email"`              | Client Ip address                                                                    |
| `POST` | `"/clearClientIps/:email"`         | Clear Client Ip address                                                              |
| `POST` | `"/addClient"`                     | Add Client to inbound                                                                |
| `POST` | `"/:id/delClient/:clientId"`       | Delete Client by clientId\*                                                          |
| `POST` | `"/updateClient/:clientId"`        | Update Client by clientId\*                                                          |
| `POST` | `"/:id/resetClientTraffic/:email"` | Reset Client's Traffic                                                               |
| `POST` | `"/resetAllTraffics"`              | Reset traffics of all inbounds                                                       |
| `POST` | `"/resetAllClientTraffics/:id"`    | Reset traffics of all clients in an inbound                                          |
| `POST` | `"/delDepletedClients/:id"`        | Delete inbound depleted clients (-1: all)                                            |

\*- The field `clientId` should be filled by:

- `client.id` for VMESS and VLESS
- `client.password` for TROJAN
- `client.email` for Shadowsocks

# Panel entwickeln

Sie müssen die Sprache `1.2` eingeben

```bash
git clone https://github.com/4xmen/radepa-x-ui.git 
cd radepa-x-ui.git
go get 
go run main.go
```

Dieses Panel wurde von 3x-UI ohne Pausenänderungen, stabile Version und verstecktes Panel vom GFW-System erweitert.

# Vorschlagssystem

- Ubuntu 20.04+
- Debian 10+
- CentOS 8+
-Fedora 36+

[//]: # (# Pictures)

[//]: # ()
[//]: # (![1]&#40;./media/1.png&#41;)

[//]: # (![2]&#40;./media/2.png&#41;)

[//]: # (![3]&#40;./media/3.png&#41;)

[//]: # (![4]&#40;./media/4.png&#41;)

[//]: # (![5]&#40;./media/5.png&#41;)

[//]: # (![6]&#40;./media/6.png&#41;)

