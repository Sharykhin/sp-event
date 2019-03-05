FROM golang:1.12

ADD . /go/src/github.com/Sharykhin/sp-event

WORKDIR /go/src/github.com/Sharykhin/sp-event


RUN go get github.com/canthefason/go-watcher && \
    go install github.com/canthefason/go-watcher/cmd/watcher && \
    go get -u gopkg.in/alecthomas/gometalinter.v2 && \
    gometalinter.v2 --install && \
    go get -u github.com/pressly/goose/cmd/goose

EXPOSE 8003

CMD watcher -run github.com/Sharykhin/sp-event/cmd/api -watch github.com/Sharykhin/sp-event
