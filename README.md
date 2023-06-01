# radepa-x-ui
![4](./web/assets/svg/footprint.svg)

Hide your footprint with Radepa

> **Disclaimer: This project is only for personal learning and communication, please do not use it for illegal purposes, please do not use it in a production environment**


[![License](https://img.shields.io/badge/license-GPL%20V3-blue.svg?longCache=true)](https://www.gnu.org/licenses/gpl-3.0.en.html)

radepa-x-ui panel supporting multi-protocol, **Multi-lang (English,Farsi,Chinese,Russian)**

**If you think this project is helpful to you, you may wish to give a** :star2:


# Install & Upgrade

```
bash <(curl -Ls https://raw.githubusercontent.com/4xmen/radepa-x-ui/master/install.sh)
```

[//]: # (## Install custom version)

[//]: # ()
[//]: # (To install your desired version you can add the version to the end of install command. Example for ver `v1.6.0`:)

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

or you can use x-ui menu then number '16' (Apply for an SSL Certificate)

# Default settings

- Port: 2086
- username and password will be generated randomly if you skip to modify your own security(x-ui "7")
- database path: /etc/x-ui/x-ui.db
- xray config path: /usr/local/x-ui/bin/config.json

Before you set ssl on settings

- http://ip:2086/
- http://domain:2086/

After you set ssl on settings

- https://yourdomain:2086/

# Environment Variables

| Variable       |                      Type                      | Default       |
| -------------- | :--------------------------------------------: | :------------ |
| XUI_LOG_LEVEL  | `"debug"` \| `"info"` \| `"warn"` \| `"error"` | `"info"`      |
| XUI_DEBUG      |                   `boolean`                    | `false`       |
| XUI_BIN_FOLDER |                    `string`                    | `"bin"`       |
| XUI_DB_FOLDER  |                    `string`                    | `"/etc/x-ui"` |

Example:

```sh
XUI_BIN_FOLDER="bin" XUI_DB_FOLDER="/etc/x-ui" go build main.go
```

# Xray Configurations:

**copy and paste to xray Configuration :** (you don't need to do this if you have a fresh install)

- [traffic](./media/configs/traffic.json)
- [traffic + Block all Iran IP address](./media/configs/traffic+block-iran-ip.json)
- [traffic + Block all Iran Domains](./media/configs/traffic+block-iran-domains.json)
- [traffic + Block Ads + Use IPv4 for Google](./media/configs/traffic+block-ads+ipv4-google.json)
- [traffic + Block Ads + Route Google + Netflix + Spotify + OpenAI (ChatGPT) to WARP](./media/configs/traffic+block-ads+warp.json)

# [WARP Configuration](https://github.com/fscarmen/warp) (Optional)

If you want to use routing to WARP follow steps as below:

1. If you already installed warp, you can uninstall using below command:

   ```sh
   warp u
   ```

2. Install WARP on **socks proxy mode**:

   ```sh
   curl -fsSL https://gist.githubusercontent.com/hamid-gh98/dc5dd9b0cc5b0412af927b1ccdb294c7/raw/install_warp_proxy.sh | bash
   ```

3. Turn on the config you need in panel or [Copy and paste this file to Xray Configuration](./media/configs/traffic+block-ads+warp.json)

   Config Features:

   - Block Ads
   - Route Google + Netflix + Spotify + OpenAI (ChatGPT) to WARP
   - Fix Google 403 error

# Features

- System Status Monitoring
- Search within all inbounds and clients
- Support Dark/Light theme UI
- Support multi-user multi-protocol, web page visualization operation
- Supported protocols: vmess, vless, trojan, shadowsocks, dokodemo-door, socks, http
- Support for configuring more transport configurations
- Traffic statistics, limit traffic, limit expiration time
- Customizable xray configuration templates
- Support https access panel (self-provided domain name + ssl certificate)
- Support one-click SSL certificate application and automatic renewal
- For more advanced configuration items, please refer to the panel
- Support to change configs by different items provided in panel
- Support export/import database from panel
- Stable APIs
- Hide your panel form GFW

# Tg robot use

X-UI supports daily traffic notification, panel login reminder and other functions through the Tg robot. To use the Tg robot, you need to apply for the specific application tutorial. You can refer to the [blog](https://coderfan.net/how-to-use-telegram-bot-to-alarm-you-when-someone-login-into-your-vps.html)
Set the robot-related parameters in the panel background, including:

- Tg robot Token
- Tg robot ChatId
- Tg robot cycle runtime, in crontab syntax
- Tg robot Expiration threshold
- Tg robot Traffic threshold
- Tg robot Enable send backup in cycle runtime
- Tg robot Enable CPU usage alarm threshold

Reference syntax:

- 30 \* \* \* \* \* //Notify at the 30s of each point
- 0 \*/10 \* \* \* \* //Notify at the first second of each 10 minutes
- @hourly // hourly notification
- @daily // Daily notification (00:00 in the morning)
- @weekly // weekly notification
- @every 8h // notify every 8 hours

# Telegram Bot Features

- Report periodic
- Login notification
- CPU threshold notification
- Threshold for Expiration time and Traffic to report in advance
- Support client report menu if client's telegram username added to the user's configurations
- Support telegram traffic report searched with UID (VMESS/VLESS) or Password (TROJAN) - anonymously
- Menu based bot
- Search client by email ( only admin )
- Check all inbounds
- Check server status
- Check depleted users
- Receive backup by request and in periodic reports

## API v1 routes

- `cpsess7945419007/frontend/jupiter/api/v1/login` with `PUSH` user data: `{username: '', password: ''}` for login
- `/cpsess7945419007/frontend/jupiter/api/v1` base for following actions:

| Method | Path                        | Action                                                                               |
|:------:|-----------------------------|--------------------------------------------------------------------------------------|
| `GET`  | `"/list"`                   | Get all inbounds                                                                     |
| `GET`  | `"/get/:id"`                | Get inbound with inbound.id                                                          |
| `GET`  | `"/getClientTraffics/:email"` | Get Client Traffics with email                                                       |
| `POST` | `"/getDate"`                | Get single client usage, not need login ,required paramaters: `email` and `password` |
| `GET`  | `"/createbackup"`           | Telegram bot sends backup to admins                                                  |
| `POST` | `"/add"`                    | Add inbound                                                                          |
| `POST` | `"/del/:id"`                | Delete Inbound                                                                       |
| `POST` | `"/update/:id"`             | Update Inbound                                                                       |
| `POST` | `"/clientIps/:email"`       | Client Ip address                                                                    |
| `POST` | `"/clearClientIps/:email"`  | Clear Client Ip address                                                              |
| `POST` | `"/addClient"`              | Add Client to inbound                                                                |
| `POST` | `"/:id/delClient/:clientId"` | Delete Client by clientId\*                                                          |
| `POST` | `"/updateClient/:clientId"` | Update Client by clientId\*                                                          |
| `POST` | `"/:id/resetClientTraffic/:email"` | Reset Client's Traffic                                                               |
| `POST` | `"/resetAllTraffics"`       | Reset traffics of all inbounds                                                       |
| `POST` | `"/resetAllClientTraffics/:id"` | Reset traffics of all clients in an inbound                                          |
| `POST` | `"/delDepletedClients/:id"` | Delete inbound depleted clients (-1: all)                                            |

\*- The field `clientId` should be filled by:

- `client.id` for VMESS and VLESS
- `client.password` for TROJAN
- `client.email` for Shadowsocks

# develop panel

You need go lang `1.2`

```bash
git clone https://github.com/4xmen/radepa-x-ui.git 
cd radepa-x-ui.git
go get 
go run main.go
```

This panel extended from 3x-ui with no break changes, stable release and hidden panel from GFW system. 

# Suggestion System

- Ubuntu 20.04+
- Debian 10+
- CentOS 8+
- Fedora 36+

[//]: # (# Pictures)

[//]: # ()
[//]: # (![1]&#40;./media/1.png&#41;)

[//]: # (![2]&#40;./media/2.png&#41;)

[//]: # (![3]&#40;./media/3.png&#41;)

[//]: # (![4]&#40;./media/4.png&#41;)

[//]: # (![5]&#40;./media/5.png&#41;)

[//]: # (![6]&#40;./media/6.png&#41;)

