FROM golang:1.19.13-alpine as build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build
RUN ls -alrt

# Runner image
FROM alpine
WORKDIR /
COPY --from=build /build/wheresthejokebot .
CMD ./wheresthejokebot -token $DISCORD_TOKEN