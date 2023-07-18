package netx

import (
	"bytes"
	"github.com/ml444/gutil/str"
	"net/url"
	"sort"
)

func GenSign(args url.Values, nonce string) string {
	keys := make([]string, 0, len(args))
	for k := range args {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))

	buffer := bytes.NewBuffer(nil)
	for _, k := range keys {
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(args.Get(k))
	}

	buffer.WriteString(nonce)

	return str.Md5(buffer.String())
}
