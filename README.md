# jira-board-merger
A tool to make tasks from multiple jira server appear in a single confluence board

```
cp config.tpl.json config.json
docker build . -t jira-board-merger
docker run --rm --name jira-board-merger -v ~/Projects/jira-board-merger/config.json:/go/src/jira_merger/config.json -p 8080:8080 -d jira-board-merger
curl -X GET localhost:8080
```

# Configuration:
You must put the additional servers to query in the configuration.

# Dev to stat the debug API
the api simulate a confluence server answer.
```
docker run --name jira-board-merger-dev-api --restart=always -p 2564:80 -v ~/Projects/jira-board-merger/dev_api/:/app -d jazzdd/alpine-flask
```

you can then configure the merger with the following:
```
{
  "servers": [
    {"host": "http://<Your Ip>:2564"},
    {"host": "http://<Your Ip>:2564"}
  ]
}

```
