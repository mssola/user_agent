package user_agent

import (
	"strings"
)

// detectModel some properties of the model from the given section.
func (p *UserAgent) detectModel(s section) {
	if !p.mobile {
		return
	}
	if p.platform == "iPhone" || p.platform == "iPad" {
		p.model = p.platform
		return
	}
	// Android model
	if s.name == "Mozilla" && p.platform == "Linux" && len(s.comment) > 2 {
		mostAndroidModel := s.comment[2]
		if strings.Contains(mostAndroidModel, "Android") || strings.Contains(mostAndroidModel, "Linux") {
			mostAndroidModel = s.comment[len(s.comment)-1]
		}
		tmp := strings.Split(mostAndroidModel, "Build")
		if len(tmp) > 0 {
			p.model = strings.Trim(tmp[0], " ")
			return
		}
	}
	// traverse all item
	for _, v := range s.comment {
		if strings.Contains(v, "Build") {
			tmp := strings.Split(v, "Build")
			p.model = strings.Trim(tmp[0], " ")
		}
	}
}
