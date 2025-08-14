<p align="center">
  <h2 align="center">Housekeepy</h2>
</p>

<p align="center">
  Never forget about your household responsibilites again
</p>

<div align="center">

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub Release](https://img.shields.io/github/v/release/ashiven/household-planner)](https://github.com/ashiven/household-planner/releases)
[![GitHub Issues or Pull Requests](https://img.shields.io/github/issues/ashiven/household-planner)](https://github.com/ashiven/household-planner/issues)
[![GitHub Issues or Pull Requests](https://img.shields.io/github/issues-pr/ashiven/household-planner)](https://github.com/ashiven/household-planner/pulls)

<img src="./assets/demo.gif"/>
</div>

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Features

- :clipboard: Automatically assigns household tasks
- :alarm_clock: Daily reminders via SMS or Whatsapp
- :wrench: Browser interface for configuration

## Getting Started

### Prerequisites

1. Register for an account on [Twilio](https://www.twilio.com/en-us) and create a [Messaging Service](https://console.twilio.com/us1/develop/sms/services).
2. Have [Docker](https://docs.docker.com/get-started/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed.
3. If you want to host this service, start up your host machine and prepare your domain name.

### Setup

1. Clone the repository to your machine

```bash
git clone https://github.com/ashiven/household-planner.git
```

2. Fill out the `.env` file with the required variables

```bash
TWILIO_ACCOUNT_SID=AC557...
TWILIO_AUTH_TOKEN=25fef...
WHATSAPP_SENDER=+155...
SMS_SENDER=+174...
TEMPLATE_SID=HX8b...
SERVICE_SID=MGf7...
ADMIN_PASSWORD=myAdminPassword
DOMAIN_NAME=myDomainName.com
```

3. Create an SSL certificate

```bash
apt install -y certbot python3-certbot-nginx
```

```bash
certbot certonly --standalone -d <domain-name>
```

4. Start the services

```bash
docker compose up --detach --build
```

## Configuration

- Navigate to your domain name in the browser to open up the configuration interface
- Add members that should be notified about household tasks with phone numbers in the ... format
- Add household tasks that need to be completed every day/week/month
- Once a day at 12:00 PM, every member will be reminded of their tasks via SMS or Whatsapp

## Contributing

Please feel free to submit a [pull request](https://github.com/ashiven/household-planner/pulls) or open an [issue](https://github.com/ashiven/household-planner/issues).

1. Fork the repository
2. Create a new branch: `git checkout -b feature-name`.
3. Make your changes
4. Push your branch: `git push origin feature-name`.
5. Submit a PR

## License

This project is licensed under the [MIT License](./LICENSE).

---

> GitHub [@ashiven](https://github.com/Ashiven) &nbsp;&middot;&nbsp;
> Twitter [ashiven\_](https://twitter.com/ashiven_)
