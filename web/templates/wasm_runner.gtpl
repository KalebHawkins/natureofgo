<!DOCTYPE html>
<script src="/static/wasm_exec.js"></script>
<script>
const go = new Go();
WebAssembly.instantiateStreaming(fetch("/static/{{.WasmFile}}"), go.importObject).then(result => {
    go.run(result.instance);
});
</script>