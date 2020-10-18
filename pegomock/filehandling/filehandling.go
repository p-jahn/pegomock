package filehandling

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/petergtz/pegomock/mockgen"
	"github.com/petergtz/pegomock/model"
	"github.com/petergtz/pegomock/modelgen/gomock"
	"github.com/petergtz/pegomock/modelgen/loader"
	"github.com/petergtz/pegomock/pegomock/util"
)

func GenerateMockFileInOutputDir(
	args []string,
	outputDirPath string,
	outputFilePathOverride string,
	nameOut string,
	packageOut string,
	selfPackage string,
	debugParser bool,
	out io.Writer,
	useExperimentalModelGen bool,
	shouldGenerateMatchers bool,
	matchersDestination string) {

	// if a file path override is specified
	// ensure all directories in the path are created
	if outputFilePathOverride != "" {
		if err := os.MkdirAll(filepath.Dir(outputFilePathOverride), 0755); err != nil {
			panic(fmt.Errorf("Failed to make output directory, error: %v", err))
		}
	}

	GenerateMockFile(
		args,
		OutputFilePath(args, outputDirPath, outputFilePathOverride),
		nameOut,
		packageOut,
		selfPackage,
		debugParser,
		out,
		useExperimentalModelGen,
		shouldGenerateMatchers,
		matchersDestination)
}

func OutputFilePath(args []string, outputDirPath string, outputFilePathOverride string) string {
	if outputFilePathOverride != "" {
		return outputFilePathOverride
	} else if util.SourceMode(args) {
		return filepath.Join(outputDirPath, "mock_"+strings.TrimSuffix(args[0], ".go")+"_test.go")
	} else {
		return filepath.Join(outputDirPath, "mock_"+strings.ToLower(args[len(args)-1])+"_test.go")
	}
}

func GenerateMockFile(args []string, outputFilePath string, nameOut string, packageOut string, selfPackage string, debugParser bool, out io.Writer, useExperimentalModelGen bool, shouldGenerateMatchers bool, matchersDestination string) {
	mockSourceCode, matcherSourceCodes := GenerateMockSourceCode(args, nameOut, packageOut, selfPackage, debugParser, out, useExperimentalModelGen)

	err := ioutil.WriteFile(outputFilePath, mockSourceCode, 0664)
	if err != nil {
		panic(fmt.Errorf("Failed writing to destination: %v", err))
	}

	if shouldGenerateMatchers {
		matchersPath := filepath.Join(filepath.Dir(outputFilePath), "matchers")
		if matchersDestination != "" {
			matchersPath = matchersDestination
		}
		err = os.MkdirAll(matchersPath, 0755)
		if err != nil {
			panic(fmt.Errorf("Failed making dirs \"%v\": %v", matchersPath, err))
		}
		for matcherTypeName, matcherSourceCode := range matcherSourceCodes {
			err := ioutil.WriteFile(filepath.Join(matchersPath, matcherTypeName+".go"), []byte(matcherSourceCode), 0664)
			if err != nil {
				panic(fmt.Errorf("Failed writing to destination: %v", err))
			}
		}
	}
}

func GenerateMockSourceCode(args []string, nameOut string, packageOut string, selfPackage string, debugParser bool, out io.Writer, useExperimentalModelGen bool) ([]byte, map[string]string) {
	var err error

	var ast *model.Package
	var src string
	if util.SourceMode(args) {
		ast, err = gomock.ParseFile(args[0])
		src = args[0]
	} else {
		if len(args) != 2 {
			log.Fatal("Expected exactly two arguments, but got " + fmt.Sprint(args))
		}
		pkgName := strings.ReplaceAll(args[0], "\\", "/")
		if useExperimentalModelGen {
			ast, err = loader.GenerateModel(pkgName, args[1])
		} else {
			ast, err = gomock.Reflect(pkgName, strings.Split(args[1], ","))
		}
		src = fmt.Sprintf("%v (interfaces: %v)", args[0], args[1])
	}
	if err != nil {
		panic(fmt.Errorf("Loading input failed: %v", err))
	}

	if debugParser {
		ast.Print(out)
	}

	return mockgen.GenerateOutput(ast, src, nameOut, packageOut, selfPackage)
}
