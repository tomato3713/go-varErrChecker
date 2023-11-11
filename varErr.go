package varErrChecker

import (
	"fmt"
	"go/ast"
	"go/types"
	"log/slog"
	"os"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "varErrChecker",
	Doc:  "varErrChecker reports whether error variables at the top of the package match naming conventions.",
	Run:  run,
}

const msg = "Error variable %s does not follow naming conventions, rule is %s"

var (
	pattern string
	debug   bool
)

func init() {
	Analyzer.Flags.StringVar(&pattern, "pattern", `^Err[\d\w]+$`, `pattern is a regular expression string that represents a naming convention for errors. default: ^Err[\d\w]+$`)
	Analyzer.Flags.BoolVar(&debug, "debug", false, "debug is a flag for debugging.")
}

var errorType = types.Universe.Lookup("error").Type()
var errorInterface = errorType.Underlying().(*types.Interface)

func run(pass *analysis.Pass) (interface{}, error) {
	// setup logger
	programLevel := new(slog.LevelVar)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})

	if debug {
		// enable debug mode
		programLevel.Set(slog.LevelDebug)
	} else {
		slog.SetDefault(slog.New(handler))
	}

	logger := slog.New(handler)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for _, f := range pass.Files {
		logger.Debug("start to check", "file", pass.Fset.File(f.Pos()).Name())

		// check all declared variables whoose error interface is satisfied
		for _, decl := range f.Decls {
			// ignore import declation
			// ref: https://pkg.go.dev/go/ast#GenDecl
			if decl, ok := decl.(*ast.GenDecl); ok {
				for _, spec := range decl.Specs {
					if typeSpec, ok := spec.(*ast.ValueSpec); ok {
						logger.Debug("find variable declaration", "line", typeSpec.Pos())
						for _, name := range typeSpec.Names {
							obj := pass.TypesInfo.ObjectOf(name)
							if obj == nil {
								continue
							}
							logger.Debug("check if error interface is satisfied", "line", obj.Pos(), "variable", fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name()))
							if types.Implements(obj.Type(), errorInterface) {
								logger.Debug("check if follow naming conventions since error interface is satisfied", "line", obj.Pos(), "variable", fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name()))
								if !re.MatchString(obj.Name()) {
									logger.Debug("did not follow naming conventions", "line", obj.Pos(), "variable", fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name()), "naming conventions", pattern)
									str := fmt.Sprintf(msg, obj.Name(), re.String())
									pass.Reportf(obj.Pos(), str)
									continue
								}
								logger.Debug("followed naming conventions", "line", obj.Pos(), "variable", fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name()), "naming conventions", pattern)
							}
							logger.Debug("error interface is not satisfied", "line", obj.Pos(), "variable", fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name()))
						}
					}
				}
			}
		}
	}

	return nil, nil
}
