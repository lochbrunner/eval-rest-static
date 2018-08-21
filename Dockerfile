FROM golang

ADD . /go/src/github.com/lochbrunner/eval-rest-static

RUN go get github.com/gorilla/mux && \
    go get -u github.com/gobuffalo/packr/...

RUN packr install github.com/lochbrunner/eval-rest-static

ENTRYPOINT /go/bin/eval-rest-static

EXPOSE 3001