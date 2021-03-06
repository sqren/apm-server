# Node agent testing

This is to setup an environment for the manual testing of the [node agent](https://github.com/elastic/apm-agent-nodejs) with the server.


# Setup

There are two possible setups. If you already have npm / node locally you can run the following to commands to get started:

```
make setup
node app.js
```

There is also a docker environment which comes with all the requisits. Run the follwing command to start the environment:

```
make start
```

## Access node endpoint

In both setups the endpoint can be accessed through `localhost:8081`.


### Docker

If you have the docker setup, you need to change the `apm-server.yml` to access connections not only on localhost but also from remote. Change `localhost:8080` to `:8080` and restart the server.

In `app.js` change `apiHost: 'localhost',` to contain your local IP.
