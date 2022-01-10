---
# goyacc-flex
基于goyacc和flex实现的逻辑语法解析器

## 词法符号

## 词法解析

通过flex工具生成C语言词法解析器文件：

```shell
$ flex --prefix=yy --header-file=yy.lex.h -o yy.lex.c lexer.l
```

其中`--prefix`表示生成的代码中标识符都是以`yy`前缀。在一个项目有多个flex生成代码时，可通过前缀区分。`--header-file`表示生成头问题，这样方便在其它代码中引用生成的词法分析函数。`-o`指定输出源代码文件的名字。

## 语法分析

通过goyacc工具生成代码：

```shell
$ goyacc -o yy.parser.go -p "yy" parser.ypp
```