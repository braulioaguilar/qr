FROM golang:1.14

LABEL Author="Braulio Aguilar <brauliodev@gmail.com>"

WORKDIR /go/src/app

RUN go get -d -v github.com/markbates/refresh && \
    go install -v github.com/markbates/refresh

CMD ["refresh", "run"]