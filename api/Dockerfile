FROM golang:alpine 

ENV GO111MODULE=on

WORKDIR go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
# RUN go get github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api .

EXPOSE 9090
# ENV PORT=3001
# CMD ["go", "run", "main.go"]
CMD ["./api"]

# 2nd stage.
# FROM scratch
# COPY --from=builder /app/api /app/
# EXPOSE 3001

# ENTRYPOINT ["/app/api"]
