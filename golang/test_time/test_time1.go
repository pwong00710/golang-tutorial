package main

import (
	"fmt"
	"time"
)

// @link https://golang.org/pkg/time/

func main() {

	//caution : format string is `2006-01-02 15:04:05.000000000`
	current := time.Now()

	fmt.Println("unix : ", current.Unix())
	// unix :

	fmt.Println("origin : ", current.String())
	// origin :  2016-09-02 15:53:07.159994437 +0800 CST

	fmt.Println("mm-dd-yyyy : ", current.Format("01-02-2006"))
	// mm-dd-yyyy :  09-02-2016

	fmt.Println("yyyy-mm-dd : ", current.Format("2006-01-02"))
	// yyyy-mm-dd :  2016-09-02

	// separated by .
	fmt.Println("yyyy.mm.dd : ", current.Format("2006.01.02"))
	// yyyy.mm.dd :  2016.09.02

	fmt.Println("yyyy-mm-dd HH:mm:ss : ", current.Format("2006-01-02 15:04:05"))
	// yyyy-mm-dd HH:mm:ss :  2016-09-02 15:53:07

	// StampMicro
	fmt.Println("yyyy-mm-dd HH:mm:ss: ", current.Format("2006-01-02 15:04:05.000000"))
	// yyyy-mm-dd HH:mm:ss:  2016-09-02 15:53:07.159994

	//StampNano
	fmt.Println("yyyy-mm-dd HH:mm:ss: ", current.Format("2006-01-02 15:04:05.000000000"))
	// yyyy-mm-dd HH:mm:ss:  2016-09-02 15:53:07.159994437

	var duration int
	duration = 600
	current = current.Add(time.Duration(duration) * time.Second)
	//current = current.Add(600 * time.Hour)
	fmt.Println("yyyy-mm-dd HH:mm:ss: ", current.Format("2006-01-02 15:04:05.000000000"))
}
