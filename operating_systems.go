// Copyright (C) 2012-2014 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package user_agent

import "strings"

// Internal: normalize the name of the operating system. By now, this just
// affects to Windows.
//
// Returns a string containing the normalized name for the Operating System.
func normalizeOS(name string) string {
	sp := strings.Split(name, " ")
	if len(sp) < 3 {
		return name
	}

	switch sp[2] {
	case "5.0":
		return "Windows 2000"
	case "5.01":
		return "Windows 2000, Service Pack 1 (SP1)"
	case "5.1":
		return "Windows XP"
	case "5.2":
		return "Windows XP x64 Edition"
	case "6.0":
		return "Windows Vista"
	case "6.1":
		return "Windows 7"
	case "6.2":
		return "Windows 8"
	case "6.3":
		return "Windows 8.1"
	}
	return name
}

// Internal: guess the OS, the localization and if this is a mobile device
// for a Webkit-powered browser.
//
// The first argument p is a reference to the current UserAgent and the second
// argument is a slice of strings containing the comment.
func webkit(p *UserAgent, comment []string) {
	if p.platform == "webOS" {
		p.browser.name = p.platform
		p.os = "Palm"
		if len(comment) > 2 {
			p.localization = comment[2]
		}
		p.mobile = true
	} else if p.platform == "Symbian" {
		p.mobile = true
		p.browser.name = p.platform
		p.os = comment[0]
	} else if p.platform == "Linux" {
		p.mobile = true
		if p.browser.name == "Safari" {
			p.browser.name = "Android"
		}
		if len(comment) > 1 {
			if comment[1] == "U" {
				if len(comment) > 2 {
					p.os = comment[2]
				} else {
					p.mobile = false
					p.os = comment[0]
				}
			} else {
				p.os = comment[1]
			}
		}
		if len(comment) > 3 {
			p.localization = comment[3]
		}
	} else if len(comment) > 0 {
		if len(comment) > 3 {
			p.localization = comment[3]
		}
		if strings.HasPrefix(comment[0], "Windows NT") {
			p.os = normalizeOS(comment[0])
		} else if len(comment) < 2 {
			p.localization = comment[0]
		} else if len(comment) < 3 {
			p.os = normalizeOS(comment[1])
		} else {
			p.os = normalizeOS(comment[2])
		}
		if p.platform == "BlackBerry" {
			p.browser.name = p.platform
			if p.os == "Touch" {
				p.os = p.platform
			}
		}
	}
}

// Internal: guess the OS, the localization and if this is a mobile device
// for a Gecko-powered browser.
//
// The first argument p is a reference to the current UserAgent and the second
// argument is a slice of strings containing the comment.
func gecko(p *UserAgent, comment []string) {
	if len(comment) > 1 {
		if comment[1] == "U" {
			if len(comment) > 2 {
				p.os = normalizeOS(comment[2])
			} else {
				p.os = normalizeOS(comment[1])
			}
		} else {
			if p.platform == "Android" {
				p.mobile = true
				p.platform, p.os = normalizeOS(comment[1]), p.platform
			} else if len(comment) < 3 {
				p.mobile = true
				p.os = "FirefoxOS"
			} else {
				p.os = normalizeOS(comment[1])
			}
		}
		if len(comment) > 3 {
			p.localization = comment[3]
		}
	}
}

// Internal: guess the OS, the localization and if this is a mobile device
// for Internet Explorer.
//
// The first argument p is a reference to the current UserAgent and the second
// argument is a slice of strings containing the comment.
func trident(p *UserAgent, comment []string) {
	// Internet Explorer only runs on Windows.
	p.platform = "Windows"

	// The OS can be set before to handle a new case in IE11.
	if p.os == "" {
		if len(comment) > 2 {
			p.os = normalizeOS(comment[2])
		} else {
			p.os = "Windows NT 4.0"
		}
	}

	// Last but not least, let's detect if it comes from a mobile device.
	for _, v := range comment {
		if strings.HasPrefix(v, "IEMobile") {
			p.mobile = true
			return
		}
	}
}

// Internal: guess the OS, the localization and if this is a mobile device
// for Opera.
//
// The first argument p is a reference to the current UserAgent and the second
// argument is a slice of strings containing the comment.
func opera(p *UserAgent, comment []string) {
	if strings.HasPrefix(comment[0], "Windows") {
		p.platform = "Windows"
		p.os = normalizeOS(comment[0])
		if len(comment) > 2 {
			p.localization = comment[2]
		}
	} else {
		if strings.HasPrefix(comment[0], "Android") {
			p.mobile = true
		}
		p.platform = comment[0]
		if len(comment) > 1 {
			p.os = comment[1]
			if len(comment) > 3 {
				p.localization = comment[3]
			}
		} else {
			p.os = comment[0]
		}
	}
}

// Internal: given the comment of the first section of the UserAgent string,
// get the platform.
//
// Returns a string that is the platform.
func getPlatform(comment []string) string {
	if len(comment) > 0 {
		if comment[0] != "compatible" {
			if strings.HasPrefix(comment[0], "Windows") {
				return "Windows"
			} else if strings.HasPrefix(comment[0], "Symbian") {
				return "Symbian"
			} else if strings.HasPrefix(comment[0], "webOS") {
				return "webOS"
			} else if comment[0] == "BB10" {
				return "BlackBerry"
			}
			return comment[0]
		}
	}
	return ""
}

// Internal: detect some properties of the OS from the given section.
func (p *UserAgent) detectOS(section UASection) {
	if section.name == "Mozilla" {
		// Get the platform here. Be aware that IE11 provides a new format
		// that is not backwards-compatible with previous versions of IE.
		p.platform = getPlatform(section.comment)
		if p.platform == "Windows" && len(section.comment) > 0 {
			p.os = normalizeOS(section.comment[0])
		}

		// And finally get the OS depending on the engine.
		switch p.browser.engine {
		case "Gecko":
			gecko(p, section.comment)
		case "AppleWebKit":
			webkit(p, section.comment)
		case "Trident":
			trident(p, section.comment)
		}
	} else if section.name == "Opera" {
		if len(section.comment) > 0 {
			opera(p, section.comment)
		}
	} else {
		// At this point, we assume that this is a bot.
		p.bot = true
		p.mobile = true
		p.mozilla = ""
		p.browser.name = section.name
		p.browser.version = section.version
	}
}

// Returns a string containing the platform..
func (p *UserAgent) Platform() string {
	return p.platform
}

// Returns a string containing the name of the Operating System.
func (p *UserAgent) OS() string {
	return p.os
}

// Returns a string containing the localization.
func (p *UserAgent) Localization() string {
	return p.localization
}
