FROM bitnami/golang AS base
WORKDIR /opt/app

COPY . .

FROM base AS dev
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT ["air", "-c", "./build/config/.air.toml"]

FROM base AS prod
ENTRYPOINT ["make", "install"]