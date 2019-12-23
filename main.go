// wait-for-server attempts to contact the server of a URL.
package main

import "time"
import "net/http"
import "log"
import "fmt"
import "os"

func main() {
	for _, url := range os.Args[1:] {
		if err := WaitForServer(url); err != nil {
			fmt.Fprintf(os.Stdout, "Site is down: %v\n", err)
			continue
		}

		fmt.Fprintf(os.Stdout, "%s contacted successfully\n", url)
	}
}

// WaitForServer attempt to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}

		log.SetPrefix("wait: ")
		log.SetFlags(0)
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
