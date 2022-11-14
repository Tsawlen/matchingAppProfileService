# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app/matchingAppProfileService

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /docker-app-profile-service 

EXPOSE 8080

CMD [ "/docker-app-profile-service" ]