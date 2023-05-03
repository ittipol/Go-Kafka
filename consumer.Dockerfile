FROM golang:1.19  as build

WORKDIR /go-app

COPY ./consumer/go.mod ./consumer/go.sum ./consumer/
COPY ./events/ ./events/
COPY ./migration/ ./migration/

RUN cd consumer && go mod download

COPY ./consumer/ ./consumer/

RUN cd consumer && go build -o ./app

# =================================================
FROM golang:1.19  as production

WORKDIR /go-app

COPY --from=build /go-app/consumer/app ./
COPY --from=build /go-app/consumer/config.yaml ./config.yaml

# EXPOSE 4040

ENTRYPOINT ["./app"]
# CMD [ "tail","-F","anyf" ]