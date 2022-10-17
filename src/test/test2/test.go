package main

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"os"
	"strings"
	"time"
)

// StringPrompt asks for a string value using the label
func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		//fmt.Fprint(os.Stderr, label+"\n")
		fmt.Println(label)
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func scannerTest() {

}

func testPb() {
	count := 100

	// create and start new bar
	//bar := pb.StartNew(count)

	// start bar from 'default' template
	//bar := pb.Default.Start(count)

	// start bar from 'simple' template
	bar := pb.Simple.Start(count)

	// start bar from 'full' template
	// bar := pb.Full.Start(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(200 * time.Millisecond)
	}

	// finish bar
	bar.Finish()
}

func main() {
}
