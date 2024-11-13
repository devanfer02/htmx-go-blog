# Golang Builder
FROM golang:alpine3.20 as go-builder

RUN apk update && apk add --no-cache git 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./app/cmd 

# HTMX n Tailwind Builder
FROM node:22 as js-builder

WORKDIR /app

COPY package.json package-lock.json vite.config.js tailwind.config.js postcss.config.js ./

COPY ./resources ./resources

COPY ./templates ./templates

RUN npm install

RUN npm run build

# Final Stage
FROM alpine:3.20

WORKDIR /app

RUN apk --no-cache add curl ca-certificates tzdata

COPY --from=go-builder /app/.env /app/.env
COPY --from=go-builder /app/bin/main /app/bin/main
COPY --from=go-builder /app/templates /app/templates 
COPY --from=js-builder /app/static /app/static 

EXPOSE 8080

CMD ["./bin/main"]