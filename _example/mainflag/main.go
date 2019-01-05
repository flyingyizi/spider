package mainflag

import "flag"
import "fmt"
import "strconv"

type percentage float32

func (p *percentage) Set(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	*p = percentage(v)
	return err
}
func (p *percentage) String() string { return fmt.Sprintf("%f", *p) }

//myFlags should be the
func myFlags() []string {
	//namePtr := flag.String("name", "lyh", "user's name")
	flag.Parse()
	others := flag.Args()
	//fmt.Println("name:", *namePtr)
	fmt.Println("other:", others)
	return others
}
