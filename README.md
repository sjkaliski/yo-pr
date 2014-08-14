yo-pr
=====

Send a Yo whenever a PR is opened.

# Getting Started

## Requirements

1. Yo Token (http://yoapi.justyo.co/)
2. Github Repo

## Running

```
$ git clone https://github.com/sjkaliski/yo-pr.git
$ go get
$ ./yo-pr --yo_token=<INSERT_TOKEN_HERE>
```

By default, the service runs on port 8080.

## Setup

From your repositories' home page...

Settings > Webhooks & Services > Add webhook

Here you will add a payload url (http://addr:port/pr). Go on to select individual events, then click Pull Request.
