# FROM golang:1.20-bullseye

# COPY ./compose/local/entrypoint /entrypoint
# RUN sed -i 's/\r$//g' /entrypoint
# RUN chmod +x /entrypoint

# COPY ./compose/local/start /start
# RUN sed -i 's/\r$//g' /start
# RUN chmod +x /start

# WORKDIR /build

# COPY go.mod go.sum main.go ./
# COPY app ./app
# COPY pkg ./pkg
# COPY cmd ./cmd
# COPY platform ./platform

# RUN go mod download && \
#   go install github.com/cosmtrek/air@latest && \
#   go install github.com/swaggo/swag/cmd/swag@latest

# ENTRYPOINT ["/entrypoint"]
FROM golang:1.23-bullseye

COPY ./compose/local/entrypoint /entrypoint
RUN sed -i 's/\r$//g' /entrypoint
RUN chmod +x /entrypoint

COPY ./compose/local/start /start
RUN sed -i 's/\r$//g' /start
RUN chmod +x /start

WORKDIR /build

COPY go.mod go.sum main.go ./
COPY app ./app
COPY pkg ./pkg
COPY cmd ./cmd
COPY platform ./platform

RUN go mod download && \
  go install github.com/air-verse/air@latest  && \
  go install github.com/swaggo/swag/cmd/swag@latest

ENTRYPOINT ["/entrypoint"]