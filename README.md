# [Corona test](coronatest.nl) result checker

## Environment variables
- TELEGRAM_BOT_TOKEN: The token you got from the [BotFather](https://core.telegram.org/bots#6-botfather).
- TELEGRAM_CHAT_ID: Your chat ID with the bot. If you send <code>/id</code> to the bot, it will respond with your chat ID.
- BEARER: The Authentication you got from coronatest.nl to get your test results.

## Heroku
To run this code on one of the [Heroku servers](https://www.heroku.com/), we listen on a specific port, so it will act like a webserver. However with the free approach on Heroku, your app will go offline if it isn't accessed within an hour. Therefore [kaffeine](https://kaffeine.herokuapp.com/) is used to keep this app online.

### Online / offline
The telegram bot will also send you a message if the server starts up (<code>I'm back online!</code>), or shuts down (<code>Going offline...</code>). This message should always be sent, even if a part of the program crashed. Note that if the Telegram bot crashes, you will obviously not get a message.<br/>
By sending those messages, you will always be notified about the current status of the server.
(Keep in mind that you can always access the corresponding website on Heroku to get your server back online.)