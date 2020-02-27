FROM golang:alpine
RUN mkdir /app
COPY goendpoint.go /app

RUN cd /app/ && go build /app/goendpoint.go
EXPOSE 5003
CMD ["/app/goendpoint"]