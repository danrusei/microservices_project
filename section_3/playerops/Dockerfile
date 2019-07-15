FROM golang:1.12-alpine AS builder
RUN apk add --no-cache ca-certificates git

#ENV PROJECT github.com/Danr17/microservices_project/section_3/playerops
#WORKDIR /go/src/$PROJECT/
WORKDIR /microservices_project/section_3/playerops
COPY . .
RUN go mod download

RUN go build -ldflags="-w -s" -o /playerops main.go

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /playerops
COPY --from=builder /playerops ./server
EXPOSE 8082
ENTRYPOINT ["/playerops/server"]





