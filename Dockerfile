# docker build -t pro-gmail .
# docker run -it -p 2345:2345 pro-gmail

FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /pro
ADD ./gConnect.go /pro/
ADD ./credentials.json /pro/
ADD ./token.json /pro/
WORKDIR /pro
EXPOSE 2345
RUN go get -d -v ./...
RUN go build -o server gConnect.go

FROM scratch

RUN mkdir /pro
COPY --from=builder /pro/server /pro/server

CMD ["/pro/server"]
