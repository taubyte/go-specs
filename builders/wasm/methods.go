package wasm

import (
	"path"

	"github.com/taubyte/go-specs/builders"
)

func Wd(workDir string) Dir {
	wd := builders.Wd(workDir)
	return Dir{wd}
}

func (d Dir) WasmOutput() string {
	return path.Join(d.OutDir(), WasmFileName+WasmExt)
}

func (d Dir) WasmCompressed() string {
	return path.Join(d.String(), WasmFileName+WasmCompressedExt)
}

func (d Dir) BuildFile() string {
	return path.Join(d.TaubyteDir(), BuildFile)
}

func (d Dir) GoCodeFile() string {
	return path.Join(d.SourceDir(), GoCodeFile)
}

func (d Dir) AsFile() string {
	return path.Join(d.SourceDir(), ASCodeFile)
}

func (d Dir) RsFile() string {
	return path.Join(d.SourceDir(), RSCodeFile)
}

func (d Dir) RsCargoFile() string {
	return path.Join(d.SourceDir(), RSCargoFile)
}
