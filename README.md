# [Corona test](coronatest.nl) result checker

## Environment variables
- TELEGRAM_BOT_TOKEN: The token you got from the [BotFather](https://core.telegram.org/bots#6-botfather).
- TELEGRAM_CHAT_ID: Your chat ID with the bot. If you send <code>/id</code> to the bot, it will respond with your chat ID.
- BEARER: The Authentication you got from coronatest.nl to get your test results.

### Online / offline
The telegram bot will also send you a message if the server starts up (<code>I'm back online!</code>), or shuts down (<code>Going offline...</code>). This message should always be sent, even if a part of the program crashed. Note that if the Telegram bot crashes, you will obviously not get a message.<br/>
By sending those messages, you will always be notified about the current status of the server.
(Keep in mind that you can always access the corresponding website on Heroku to get your server back online.)