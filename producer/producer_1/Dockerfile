FROM golang:1.19-alpine

# set working directory
WORKDIR /src
# copy all directory
ADD . .
# tun applicaiton
RUN go install main.go
# build the binary
ENTRYPOINT [ "main" ]