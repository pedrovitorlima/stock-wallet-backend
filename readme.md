# Wallet Manager Backend
## GoLang based backend for a stock wallet manager application (ongoing)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

Wallet Manager is an application to create customized stock wallets given financial goals. It supports you to create different wallets, link them with specific goals, and show how much your investiments appreciated or depreciated. This is the backend codebase.

Its a GoLang based application. You probably are asking "and why is that?". It is how it is because I want to improve my GoLang skills :)

## Features

- POC: Create wallets and add goals (ongoing)
    - Basic project and dependencies [done]
    - Persist in DB [done]
    - APIs and Tests [ongoing]
    - Refactoring [ongoing]
- Link stocks by ticker (near future)
- Get stock appreciations through a third-party API (near future)
- Reports (near future)


## Tech

A few technologies being used:

- [GoLang] - Language and plugins
- [Docker] - There is a docker-compose file to create the infrastructure (database, for example)
- [Postgres] - Database being used to persist data
- [Dillinger] - Markdown editor to create this file

And of course the wallet-manager-backend itself is open source with a [public repository][dill]
 on GitHub.

## Installation



## Plugins


## Development


## Docker

Currently you just need to use docker compose to run the database. Do

```sh
docker-compose up
```

in the root folder of this project and thats it.

