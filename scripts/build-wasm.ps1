$Env:GOOS = 'js'
$Env:GOARCH = 'wasm'

go build -o web/static/example0.1.wasm ./Example0.1/
go build -o web/static/example0.2.wasm ./Example0.2/
go build -o web/static/exercise0.1.wasm ./Exercise0.1/
go build -o web/static/exercise0.3.wasm ./Exercise0.3/
go build -o web/static/example0.4.wasm ./Example0.4/
go build -o web/static/exercise0.4.wasm ./Exercise0.4/

Remove-Item Env:GOOS
Remove-Item Env:GOARCH