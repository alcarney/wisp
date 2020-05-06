const go = new Go()

const runButton = document.getElementById('run-button')
const input = document.getElementById('input')
const output = document.getElementById('output')
const log = document.getElementById('log')


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

// Given the address, length and module, extract the string from the described
// memory location
function extractText(addr, length, module) {
    let memory = module.exports.memory
    let bytes = memory.buffer.slice(addr, addr + length)
    return String.fromCharCode.apply(null, new Int8Array(bytes))
}

// Function that returns the acutal function we want to attach to the "Run" button
function onRun(runner, module) {
    return (event) => {

        log.value = ''
        output.value = ''

        // First, need to run the module in order to define everything.
        runner.run(module)

        let inputText = input.value
        let addr = insertText(inputText, module)

        // Now just pass the memory location to the relevant function.
        module.exports.runBf(addr, inputText.length)
    }
}

function logText(addr, length) {
    let text = extractText(addr, length, wasm)
    log.value += text + '\n'
}


function outputText(addr, length) {
    let text = extractText(addr, length, wasm)
    output.value += text
}

// Add our own functions to the env we pass to the wasm module
go.importObject.env['main.go.log'] = logText
go.importObject.env['main.go.output'] = outputText

WebAssembly.instantiateStreaming(fetch('js/wisp.wasm'), go.importObject)
    .then(module => {
        wasm = module.instance

        runButton.disabled = false
        runButton.addEventListener('click', onRun(go, wasm))
    })
