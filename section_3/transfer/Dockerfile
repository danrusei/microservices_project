FROM golang:1.12-alpine AS builder
RUN apk add --no-cache ca-certificates git

#ENV PROJECT github.com/Danr17/microservices_project/section_3/transfer
#WORKDIR /go/src/$PROJECT/
WORKDIR /microservices_project/section_3/transfer
COPY . .
RUN go mod download

RUN go build -ldflags="-w -s" -o /transfer main.go

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /transfer
COPY --from=builder /transfer ./server
EXPOSE 8083
ENTRYPOINT ["/transfer/server"]





