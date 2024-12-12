FROM golang:1.22

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o main ./demo/main.go

ENV API_BASE_URL=https://api.buttercms.com/v2
ENV API_KEY=your_api_key

CMD ["./main"]