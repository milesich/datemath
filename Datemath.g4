grammar Datemath;

start: e = timestamp EOF;

timestamp
	: Builtins                             # Builtin
	| timestamp op=(Plus | Minus) duration # Binary
	| timestamp Round duration             # Round
	| Identifier                           # Identifier
	| Literal                              # Literal
	| DateLiteral                          # DateLiteral
;

duration
	: EsDuration
	| GoDuration
;

/*****************************/
/*                           */
/*           LEXER           */
/*                           */
/*****************************/

EsDuration:
	IntegerLiteral? (
		'y'
		| 'M'
		| 'w'
		| 'd'
		| 'h'
		| 'H'
		| 'm'
		| 's'
	);

GoDuration: (
		(IntegerLiteral | FloatLiteral) (
			'h'
			| 'm'
			| 's'
			| 'ms'
			| 'us'
			| 'µs'
			| 'ns'
		)
	)+;

Literal: '\'' ~[']* '\'';

Plus: '+';
Minus: '-';
Round: '/';
Builtins: ( 'now' );

IntegerLiteral: '0' | [1-9] [0-9_]*;

FloatLiteral: DecimalLiteral '.' Digit+ | '.' Digit+;

Identifier: IdentifierStart IdentifierPart*;

WhiteSpaces: [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN);
DateLiteral:
	Digit Digit Digit Digit DateSeparator Digit? Digit DateSeparator Digit? Digit .*? ('||' | EOF);
fragment DateSeparator: ('-' | '/' );
fragment IdentifierStart: Letter;
fragment IdentifierPart: Letter | Digit | [_];
fragment Letter: 'A' ..'Z' | 'a' ..'z';
fragment Digit: '0' ..'9';
fragment DecimalLiteral: '0' | [1-9] Digit*;