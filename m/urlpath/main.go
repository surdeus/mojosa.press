package urlpath

import(
	"regexp"
)

var(
	RootPrefix = "/"
	StaticPrefix = RootPrefix+"s/"
	ViewPostPrefix = RootPrefix+"vp/"
	TypePostPrefix = RootPrefix+"tp/"
	TypePostHndlPrefix = RootPrefix+"tph/"
	TestPrefix = RootPrefix+"test/"
)

func
Validify(u string, re *regexp.Regexp) bool {
	if re == nil {
		return true
	}
	
	ret := re.Find([]byte(u))
	if ret == nil {
		return false
	}

	return true
}

