FROM golang
RUN mkdir /pro
ADD ./gConnect.go /pro/
ADD ./credentials.json /pro/
ADD ./token.json /pro/
WORKDIR /pro
RUN go get -d -v ./...
RUN go build -o server gConnect.go
CMD ["/pro/server"]
