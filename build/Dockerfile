FROM golang:alpine AS build
WORKDIR /webface
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o webface cmd/server.go

FROM golang:alpine
WORKDIR /webface
COPY --from=build /webface/webface .
COPY web ./web
EXPOSE 8080
USER 32000:32000
ENV LDAP_PORT=389
ENTRYPOINT [ "/webface/webface" ]



