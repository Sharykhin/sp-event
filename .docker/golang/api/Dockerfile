FROM golang:1.12


RUN go get github.com/canthefason/go-watcher && \
    go install github.com/canthefason/go-watcher/cmd/watcher && \
    go get -u gopkg.in/alecthomas/gometalinter.v2 && \
    gometalinter.v2 --install && \
    go get -u github.com/pressly/goose/cmd/goose

ADD . /sp-event

WORKDIR /sp-event

#CMD exec /bin/sh -c "trap : TERM INT; (while true; do sleep 1000; done) & wait"

# watch uses /go/src as a default directory
CMD watcher -run ./cmd/api -watch ../../sp-event/
