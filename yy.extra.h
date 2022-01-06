#ifndef _YY_EXTRA_H
#define _YY_EXTRA_H

struct yyextra_t
{
	char str_buff[4096];
	int buf_pos;
};

static void yyextra_text_append(struct yyextra_t* ext, char a){
	if(ext->buf_pos >= sizeof(ext->str_buff)){
		//TODO
	}
	ext->str_buff[ext->buf_pos++] = a;
}
static void yyextra_text_clear(struct yyextra_t* ext){
	ext->buf_pos = 0;
}
static char * yyextra_text_get(struct yyextra_t* ext){
	return ext->str_buff;
}
static int yyextra_text_len(struct yyextra_t* ext){
	return ext->buf_pos;
}

enum {
	EOF = 0,
	ILLEGAL = 10000,

	KEY = 258,
	VALUE = 259,

	LOGIC_AND = 260, // /
	LOGIC_OR = 261, // |

	LPAREN = 262, // (
	RPAREN = 263, // )
	COLON = 264, // :
};

#endif