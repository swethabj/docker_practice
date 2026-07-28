FROM golang:1.25.6
WORKDIR /app                     
COPY . . 
RUN go build -o app main.go      
CMD ["./app"]  