FROM golang:1.14.3-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go build -o /out/example .

FROM scratch AS run
COPY --from=build /out/example /
ENTRYPOINT ["./example"]
