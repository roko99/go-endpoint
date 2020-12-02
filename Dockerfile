FROM golang:alpine as golang
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'

FROM scratch
COPY --from=golang /go/bin/app /app
EXPOSE 5003
ENTRYPOINT ["/app"]