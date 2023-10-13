Applications for sending logs with errors from the LOKI(GRAFANA) logger to the Telegram messenger

The LOKI logger has an inconvenient interface for searching errors in application logs,
and no one will look there every 5 minutes -
Therefore, it is necessary to automate the process of searching for errors in logs and notifying developers.

The application does:
1. Log in to grafana
2. Takes a list of services from the services.txt file
3. For each service, searches for LOKI logs with the text: error:,panic:,ERROR:,PANIC:
4. Found logs with errors are sent to the Telegram messenger
5. Repeats the search every 10 minutes


Installation procedure:
1. Compile this repository
make build
the telegram_loki file will appear in the bin folder


2. make a .env file with filled parameters:
TELEGRAM_APP_ID=
TELEGRAM_APP_HASH=
TELEGRAM_PHONE_FROM=
TELEGRAM_PHONE_SEND_TEST=
LOKI_URL=
GRAFANA_LOGIN=
GRAFANA_PASSWORD=
TELEGRAM_CHAT_NAME=


3. Fill in the settings/services.txt file
In json format write the service name in LOKI
and the name (login) of the programmer (optional)
sample:
{
   "test-service1": "@DeveloperTelegramName"
}


3. Run the telegram_loki file


Source code in Golang language.
Tested on Linux Ubuntu
Readme from 10/13/2023

License:
Save information about the author and the site.

Made by Alexander Nikitin
https://github.com/ManyakRus/telegram_loki