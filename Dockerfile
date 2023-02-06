FROM golang:1.19.5 as build

WORKDIR /dockerdev

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go build -gcflags="all=-N -l" -o /server

FROM debian:buster

EXPOSE 3200
EXPOSE 40000

WORKDIR /

COPY --from=build /dockerdev/.env /
COPY --from=build /go/bin/dlv /
COPY --from=build /server /

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]
#CMD ["/server"]