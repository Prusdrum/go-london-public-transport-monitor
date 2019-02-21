FROM golang:1.11

WORKDIR /go/src/app
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN make install
RUN make build
CMD ["./main"]