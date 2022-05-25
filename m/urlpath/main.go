package urlpath

import(
	"regexp"
)

var(
	RootPrefix = "/"
	StaticPrefix = RootPrefix+"s/"
	ViewPostPrefix = RootPrefix+"vp/"
	TypePostPrefix = RootPrefix+"tp/"
	EditPostPrefix = RootPrefix+"ep/"
	TypePostHndlPrefix = RootPrefix+"tph/"
	GetTestPrefix = RootPrefix+"gettest/"
	PostTestPrefix = RootPrefix+"posttest/"
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

