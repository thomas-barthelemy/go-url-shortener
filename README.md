GoLang URL Shortener Sample with Redis storage
======================================

This is a Sample URL Shortener made for learning purposes,
it uses GoLang, [Gin Gonic](https://gin-gonic.github.io/gin/) web framework
and requires a redis server.

I used a docker redis container for development, which is one of many possibilities.
It has the advantage to be quick to setup:

    docker run --name url_short_db -d -p 6379:6379 redis redis-server --appendonly yes

If you are not comfortable with docker I invite you to check [redis website](redis.io)
for more information on how to set up one.

Url Shortener Logic
-------------------

Each URL is given a unique id (base 10) when created
which is in turn converted to another base depending on the character set available.
It contains `[a-zA-Z0-9]` by default so the ids are converted to base 62.

When a url is requested (`baseUrl/shortUrlCode`) the `shortUrlCode` is translated
back to the URL id and the full URL is retrieved.

The response is then made a redirect to that URL.

Endpoints
---------

    [GIN-debug] GET    /:code                    --> github.com/thomas-barthelemy/go-url-shortener/modules/urlShortener.getUrlEndpoint (1 handlers)
    [GIN-debug] POST   /api/urls                 --> github.com/thomas-barthelemy/go-url-shortener/modules/urlShortener.postUrlEndpoint (1 handlers)

The `POST /api/urls` requires a form encoded parameter `url`. 

What else?
----------

Things that could be added are legion, here are a few improvements:

 - Improving how the id is generated/stored in redis, there is a risk of overriding an existing URL
 - Using redis key expiration for the URLs
 - Adding the CHAR MAP as a config instead of being hardcoded
 - More error handling, mostly ignored so far

In that sense, feel free to PR!