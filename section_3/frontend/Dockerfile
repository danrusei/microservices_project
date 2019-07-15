FROM golang:1.12-alpine AS builder
RUN apk add --no-cache ca-certificates git

ENV PROJECT github.com/Danr17/microservices_project/section_3/frontend
WORKDIR /go/src/$PROJECT

COPY go.mod go.sum ./
COPY . .
RUN go build -o /frontend .

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /frontend
COPY --from=builder /frontend ./server
EXPOSE 8082
ENTRYPOINT ["/frontend/server"]





