FROM golang:latest AS build
WORKDIR /go/src
COPY app .

ENV CGO_ENABLED=0
RUN go install
RUN go build -a -installsuffix cgo -o main .

FROM scratch AS runtime
COPY --from=build /go/src/main ./
EXPOSE 8080/tcp
ENTRYPOINT ["./main"]
