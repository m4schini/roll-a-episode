FROM docker.io/library/golang:buster as build

RUN mkdir /build

WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o tv .

FROM docker.io/library/golang:1.20 as prod

COPY --from=build /build/tv /tv

ENV PORT=8080

EXPOSE $PORT

ENTRYPOINT ["/tv"]