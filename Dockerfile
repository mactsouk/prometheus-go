# docker build -t pro-gmail .
# docker run -it -p 2345:2345 pro-gmail

FROM golang
RUN mkdir /pro
ADD ./gConnect.go /pro/
ADD ./credentials.json /pro/
ADD ./token.json /pro/
WORKDIR /pro
EXPOSE 2345
RUN go get -d -v ./...
RUN go build -o server gConnect.go
CMD ["/pro/server"]
