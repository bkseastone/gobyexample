//go:generate protoc -I ./protos --go_out ./protos *.proto
package rpcdemo

import "strings"

type Hello struct {
}

func (Hello) Upper(args string, reply *string) error {
	*reply = strings.ToUpper(args)
	return nil
}
