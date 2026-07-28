FROM golang:1.25.6
WORKDIR /app                    # docker folder path to begin , all commands run from this path as entrypoint               
COPY . .                        # copy all from source to destination
RUN go build -o app main.go     # in app folder of docker(dockerfile) main.go file is present 
CMD ["./app"]                   # ./app == ./main.go 