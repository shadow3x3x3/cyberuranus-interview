FROM golang:latest
WORKDIR /cu
COPY . /cu
RUN go build -o app

EXPOSE 8080
ENTRYPOINT [ "./app" ]