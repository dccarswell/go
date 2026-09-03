package try

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

func Shush(tgt any) func(f TryFunc) {
	return func(f TryFunc) {
		err := Try(f)
		if err != nil {
			switch t := tgt.(type) {
			case io.Writer:
				fmt.Fprintf(t, "Shush: suppressed error: %v\n", err)
			case *slog.Logger:
				t.Warn("Shush: suppressed error", "error", err)
			case *log.Logger:
				t.Printf("Shush: suppressed error: %v\n", err)
			case nil:
				//silently discard
			default:
				fmt.Fprintf(os.Stderr, "Shush: suppressed error: %v\n", err)
			}
		}
	}
}
