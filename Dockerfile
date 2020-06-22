FROM golang:1.14-alpine as build
WORKDIR /src/translate
COPY . .
RUN go build

FROM alpine:latest
COPY --from=build /src/translate .
CMD ["./translate", "-port", "3000"]
EXPOSE 3000
