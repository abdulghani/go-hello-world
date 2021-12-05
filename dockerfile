##
# BUILD
FROM golang:1.17-buster AS build

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

# Build files
COPY ./src ./src
RUN go build -o ./go_app_binary ./src/main.go

##
# DEPLOY
FROM gcr.io/distroless/base

WORKDIR /usr/src/app

# sometimes had problems with using relative . path
COPY --from=build /usr/src/app/go_app_binary /usr/src/app/go_app_binary

EXPOSE 3000

# same. had problems with relative . path
ENTRYPOINT [ "/usr/src/app/go_app_binary" ]