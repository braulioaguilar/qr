FROM golang:1.14

LABEL Author="Braulio Aguilar <brauliodev@gmail.com>"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/braulioinf/launcher

COPY . .

RUN go mod download

RUN go get -d -v github.com/markbates/refresh && \
    go install -v github.com/markbates/refresh

CMD ["refresh", "run"]