FROM golang:1.12-alpine AS builder
RUN apk add --no-cache ca-certificates git

ENV PROJECT github.com/Danr17/microservices_project/section_3/playerops
WORKDIR /go/src/$PROJECT

COPY go.mod go.sum ./
COPY . .
RUN go build -o /playerops .

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /playerops
COPY --from=builder /playerops ./server
EXPOSE 8082
ENTRYPOINT ["/playerops/server"]





