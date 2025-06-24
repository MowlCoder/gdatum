FROM golang:1.24 AS build-stage

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 GOOS=linux go build -o ./app cmd/app/main.go

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build-stage --chown=nonroot:nonroot /build/app ./app

EXPOSE 8080
EXPOSE 8081

USER 65532:65532

ENTRYPOINT ["./app"]
