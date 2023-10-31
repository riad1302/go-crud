#  latest golang image. builder is alias
FROM golang:alpine as builder

LABEL maintainer="Habibur Rahman Riad <habibur.genex@gmail.com>"

#set working directory on container
WORKDIR /var/www/html/
# copy go.mod go.sum files from local to container
COPY --chown=www-data:www-data ./codes/go.mod ./codes/go.sum /var/www/html/

# downloads all the dependencies
RUN go mod download
# copy source from local to working directory on container
COPY --chown=www-data:www-data ./codes /var/www/html/
# build the app. After building the app executable is stored in main ( -o main)
# go help build will give more details about each parameters
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#next stage#
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /var/www/html/

# copy the pre-built binary from 1st/previous stage
COPY --from=builder /var/www/html/main .

# run the executable. User either CMD or ENTRYPOINT
ENTRYPOINT ["./main"]

RUN go build -o /golang-binary

EXPOSE 8080

CMD [ "/golang-binary" ]





#  latest golang image. builder is alias
FROM golang:1.21.1-alpine  as builder

LABEL maintainer="Habibur Rahman Riad <habibur.genex@gmail.com>"

#set working directory on container
WORKDIR /var/www/html/
# copy go.mod go.sum files from local to container
COPY --chown=www-data:www-data ./codes/go.mod ./codes/go.sum /var/www/html/

# downloads all the dependencies
RUN go mod download
# copy source from local to working directory on container
COPY --chown=www-data:www-data ./codes /var/www/html/

RUN go build -o /golang-binary

EXPOSE 8080

CMD [ "/golang-binary" ]