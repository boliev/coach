FROM golang:1.16.1-alpine
WORKDIR /app
COPY . /app
RUN go build -o /bin/coach cmd/main.go 
#src/config.go src/wire_gen.go