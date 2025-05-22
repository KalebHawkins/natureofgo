$Env:GOOS = 'js'
$Env:GOARCH = 'wasm'

go build -o web/static/example0.1.wasm ./Example0.1/
go build -o web/static/example0.2.wasm ./Example0.2/
go build -o web/static/exercise0.1.wasm ./Exercise0.1/
go build -o web/static/exercise0.3.wasm ./Exercise0.3/
go build -o web/static/example0.4.wasm ./Example0.4/
go build -o web/static/exercise0.4.wasm ./Exercise0.4/
go build -o web/static/exercise0.5.wasm ./Exercise0.5/
go build -o web/static/Example0.5.wasm ./Example0.5/
go build -o web/static/Exercise0.6.wasm ./Exercise0.6/
go build -o web/static/Example0.6.wasm ./Example0.6/
go build -o web/static/Exercise0.7.wasm ./Exercise0.7/
go build -o web/static/Exercise0.8.wasm ./Exercise0.8/

Remove-Item Env:GOOS
Remove-Item Env:GOARCH