Applications for sending error logs from the LOKI logger (or Victoria Metrics) to the Telegram messenger

Grafana has an inconvenient web interface for searching for errors in application logs,
and no one will look there every 5 minutes -
so we need to automate the process of searching for errors in logs and notifying developers.

The application does:
Search for errors in the LOKI logger (or in Victoria Metrics)
1. Logs into the LOKI (Victoria) api
2. Takes a list of services from the services.txt file,
including the service name and the name of the person responsible
3. For each service, searches for logs with the text: error:, panic:, ERROR:, PANIC:
4. Sends the found logs with errors to the Telegram messenger
5. Repeats the search every 10 minutes

Search for errors in the database:
1. Logs into the PostgreSQL database
2. Finds .sql files in the sql folder, and executes a script in the database,
if the script returns a value, a message about this will come to the Telegram messenger,
if the script does not return a single line, then everything is fine.
3. Repeats requests every 60 minutes
4. If you fill in the scripts.txt file (scripts_add.txt),
including the .sql file name and the name of the person responsible,
then it will also add the name of the person responsible.

The bot sends error logs to Telegram groups:
1. General group with all errors, all programmers.
The group ID is filled in settings.txt, TELEGRAM_CHAT_NAME=
In the Telegram group settings, you need to set the rights Group type = "Public"
2. Each programmer separately, only the logs that apply to him.
Filled in the files services.txt, scripts.txt
Each programmer must add this telegram bot to Telegram

Installation procedure:
1. Compile this repository
>make build
the telegram_loki file will appear in the bin folder

2. fill in the settings.txt (or .env) file
with the LOKI URL settings, login, password, etc.,
as well as the database connection settings (optional).

3. Fill in the settings/services.txt file
In json format, write the service name in LOKI
and the name (login) of the programmer (optional)
sample:
{
"test-service1": "@DeveloperTelegramName"
}

3. Run the telegram_loki file

Source code in Golang.
Tested on Linux Ubuntu
Readme from 08/19/2025

License:
GPL 2, save information about the author and the site.

Made by Alexander Nikitin
https://github.com/ManyakRus/telegram_loki