
%top{
	#include "yy.extra.h"

	#define YY_DECL int yylex(yyscan_t yyscanner)
	#define YY_EXTRA_TYPE struct yyextra_t*
}

%option noyywrap
%option reentrant

%x STRING
key -?[a-zA-Z_]+
%%

{key} return KEY;

<STRING>"\"" {
	BEGIN(INITIAL);
	return VALUE;
}

<STRING>. {
	yyextra_text_append(yyget_extra(yyscanner), yyget_text(yyscanner)[0]);
}

<STRING>\\. {
	yyextra_text_append(yyget_extra(yyscanner), yyget_text(yyscanner)[1]);
}

"\"" {
	yyextra_text_clear(yyget_extra(yyscanner));
	BEGIN(STRING);
}

"&&" return LOGIC_AND;
"||" return LOGIC_OR;

"(" return LPAREN;
")" return RPAREN;
":" return COLON;

%%

