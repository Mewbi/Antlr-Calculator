# 📊 Simple Calculator with ANTLR

This project is a simple calculator written in Golang that uses ANTLR for parsing expressions. It supports variables, basic arithmetic operations, and can run in three different modes.

## 🛠️ Features

- 📝 **Interactive Prompt Mode**: Enter and evaluate expressions line by line.
- 📂 **File Mode**: Read and evaluate expressions from a file.
- 🧮 **Expression Mode**: Evaluate an expression passed as a command-line argument.
- ℹ️  **Help Menu**: Provides usage information and instructions.

## 📜 Grammar

The calculator uses the following ANTLR grammar:

```antlr
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
```

Obs. Relational operators will return 1 if expression is true, and 0 if false.

## 💾 Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Mewbi/Antlr-Calculator.git
   ```
2. Build the program:
   ```sh
   go build -o ./calc main.go
   ```

3. Run the program:
   ```sh
   ./calc
   ```

## 📋 Usage

### Interactive Prompt Mode

Run the program without any arguments to enter the interactive prompt mode:

```sh
./calc
```

Type your expressions, and the program will print the results. Type `exit` to quit.

### File Mode

Use the `-file` flag to read expressions from a file:

```sh
./calc -file filename
```

### Expression Mode

Pass an expression as a command-line argument to evaluate it:

```sh
./calc "2 + 3 * 6 + (5 + 2)"
```

### Help Menu

Use the `-help` flag to display the help menu:

```sh
./calc -help
```

## 🔍 Example

Here's an example of the calculator in action:

```sh
$ ./calc
> x = 5
> y = 10
> z = x + y * 2
> z
25
> exit
```

Enjoy using the simple calculator! 🎉
