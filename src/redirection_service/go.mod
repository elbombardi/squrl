module github.com/elbombardi/squrl/src/redirection_service

go 1.21.0

replace github.com/elbombardi/squrl/src/db => ../db

require (
	github.com/elbombardi/squrl/src/db v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.48.0
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.48.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
)