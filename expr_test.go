package goyacc_lex

import "testing"

func TestExprLexer(t *testing.T) {
	rule := "((b:\"90rule = b:\\\"90sdf||\\\" || c:\\\"asdf\\\"sdf\"||e:\"asdfeer\") && a:\"asdf|||||&&&&d|fd&&||\" || c:\"aasdfasdfasdf\")"

	//rule = "b:\"90sdf||\" || c:\"asdf\""

	lex := newExprLexer([]byte(rule), nil)
	ret := yyParse(lex)
	if ret != 0 {
		t.Fatalf("parse failed, ret:%d\n", ret)
	}

}
