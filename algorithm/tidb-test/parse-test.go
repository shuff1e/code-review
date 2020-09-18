package main

import (
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

type visitor struct{}

func (v *visitor) Enter(in ast.Node) (out ast.Node, skipChildren bool) {
	fmt.Printf("%T\n", in)
	return in, false
}

func (v *visitor) Leave(in ast.Node) (out ast.Node, ok bool) {
	return in, true
}

func main() {

	sql := "SELECT /*+ TIDB_SMJ(employees) */ emp_no, first_name, last_name " +
		"FROM employees USE INDEX (last_name) " +
		"where last_name='Aamodt' and gender='F' and birth_date > '1960-01-01'"

	sqlParser := parser.New()
	stmtNodes, warns, err := sqlParser.Parse(sql, "", "")
	for _,v := range warns {
		fmt.Println(v)
	}
	if err != nil {
		fmt.Printf("parse error:\n%v\n%s", err, sql)
		return
	}
	for _, stmtNode := range stmtNodes {
		v := visitor{}
		stmtNode.Accept(&v)
	}
}