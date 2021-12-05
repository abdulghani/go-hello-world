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
RUN CGO_ENABLED=0 go build -o ./go_app_binary ./src/main.go

##
# DEPLOY
FROM gcr.io/distroless/static

WORKDIR /usr/src/app

# sometimes had problems with using relative . path
COPY --from=build /usr/src/app/go_app_binary ./
COPY ./.env ./
COPY ./jwt_key.example ./

EXPOSE 3000

CMD [ "./go_app_binary" ]