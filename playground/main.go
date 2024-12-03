package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"os"
	"slices"
)

// parse the Go file, modify the AST, and save the result to another file
func addPrintfStatements(inputFilePath, outputFilePath string) error {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, inputFilePath, file, 0)
	if err != nil {
		return fmt.Errorf("could not parse file: %v", err)
	}

	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	var conf types.Config
	_, err = conf.Check("fib", fset, []*ast.File{node}, &info)
	if err != nil {
		log.Fatal(err)
	}

	// Walk through the AST and modify it
	ast.Inspect(node, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			seenVariables := make(map[string]bool)

			// If the function has a receiver, add its fields to the seen variables
			// if funcDecl.Recv != nil {
			// 	fmt.Println("not null recv list")
			// 	// Loop through each receiver field (if it's a struct type)
			// 	for _, field := range funcDecl.Recv.List {
			// 		// if ident, ok := field.Type.(*ast.Ident); ok {
			// 		// If the receiver is a simple type (like a struct), add its fields
			// 		// If the receiver is a pointer type, we unwrap it
			// 		if starExpr, ok := field.Type.(*ast.StarExpr); ok {
			// 			// Check if the underlying type is a struct

			// 			// If it's a pointer, check the type of the underlying type (starExpr.X)
			// 			if ident, ok := starExpr.X.(*ast.Ident); ok {
			// 				// Look up the type of the identifier (e.g., "data")
			// 				typ := info.Types[ident].Type
			// 				if structType, ok := typ.(*types.Struct); ok {
			// 					// The receiver is a pointer to a struct
			// 					fmt.Printf("Receiver is a pointer to a struct: %v\n", structType)
			// 				} else {
			// 					// It's not a struct
			// 					fmt.Printf("Receiver is not a struct: %v\n", typ)
			// 				}
			// 			}
			// 			if structType, ok := starExpr.X.(*ast.StructType); ok {
			// 				// Add each field's name to the seen variables
			// 				for _, f := range structType.Fields.List {
			// 					for _, fieldName := range f.Names {
			// 						seenVariables[fieldName.Name] = true
			// 					}
			// 				}
			// 			} else {
			// 				fmt.Printf("haha %v\n", reflect.TypeOf(starExpr.X))
			// 			}
			// 		} else if structType, ok := field.Type.(*ast.StructType); ok {
			// 			// If the receiver is not a pointer, but directly a struct type
			// 			for _, f := range structType.Fields.List {
			// 				for _, fieldName := range f.Names {
			// 					seenVariables[fieldName.Name] = true
			// 				}
			// 			}
			// 		} else {
			// 			fmt.Printf("not star expr or struct type\n")
			// 		}
			// 	}
			// }

			for _, param := range funcDecl.Type.Params.List {
				for _, name := range param.Names {
					seenVariables[name.Name] = true
				}
			}

			for i := 0; i < len(funcDecl.Body.List); i++ {
				stmt := funcDecl.Body.List[i]

				if exprStmt, ok := stmt.(*ast.ExprStmt); ok {
					if callExpr, ok := exprStmt.X.(*ast.CallExpr); ok {
						if fun, ok := callExpr.Fun.(*ast.Ident); ok && fun.Name == "Debug" {
							printStmt := createPrintStmt(seenVariables)

							funcDecl.Body.List = slices.Insert(funcDecl.Body.List, i, printStmt)
							i++
						}
					}
				}

				if assignStmt, ok := stmt.(*ast.AssignStmt); ok {
					for _, lhs := range assignStmt.Lhs {
						if ident, ok := lhs.(*ast.Ident); ok {
							seenVariables[ident.Name] = true
						}
					}
				}
			}
		}
		return true
	})

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("could not create output file: %v", err)
	}
	defer outFile.Close()

	err = printer.Fprint(outFile, fset, node)
	if err != nil {
		return fmt.Errorf("could not print modified AST: %v", err)
	}

	return nil
}

// create a fmt.Printf statement that prints all variable names seen so far
func createPrintStmt(seenVariables map[string]bool) ast.Stmt {
	format := "Variables seen so far:"
	args := []ast.Expr{}
	for varName := range seenVariables {
		format += " %v"
		args = append(args, &ast.Ident{Name: varName})
	}
	format += "\\n"

	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "fmt"},
				Sel: &ast.Ident{Name: "Printf"},
			},
			Args: append([]ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"` + format + `"`, ValuePos: token.NoPos},
			}, args...),
		},
	}
}

func main() {
	inputFilePath := "simple.go"
	outputFilePath := "simple.go"

	err := addPrintfStatements(inputFilePath, outputFilePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		fmt.Printf("Successfully saved modified Go code to %s\n", outputFilePath)
	}
}
