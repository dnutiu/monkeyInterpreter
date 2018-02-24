package repl

import (
	"io"
	"bufio"
	"fmt"
	"github.com/metonimie/monkeyInterpreter/lexer"
	"github.com/metonimie/monkeyInterpreter/token"
)

const PROMPT = ">"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, "%s", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			goto exit
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "+%v\n", tok)
		}
	}
exit:
	return
}
