FROM golang:alpine
LABEL maintainer="Simon Woldemichael"
RUN apk update && apk add git && apk add ca-certificates
RUN adduser -D -g '' nervbot
WORKDIR $GOPATH/src/github.com/swoldemi/NervBot
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o $GOPATH/bin/NervBot
CMD ["NervBot"]
