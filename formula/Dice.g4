grammar Dice;

formula : count? D sides modifier? extensions? EOF;
count : DIGIT+ ;
sides : DIGIT+ ;
modifier : SIGN DIGIT+ ;
parameter: DIGIT+ | DQUOTE WORD+ DQUOTE;
parameters : parameter (COMMA SPACE? parameter) ;
extensionname : ALPHA WORD+ ;
extension : extensionname LPAREN parameters RPAREN ;
extensions : SPACE extension (SPACE extension)*;

fragment D : ('d' | 'D') ;

DIGIT   : [0-9] ;
ALPHA   : [a-zA-Z] ;
WORD    : [0-9A-Za-z_] ;
SIGN    : (PLUS | MINUS) ;
LPAREN  : '(' ;
RPAREN  : ')' ;
DQUOTE  : '"' ;
COMMA   : ',' ;
PLUS    : '+' ;
MINUS   : '-' ;
SPACE   : ' ' ;




