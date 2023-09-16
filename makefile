public/wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./public/wasm_exec.js

wasm-bin: public/wasm_exec.js
	cd wasm && \
		GOOS=js GOARCH=wasm go build -o ../public/index.wasm  

devserver:
	cd cmd/devserver && \
		go run main.go
