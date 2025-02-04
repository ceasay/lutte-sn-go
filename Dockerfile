FROM golang:1.16-alpine AS builder

WORKDIR /build

RUN go env -w GO111MODULE=auto

RUN go mod init v1
RUN go mod tidy

RUN go mod download

COPY . .

RUN go build -o video-streaming

EXPOSE 8080

CMD [ "/build/video-streaming" ]