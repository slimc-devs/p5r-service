FROM golang:1.15 as builder

WORKDIR /app
COPY . .
#RUN go test
RUN CGO_ENABLED=0 GOOS=linux go install cmd/*


FROM scratch
COPY --from=builder /go/bin/app /usr/local/bin/
CMD ["app"]
