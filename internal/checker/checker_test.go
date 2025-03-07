package checker_test

import (
	"strings"
	"testing"

	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/bundled"
	"github.com/microsoft/typescript-go/internal/checker"
	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/repo"
	"github.com/microsoft/typescript-go/internal/tspath"
	"github.com/microsoft/typescript-go/internal/vfs/osvfs"
	"github.com/microsoft/typescript-go/internal/vfs/vfstest"
)

func TestGetSymbolAtLocation(t *testing.T) {
	t.Parallel()

	content := `interface Foo {
  bar: string;
}
declare const foo: Foo;
foo.bar;`
	fs := vfstest.FromMap(map[string]string{
		"/foo.ts": content,
		"/tsconfig.json": `
				{
					"compilerOptions": {}
				}
			`,
	}, false /*useCaseSensitiveFileNames*/)
	fs = bundled.WrapFS(fs)

	cd := "/"
	host := compiler.NewCompilerHost(nil, cd, fs, bundled.LibPath())
	opts := compiler.ProgramOptions{
		Host:           host,
		ConfigFileName: "/tsconfig.json",
	}
	p := compiler.NewProgram(opts)
	p.BindSourceFiles()
	c := p.GetTypeChecker()
	file := p.GetSourceFile("/foo.ts")
	interfaceId := file.Statements.Nodes[0].Name()
	varId := file.Statements.Nodes[1].AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes[0].Name()
	propAccess := file.Statements.Nodes[2].AsExpressionStatement().Expression
	nodes := []*ast.Node{interfaceId, varId, propAccess}
	for _, node := range nodes {
		symbol := c.GetSymbolAtLocation(node)
		if symbol == nil {
			t.Fatalf("Expected symbol to be non-nil")
		}
	}
}

func TestCheckSrcCompiler(t *testing.T) {
	t.Parallel()

	repo.SkipIfNoTypeScriptSubmodule(t)
	fs := osvfs.FS()
	fs = bundled.WrapFS(fs)

	rootPath := tspath.CombinePaths(tspath.NormalizeSlashes(repo.TypeScriptSubmodulePath), "src", "compiler")

	host := compiler.NewCompilerHost(nil, rootPath, fs, bundled.LibPath())
	opts := compiler.ProgramOptions{
		Host:           host,
		ConfigFileName: tspath.CombinePaths(rootPath, "tsconfig.json"),
	}
	p := compiler.NewProgram(opts)
	p.CheckSourceFiles()
}

func BenchmarkNewChecker(b *testing.B) {
	repo.SkipIfNoTypeScriptSubmodule(b)
	fs := osvfs.FS()
	fs = bundled.WrapFS(fs)

	rootPath := tspath.CombinePaths(tspath.NormalizeSlashes(repo.TypeScriptSubmodulePath), "src", "compiler")

	host := compiler.NewCompilerHost(nil, rootPath, fs, bundled.LibPath())
	opts := compiler.ProgramOptions{
		Host:           host,
		ConfigFileName: tspath.CombinePaths(rootPath, "tsconfig.json"),
	}
	p := compiler.NewProgram(opts)

	b.ReportAllocs()

	for b.Loop() {
		checker.NewChecker(p)
	}
}

func BenchmarkCheckSrcCompiler(b *testing.B) {
	repo.SkipIfNoTypeScriptSubmodule(b)
	fs := osvfs.FS()
	fs = bundled.WrapFS(fs)

	rootPath := tspath.CombinePaths(tspath.NormalizeSlashes(repo.TypeScriptSubmodulePath), "src", "compiler")

	host := compiler.NewCompilerHost(nil, rootPath, fs, bundled.LibPath())
	opts := compiler.ProgramOptions{
		Host:           host,
		ConfigFileName: tspath.CombinePaths(rootPath, "tsconfig.json"),
	}
	p := compiler.NewProgram(opts)
	p.BindSourceFiles()

	files := p.SourceFiles()
	findFile := func(name string) *ast.SourceFile {
		return core.Find(files, func(f *ast.SourceFile) bool {
			return strings.HasSuffix(f.FileName(), name)
		})
	}

	tests := []string{
		"/compiler/checker.ts",
		"/compiler/core.ts",
		"/compiler/program.ts",
		"/compiler/sys.ts",
		"/compiler/types.ts",
	}

	for _, test := range tests {
		b.Run(tspath.GetBaseFileName(test), func(b *testing.B) {
			file := findFile(test)

			b.ReportAllocs()

			for b.Loop() {
				c := checker.NewChecker(p)
				c.CheckSourceFile(file)
			}
		})
	}
}
