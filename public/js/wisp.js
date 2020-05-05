const go = new Go()

const runButton = document.getElementById('run-button')
const input = document.getElementById('input')
const log = document.getElementById('log')
log.value = ''

let wasm

// Given the text and a module, insert the string into its memory.
function insertText(text, module) {

    // Get the address of the writable memory.
    let addr = module.exports.getBuffer()
    let buf = module.exports.memory.buffer

    let mem = new Int8Array(buf)
    let view = mem.subarray(addr, addr + text.length)

    for (let i = 0; i < text.length; i++) {
        view[i] = text.charCodeAt(i)
    }

    // Return the address we started at.
    return addr
}

// Function that returns the acutal function we want to attach to the "Run" button
function onRun(runner, module) {
    return (event) => {
        // First, need to run the module in order to define everything.
        runner.run(module)

        let inputText = input.value
        let addr = insertText(inputText, module)

        // Now just pass the memory location to the relevant function.
        module.exports.echo(addr, inputText.length)
    }
}

function logText(addr, length) {
    let memory = wasm.exports.memory
    let bytes = memory.buffer.slice(addr, addr + length)
    let text = String.fromCharCode.apply(null, new Int8Array(bytes))

    log.value += text + '\n'
}

// Add our own functions to the env we pass to the wasm module
go.importObject.env['main.go.log'] = logText

WebAssembly.instantiateStreaming(fetch('js/wisp.wasm'), go.importObject)
    .then(module => {
        wasm = module.instance

        runButton.disabled = false
        runButton.addEventListener('click', onRun(go, wasm))
    })
