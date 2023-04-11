FROM golang:1.18-alpine
WORKDIR /dev_repo
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . . 
RUN go build -o ./out/dist .
CMD ./out/dist
