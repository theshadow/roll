grammar Dice;

D       : [dD] ;
SIGN    : [\-+] ;
LPAREN  : '(' ;
RPAREN  : ')' ;
COMMA   : ',' ;
SPACE   : ' ' ;

WS      : [\r\n\t] -> skip ;

Integer       : [0-9]+ ;
Id            : [a-zA-Z][A-Za-z0-9]+ ;
StringLiteral : '"' ~('\'' | '\\' | '\n' | '\r')+ '"' ;

formula       : count? sides modifier? extensions ;
extensions    : (SPACE funccall)* ;
count         : Integer ;
sides         : Id ;
modifier      : SIGN Integer ;
parameter     : Integer | StringLiteral ;
parameters    : (parameter COMMA)* parameter ;
funcname      : Id ;
funccall      : funcname '(' parameters? ')' ;