# Clexicon
Customizable lexicon webserver container


## Clexicon-API
This is the back-end which stores the data and provides it via a ReST API.
It can be started either by cloning it to a linux system, cd-ing to the `api`
directory, and running

    make run

Or by running the following on a server with docker, filling in the appropriate information:

    docker run -d --name <container-name> -p <host-port>:34573 bismarck6502/clexicon_api:latest

Once it's up, you can open the root document in a web-browser to see the
documentation.

## Clexicon-Web
The front-end is in progress...
