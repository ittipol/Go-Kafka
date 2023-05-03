FROM golang:1.19  as build

WORKDIR /go-app

COPY ./account/go.mod ./account/go.sum ./account/
COPY ./migration/ ./migration/

RUN cd account && go mod download

COPY ./account/ ./account/

RUN cd account && go build -o ./app

# =================================================
FROM golang:1.19  as production

WORKDIR /go-app

COPY --from=build /go-app/account/app ./
COPY --from=build /go-app/account/config.yaml ./config.yaml

EXPOSE 5000

ENTRYPOINT ["./app"]
# CMD [ "tail","-F","anyfile" ]