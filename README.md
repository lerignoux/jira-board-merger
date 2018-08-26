# jira-board-merger
A tool to make tasks from multiple jira server appear in a single confluence board

```
cp config.tpl.json config.json
docker build . -t jira-board-merger
docker run --name jira-board-merger -v config.json:/go/src/app/config.json -p 8080:8080 -d jira-board-merger
```

# Configuration:
You must put the additional servers to query in the configuration.
