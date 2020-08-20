# Build .wasm file
$env:GOOS='js'
$env:GOARCH='wasm'
go build -o web/dragons.wasm cmd/wasm/main.go
$env:GOOS=''
$env:GOARCH=''

# Run development server
go run cmd/web/main.go
