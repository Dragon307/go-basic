

## [go-wasm](https://github.com/golang/go/wiki/WebAssembly)

> GOOS=js GOARCH=wasm go build -o main.wasm

> cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

install goexec: go get -u github.com/shurcooL/goexec
> goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'



## run wasm in node

