FROM golang:1.24
ENV GOPROXY=https://goproxy.io,direct
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o fizzBuzz
EXPOSE 8080
CMD ["./fizzBuzz"]