package sanitize

import(
	"github.com/microcosm-cc/bluemonday"
)

func
Sanitize(input []byte) []byte {
	ret := bluemonday.UGCPolicy().SanitizeBytes(input)
	return ret
}
