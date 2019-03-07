package rpcdemo

import "strings"

type Hello struct {
}

func (Hello) Upper(args string, reply *string) error {
	*reply = strings.ToUpper(args)
	return nil
}
