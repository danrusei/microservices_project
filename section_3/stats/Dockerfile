FROM golang:1.12-alpine AS builder
RUN apk add --no-cache ca-certificates git

ENV PROJECT github.com/Danr17/microservices_project/section_3/stats
WORKDIR /go/src/$PROJECT

COPY go.mod go.sum ./
COPY . .
RUN go build -o /stats .

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /stats
COPY --from=builder /stats ./server
EXPOSE 8081
ENTRYPOINT ["/stats/server"]





