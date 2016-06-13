OKCoin Api - at the moment development is focused on the RESTful API provided by OKCoin and will eventually move to support websockets.

Must contain two environment variables to test: OKCOIN_API_KEY and OKCOIN_SECRET_KEY.

Add them to your .bash_profile or .bashrc file in the following way:

    export OKCOIN_API_KEY=your okcoin api key goes here
    export OKCOIN_SECRET_KEY=your okcoin secret key goes here


View the OKCoin REST API here: https://www.okcoin.com/about/rest_api.do


A little TL;DR:

OKCoin's exchange is responsible for a lot of traffic that happens on today's bitcoin blockchain. This project is to help developers (and my self) use OKCoin's API to commit trades/monitor changes on their exchange.

**Why Golang:** New language, good future, fast, learning project, (did I mention fast?). As this repo progresses, so is my hope for the use of Golang in bitcoin trades.

**Want to help?** Yes, please! Make a PR, criticize what I'm doing, I'm open to any and all suggestions and any help I can get. Just branch off master (I should start doing that too), do some stuff, and create a PR with master as its target.                                                
