ARG VERSION

FROM bitnami/golang AS base
RUN mkdir -p -m 0600 ~/.ssh 
WORKDIR /opt/app

COPY . .
RUN go install github.com/cosmtrek/air@latest

FROM base as dev
ENTRYPOINT ["air", "-c", "./build/config/.air.toml"]

FROM base as prod
# ENTRYPOINT ["make", "run"]