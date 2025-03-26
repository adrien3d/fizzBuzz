FROM golang:1.24
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o fizzBuzz
EXPOSE 8080
CMD ["./fizzBuzz"]