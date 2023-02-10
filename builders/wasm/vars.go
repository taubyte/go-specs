package wasm

const (
	Utils         = "utils"
	ConfigVersion = "0.1.0"
	BuildFile     = "build.sh"

	WasmFileName      = "artifact"
	WasmExt           = ".wasm"
	WasmCompressedExt = ".zwasm"

	BufferSize = 1024
)

// Supported Languages
const (
	Rust           SupportedLanguage = "Rust"
	Go             SupportedLanguage = "Go"
	AssemblyScript SupportedLanguage = "Assembly_Script"
)

var supportedLanguages = map[SupportedLanguage]string{
	Rust:           ".rs",
	Go:             ".go",
	AssemblyScript: ".ts",
}

func SupportedLanguages() map[SupportedLanguage]string {
	return supportedLanguages
}
