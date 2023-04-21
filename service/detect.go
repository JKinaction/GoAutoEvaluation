package service

import (
	"bytes"
	"domo1/util/common"
	"domo1/util/dto"
	"domo1/util/model"
	"domo1/util/response"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 运行路径./file/user/Userid/Questionid/code.go
// 输出路径./file/user/Userid/Questionid/user.out
// 输入路径./file/question/Questionid/Intputanwserid/user.in
// 对比路径./file/question/Questionid/Intputanwserid/answer.out
func RuncodeService(request dto.CodeDto) response.ResponseStruct {
	res := response.NewResponse()

	userid := request.Userid
	questionid := request.Questionid
	code := request.Code

	filename, err := SaveCode(userid, questionid, code)
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	cmd := exec.Command("go", "run", filename)

	outpath := fmt.Sprintf("./file/user/%v/%v/user.out", userid, questionid)

	out, err := os.Create(outpath)
	defer out.Close()
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	var ia model.InputAnswer
	common.GetDB().Where("questionid = ?", questionid).Where("path != ?", "").First(&ia) //-------------------------------
	inpath := fmt.Sprintf("./file/question/%v/%v/user.in", questionid, ia.ID)
	in, err := os.Open(inpath)
	defer in.Close()
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	var stderr bytes.Buffer
	cmd.Stdin = in
	cmd.Stdout = out //运行用户代码输出
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		res.Data = gin.H{
			"err": err,
		}
		return res
	}
	// read answer file
	answerpath := fmt.Sprintf("./file/question/%v/%v/answer.out", questionid, ia.ID)
	answerBytes, err := ioutil.ReadFile(answerpath)
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	answer := string(answerBytes)

	// read user output file
	userBytes, err := ioutil.ReadFile(outpath)
	if err != nil {
		logrus.Println(err)
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	user := string(userBytes)

	// match output with regex 改进方法按行读入，再每行split对比——————————————————————————————————————————————————
	r := regexp.MustCompile(`([\w]+)`)

	answerMatch := r.FindStringSubmatch(answer)
	userMatch := r.FindStringSubmatch(user)
	for i := 0; i < len(answerMatch) && i < len(userMatch); i++ {
		// compare results
		if answerMatch[i] != userMatch[i] {
			res.HttpStatus = http.StatusBadRequest
			res.Code = response.FailCode
			res.Msg = response.OutputIncorrect
			res.Data = gin.H{
				"data": "答案错误",
			}
			return res
		}
	}
	res.Data = gin.H{
		"data": "答案正确",
	}
	return res
}

func CheckFuncVarService(request dto.FuncVarDto) response.ResponseStruct {
	res := response.NewResponse()
	userid := request.Userid
	questionid := request.Questionid
	inVars := request.Vars
	inFuncs := request.Funcs
	code := request.Code
	//保存代码到本地
	filename, err := SaveCode(userid, questionid, code)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	//查找所有变量和函数
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
				delete(funcs, k)
			}
		}
	}
	res.Data = gin.H{
		"未匹配到的funcs": funcs,
		"未匹配到的vars":  vars,
	}
	return res
}

func SaveCode(userid, questionid int, code string) (string, error) {

	if err := os.MkdirAll(fmt.Sprintf("./file/user/%v/%v", userid, questionid), os.ModePerm); err != nil {
		logrus.Println(err)
		return "", err
	}
	filename := fmt.Sprintf("./file/user/%v/%v/code.go", userid, questionid)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if _, err := file.WriteString(code); err != nil {
		return "", err
	}
	return filename, nil
}

func SearchFuncVar(filename string) (map[string]bool, map[string]bool, error) {
	vars := make(map[string]bool)
	funcs := make(map[string]bool)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
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
	return nil, nil, err
}
