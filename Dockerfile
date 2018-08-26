FROM golang:alpine3.7
MAINTAINER Laurent Erignoux lerignoux@gmail.com

WORKDIR /go/src
COPY . .

RUN go get -d -v ./jira_merger
RUN go install ./jira_merger

CMD ["jira_merger"]
