FROM golang:1.23 as local

WORKDIR /go/src/github.com/famiphoto/famiphoto/api
COPY ../../api .
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN mkdir -p /var/www/famiphoto
EXPOSE 8080

FROM golang:1.23 as builder
WORKDIR /go/src/github.com/famiphoto/famiphoto/api
COPY ../../api .
RUN make build

FROM golang:1.23
ARG UID
ARG GID
ARG USERNAME
RUN groupadd -g ${GID} ${USERNAME}
RUN useradd -u ${UID} -g ${USERNAME} -m ${USERNAME}
WORKDIR /go/src/github.com/famiphoto/famiphoto/api
COPY --from=builder /go/src/github.com/hiroyky/famiphoto/dst/api ./api
