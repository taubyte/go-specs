package wasm

var (
	Utils         = "utils"
	ConfigVersion = "0.1.0"

	GoCodeFile  = "main.go"
	ASCodeFile  = "index.ts"
	RSCodeFile  = "lib.rs"
	RSCargoFile = "Cargo.toml"
	BuildFile   = "build.sh"

	WasmFileName      = "artifact"
	WasmExt           = ".wasm"
	WasmCompressedExt = ".zwasm"

	GolangImage         = "taubyte/builder:golang"
	AssemblyScriptImage = "taubyte/builder:assembly-script"
	RustImage           = "taubyte/builder:rust"
	CustomImage         = "taubyte/builder:custom"

	BufferSize = 1024
)
