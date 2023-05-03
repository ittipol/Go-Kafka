FROM golang:1.19  as build

WORKDIR /go-app

COPY ./producer/go.mod ./producer/go.sum ./producer/
COPY ./events/ ./events/

RUN cd producer && go mod download

COPY ./producer/ ./producer/

RUN cd producer && go build -o ./app

# =================================================
FROM golang:1.19  as production

WORKDIR /go-app

COPY --from=build /go-app/producer/app ./
COPY --from=build /go-app/producer/config.yaml ./config.yaml

EXPOSE 4000

ENTRYPOINT ["./app"]
# CMD [ "tail","-F","anyfile" ]