p := public/files/wasm/index.wasm
j := public/files/

dev:
	go run example/serve/main.go
wasm-prod:
	GOARCH=wasm GOOS=js go build -ldflags '-w -s' -o $(p) example/client/main.go
	wasm-opt -Oz -o ${p}.opt ${p}
	du -sh $(p).opt
	gzip -5 -f -k $(p).opt && du -sh $(p).gz
wasm:
	GOARCH=wasm GOOS=js go build -ldflags '-w -s' -o $(p) example/client/main.go
build-js:
	gopherjs build example/client/main.go
	esbuild main.js --bundle --outfile=out.js --minify
ser:
	echo "serve-run"
tailwind:
	tailwindcss -i ./tailwind.css -o ./public/tailwind.css --watch
build-client:
	echo "build-client"
build-serve:
	echo "build-client"
test:
	go test ./...go test ./...
