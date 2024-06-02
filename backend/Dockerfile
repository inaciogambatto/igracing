FROM golang:1.22.3 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o server
#RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags "-w -extldflags -static" -o server

ENTRYPOINT [ "/app/server" ]

#FROM scratch
#COPY --from=builder /app/server /server
#ENTRYPOINT [ "server" ]