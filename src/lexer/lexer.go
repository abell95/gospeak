package lexer

// Lexer processes the text into tokens
type Lexer struct {
	input        string
	position     int  // current character in input string
	readPosition int  // reading position in input
	ch           byte // the character being read
}

// New initializes and returns a lexer
func New(input string) *Lexer {
	l := &Lexer{input: input} // init the lexer, feed it the input and assign it as a pointer
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // if the readPos is greater than the length of the input, we're done with the string
	} else {
		l.ch = l.input[l.readPosition] // otherwise, assign the next character to be analyzed
	}
	l.position = l.readPosition // keep track of our reading position
	l.readPosition++            // and then increment it
}
