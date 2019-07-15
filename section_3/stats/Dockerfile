FROM golang:1.12-alpine AS builder
RUN apk add --no-cache ca-certificates git

#ENV PROJECT github.com/Danr17/microservices_project/section_3/stats
#WORKDIR /go/src/$PROJECT/
WORKDIR /microservices_project/section_3/stats
COPY . .
RUN go mod download

RUN go build -ldflags="-w -s" -o /stats main.go

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /stats
COPY --from=builder /stats ./server
EXPOSE 8081
ENTRYPOINT ["/stats/server"]





