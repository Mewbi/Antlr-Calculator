grammar Calculator;

// Parser rules
prog:   stat+ ;

stat:   expr NEWLINE                # printExpr
    |   ID '=' expr NEWLINE         # assign
    |   NEWLINE                     # blank
    ;

expr:   expr op=('*'|'/') expr                          # MulDiv
    |   expr op=('+'|'-') expr                          # AddSub
    |   expr op=('=='|'!='|'<'|'>'|'<='|'>=') expr      # OpRel
    |   NUMBER                                          # number
    |   ID                                              # id
    |   '(' expr ')'                                    # parens
    ;

// Lexer rules
MUL:    '*' ;
DIV:    '/' ;
ADD:    '+' ;
SUB:    '-' ;

EQUAL:        '==' ;
DIFF:         '!=' ;
LESS:         '<' ;
GREATER:      '>' ;
LESSEQ:       '<=' ;
GREATEREQ:    '>=' ;

ID:         [a-zA-Z_][a-zA-Z_0-9]* ;
NUMBER:     [0-9]+ ('.'[0-9]+)? ;

NEWLINE:    '\r'? '\n' ;
WS:         [ \t]+ -> skip ;
