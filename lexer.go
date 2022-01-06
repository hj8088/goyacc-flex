package goyacc_lex

/*
#include "yy.lex.h"
extern int yylex(yyscan_t yyscanner);
*/
import "C"
import (
	"github.com/pkg/errors"
	"log"
	"reflect"
	"unsafe"
)

type ExprLexer struct {
	yyLineno int
	yyText   string
	lastErr  error

	scanner unsafe.Pointer
	extra   C.struct_yyextra_t

	caller Caller
}

var _ yyLexer = (*ExprLexer)(nil)

func newExprLexer(data []byte, caller Caller) *ExprLexer {
	p := &ExprLexer{
		scanner: unsafe.Pointer(nil),
		caller:  caller,
	}

	if caller == nil || reflect.ValueOf(caller).IsNil() {
		p.caller = &DefaultPrintCaller{}
	}

	C.yylex_init_extra((*C.struct_yyextra_t)(&p.extra), (*C.yyscan_t)(&p.scanner))

	C.yy_scan_bytes(
		(*C.char)(C.CBytes(data)),
		C.yy_size_t(len(data)),
		(C.yyscan_t)(p.scanner),
	)

	return p
}

func (e *ExprLexer) copyStr(src string) string {
	dst := make([]byte, len(src))
	copy(dst, src)
	return string(dst)
}

func (e *ExprLexer) Lex(yylval *yySymType) int {

	e.lastErr = nil

	scanner := (C.yyscan_t)(e.scanner)

	var tok = C.yylex(scanner)

	e.yyLineno = int(C.yyget_lineno(scanner))
	e.yyText = *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(C.yyget_text(scanner))),
		Len:  int(C.yyget_leng(scanner)),
	}))

	switch tok {
	case C.KEY:
		yylval.str_ = e.copyStr(e.yyText)
		return KEY

	case C.VALUE:
		pExtra := (*C.struct_yyextra_t)(&e.extra)
		yyValue := *(*string)(unsafe.Pointer(&reflect.StringHeader{
			Data: uintptr(unsafe.Pointer(C.yyextra_text_get(pExtra))),
			Len:  int(C.yyextra_text_len(pExtra)),
		}))
		e.yyText = yyValue
		yylval.str_ = e.copyStr(yyValue)
		return VALUE

	case C.LOGIC_AND:
		return LOGICAL_AND
	case C.LOGIC_OR:
		return LOGICAL_OR
	case C.LPAREN:
		return LEFT_PAREN
	case C.RPAREN:
		return RIGHT_PAREN
	case C.COLON:
		return COLON
	}

	if tok == C.ILLEGAL {
		log.Printf("lex: ILLEGAL token, yytext = %q, yylineno = %d", e.yyText, e.yyLineno)
	}

	return 0 // eof
}

func (e *ExprLexer) Error(s string) {
	e.lastErr = errors.New("yacc: " + s)
	if err := e.lastErr; err != nil {
		log.Println(err)
	}
}
