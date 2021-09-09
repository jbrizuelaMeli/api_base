FROM golang:alpine AS build

WORKDIR /go/src/github.com/api_base
COPY . .
RUN go build -o /go/bin/api_base cmd/main.go

FROM scratch
COPY --from=build /go/bin/api_base /go/bin/api_base
ENTRYPOINT ["/go/bin/api_base"]