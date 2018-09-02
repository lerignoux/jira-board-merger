# jira-board-merger
A tool to make tasks from multiple jira server appear in a single confluence board

```
cp config.tpl.json config.json
docker build . -t jira-board-merger
docker run --rm --name jira-board-merger -v ~/Projects/jira-board-merger/config.json:/go/src/jira_merger/config.json -p 8080:8080 -d jira-board-merger
docker run --name jira-merger-nginx -v ~/Projects/jira-board-merger/nginx.conf:/etc/nginx/nginx.conf:ro -d nginx
```

# Configuration:
You must put the additional servers to query in the configuration.

edit the nginx configuration and mount it in it's container:
```
~/Projects/jira-board-merger/nginx.conf:/etc/nginx/nginx.conf:ro -d nginx
```

# Usage
go to nginx url to access confluence merged data.

Data is readonly

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
