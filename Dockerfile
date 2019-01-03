FROM golang as builder

ADD . /go/src/github.com/lochbrunner/eval-rest-static

RUN go get github.com/gorilla/mux && \
    go get -u github.com/gobuffalo/packr/...

RUN CGO_ENABLED=0 GOOS=linux \
    packr install -a -installsuffix cgo github.com/lochbrunner/eval-rest-static


FROM scratch

COPY --from=builder /go/bin/eval-rest-static /
ENTRYPOINT ["/eval-rest-static"]

EXPOSE 3001
