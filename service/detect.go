package service

import (
	"bytes"
	"domo1/util/dto"
	"domo1/util/response"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
)

func RuncodeService(request dto.CodeDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	userid := request.Userid
	problemid := request.Problemid
	code := request.Code

	filename, err := SaveCode(userid, problemid, code)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	cmd := exec.Command("go", "run", filename)
	out, err := os.Create("user.out")
	defer out.Close()
	if err != nil {
		log.Fatalln(err)
	}
	in, err := os.Open("user.in")
	defer in.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var stderr bytes.Buffer
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	// read answer file
	answerBytes, err := ioutil.ReadFile("answer.out")
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	answer := string(answerBytes)

	// read user output file
	userBytes, err := ioutil.ReadFile("user.out")
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	user := string(userBytes)

	// match output with regex
	r := regexp.MustCompile(`([\w]+)`)

	answerMatch := r.FindStringSubmatch(answer)
	userMatch := r.FindStringSubmatch(user)
	// compare results
	if answerMatch[0] != userMatch[0] {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.OutputIncorrect
		return res
	}
	return res
}

func CheckFuncVarService(request dto.FuncVarDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	userid := request.Userid
	problemid := request.Problemid
	inVars := request.Vars
	inFuncs := request.Funcs
	code := request.Code
	filename, err := SaveCode(userid, problemid, code)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	vars, funcs, err := SearchFuncVar(filename)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	for k := range vars {
		for _, v := range inVars {
			if k != v {
				delete(vars, k)
			}
		}
	}

	for k := range funcs {
		for _, v := range inFuncs {
			if k == v {
				delete(vars, k)
			}
		}
	}
	res.Data = gin.H{
		"funcs": funcs,
		"vars":  vars,
	}
	return res
}

func SaveCode(userid, problemid, code string) (string, error) {
	var filename string
	filename = fmt.Sprintf("./file/%v/%v/code.out", userid, problemid)
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	os.WriteFile(file.Name(), []byte(code), 0777)
	return filename, nil
}

func SearchFuncVar(filename string) (vars, funcs map[string]bool, err error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return
	}

	for _, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			funcs[d.Name.Name] = true
			ast.Inspect(d.Body, func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.CallExpr:
					if ident, ok := x.Fun.(*ast.Ident); ok {
						funcs[ident.Name] = true
						if ident.Obj != nil && ident.Obj.Kind == ast.Fun {
							if typ, ok := ident.Obj.Decl.(*ast.FuncDecl); ok {
								for _, field := range typ.Type.Params.List {
									for _, name := range field.Names {
										vars[name.Name] = true
									}
								}
							}
						}
					}
				case *ast.AssignStmt:
					for _, ident := range x.Lhs {
						if ident, ok := ident.(*ast.Ident); ok {
							vars[ident.Name] = true
						}
					}
				case *ast.DeclStmt:
					switch x := x.Decl.(type) {
					case *ast.GenDecl:
						if x.Tok == token.VAR || x.Tok == token.CONST {
							for _, spec := range x.Specs {
								switch s := spec.(type) {
								case *ast.ValueSpec:
									for _, ident := range s.Names {
										vars[ident.Name] = true
									}
								}
							}
						}
					}

				}
				return true
			})
		case *ast.GenDecl:
			if d.Tok == token.VAR || d.Tok == token.CONST {
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.ValueSpec:
						for _, ident := range s.Names {
							vars[ident.Name] = true
						}
					}
				}
			}
		}
	}
	return
}
