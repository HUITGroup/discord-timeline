FROM golang:alpine
RUN apk update && apk add git
RUN mkdir /dctimeline
WORKDIR /dctimeline
ADD . /dctimeline
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /dctimeline/app .
CMD ["./app"]