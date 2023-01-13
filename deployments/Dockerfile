# build the binary
FROM golang:1.19 as build

RUN useradd -u 10001 gopher

ENV NAME "service"
WORKDIR /opt/${NAME}

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./bin/${NAME} ./cmd/main.go

# run the binary
FROM scratch

ENV NAME "service"

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd

USER gopher

COPY --from=build /opt/${NAME}/bin/${NAME} /${NAME}

CMD ["./service"]