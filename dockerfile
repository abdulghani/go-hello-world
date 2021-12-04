FROM golang:alpine

# Create app directory
WORKDIR /app

# Install app dependencies
COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./src ./src
RUN go build -o ./build/main ./src/main.go

EXPOSE 8080
CMD [ "./build/main" ]