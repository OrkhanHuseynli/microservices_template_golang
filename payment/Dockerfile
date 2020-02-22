ARG GO_VERSION=1.13.1
# First stage: build the executable.
FROM golang:${GO_VERSION} AS builder

MAINTAINER https://github.com/OrkhanHuseynli

WORKDIR /app
RUN ls
COPY ./ ./
RUN echo " ******** LIST DIR AFTER COPY ******** "
RUN ls
WORKDIR /app/src
RUN echo " ******** LIST DIR IN src ******** "
RUN ls
RUN echo " ******** RUN BUILD ******** "
RUN go build -mod=vendor main.go
RUN echo " ******** LIST DIR AFTER BUILD ******** "
RUN ls
EXPOSE 8080
ENTRYPOINT ["./main"]