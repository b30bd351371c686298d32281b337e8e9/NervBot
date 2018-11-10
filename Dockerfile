FROM golang:latest
LABEL maintainer="Simon Woldemichael"
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep
WORKDIR /go/src/github.com/swoldemi/NervBot
COPY . .
RUN dep ensure
RUN go build -o main . 
EXPOSE 8080
CMD ["./main"]
