package urlexpl

import (
	"fmt"
	"net"
	"net/url"

	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type UrlExpl struct{}

func (ue *UrlExpl) RunExample(inputParams *goexpl.InputParams) error {
	// here is a url example
	// protocol://user:password@domain:port/path?query#fragment
	// * protocol aka schema
	// * domain aka host
	// * query are kv pairs, aka parameters
	s := "https://user:pass@host.com:8080/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	fmt.Println(u.Scheme)                // protocol/schema
	fmt.Println(u.User)                  // username:pass
	fmt.Println(u.User.Username())       // only username
	p, havePassword := u.User.Password() // get password and if there is a password
	fmt.Println(p, havePassword)

	// Host contains host and port, sometimes there are no port or default port?
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// path and fragment
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// get query, or should I say parameters
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

	return nil
}

func (ue *UrlExpl) Init() {
	err := goexpl.RegisterGoExample(ue.GetGoExampleName(), &UrlExpl{})
	if err != nil {
		panic(err)
	}
}

func (ue *UrlExpl) GetGoExampleName() string {
	return common.UrlExpl
}
