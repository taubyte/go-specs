package wasm

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/otiai10/copy"
	git "github.com/taubyte/go-simple-git"
	"github.com/taubyte/go-specs/builders"
	functionSpec "github.com/taubyte/go-specs/function"
)

func Wd(workDir string) Dir {
	wd := builders.Wd(workDir)
	return Dir{wd}
}

func (d Dir) BoilerPlate() error {
	if err := os.MkdirAll(d.SourceDir(), 0755); err != nil {
		return fmt.Errorf("creating source dir failed with: %s", err)
	}

	if err := os.Mkdir(d.OutDir(), 0755); err != nil {
		return fmt.Errorf("creating out dir failed with: %s", err)
	}

	return nil
}

func (d Dir) WasmOutput() string {
	return path.Join(d.OutDir(), WasmFileName+WasmExt)
}

func (d Dir) WasmCompressed() string {
	return path.Join(d.String(), WasmFileName+WasmCompressedExt)
}

func (s SupportedLanguage) Extension() string {
	return supportedLanguages[s]
}

func LangFromExt(ext string) *SupportedLanguage {
	for language := range supportedLanguages {
		if language.Extension() == ext {
			return &language
		}
	}

	return nil
}

func (s SupportedLanguage) CopyFunctionTemplateConfig(templateRepo *git.Repository, wd, destination string) error {
	if templateRepo == nil {
		return errors.New("template is nil")
	}

	if _, err := os.Stat(destination); err != nil {
		return err
	}

	templatePath := path.Join(wd, templateRepo.Dir(), "code", functionSpec.PathVariable.String(), string(s), "common")
	return copy.Copy(templatePath, destination)
}
