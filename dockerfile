FROM golang:1.17-alpine

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

# Build files
COPY ./src ./src
RUN go build -o ./build/main ./src/main.go
RUN rm -rf ./src

# Expose port and entrypoint
EXPOSE 3000
CMD [ "./build/main" ]