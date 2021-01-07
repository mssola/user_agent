// Copyright (C) 2012-2020 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package user_agent

import (
	"regexp"
	"strings"
)

var ie11Regexp = regexp.MustCompile("^rv:(.+)$")

// browser is a struct containing all the information that we might be
// interested from the browser.
type browser struct {
	// The name of the browser's engine.
	engine string

	// The version of the browser's engine.
	engineVersion string

	// The name of the browser.
	name string

	// The version of the browser.
	version string
}

// Extract all the information that we can get from the User-Agent string
// about the browser and update the receiver with this information.
//
// The function receives just one argument "sections", that contains the
// sections from the User-Agent string after being parsed.
func (p *UserAgent) detectBrowser(sections []section) {
	slen := len(sections)

	if sections[0].name == "Opera" {
		p.browser.name = "Opera"
		p.browser.version = sections[0].version
		p.browser.engine = "Presto"
		if slen > 1 {
			p.browser.engineVersion = sections[1].version
		}
	} else if sections[0].name == "Dalvik" {
		// When Dalvik VM is in use, there is no browser info attached to ua.
		// Although browser is still a Mozilla/5.0 compatible.
		p.mozilla = "5.0"
	} else if slen > 1 {
		engine := sections[1]
		p.browser.engine = engine.name
		p.browser.engineVersion = engine.version
		if slen > 2 {
			sectionIndex := 2
			// The version after the engine comment is empty on e.g. Ubuntu
			// platforms so if this is the case, let's use the next in line.
			if sections[2].version == "" && slen > 3 {
				sectionIndex = 3
			}
			p.browser.version = sections[sectionIndex].version
			if engine.name == "AppleWebKit" {
				for _, comment := range engine.comment {
					if len(comment) > 5 &&
						(strings.HasPrefix(comment, "Googlebot") || strings.HasPrefix(comment, "bingbot")) {
						p.undecided = true
						break
					}
				}
				switch sections[slen-1].name {
				case "Edge":
					p.browser.name = "Edge"
					p.browser.version = sections[slen-1].version
					p.browser.engine = "EdgeHTML"
					p.browser.engineVersion = ""
				case "Edg":
					if p.undecided != true {
						p.browser.name = "Edge"
						p.browser.version = sections[slen-1].version
						p.browser.engine = "AppleWebKit"
						p.browser.engineVersion = sections[slen-2].version
					}
				case "OPR":
					p.browser.name = "Opera"
					p.browser.version = sections[slen-1].version
				default:
					switch sections[slen-3].name {
					case "YaBrowser":
						p.browser.name = "YaBrowser"
						p.browser.version = sections[slen-3].version
					default:
						switch sections[slen-2].name {
						case "Electron":
							p.browser.name = "Electron"
							p.browser.version = sections[slen-2].version
						default:
							switch sections[sectionIndex].name {
							case "Chrome", "CriOS":
								p.browser.name = "Chrome"
							case "Chromium":
								p.browser.name = "Chromium"
							case "FxiOS":
								p.browser.name = "Firefox"
							default:
								p.browser.name = "Safari"
							}
						}
					}
					// It's possible the google-bot emulates these now
					for _, comment := range engine.comment {
						if len(comment) > 5 &&
							(strings.HasPrefix(comment, "Googlebot") || strings.HasPrefix(comment, "bingbot")) {
							p.undecided = true
							break
						}
					}
				}
			} else if engine.name == "Gecko" {
				name := sections[2].name
				if name == "MRA" && slen > 4 {
					name = sections[4].name
					p.browser.version = sections[4].version
				}
				p.browser.name = name
			} else if engine.name == "like" && sections[2].name == "Gecko" {
				// This is the new user agent from Internet Explorer 11.
				p.browser.engine = "Trident"
				p.browser.name = "Internet Explorer"
				for _, c := range sections[0].comment {
					version := ie11Regexp.FindStringSubmatch(c)
					if len(version) > 0 {
						p.browser.version = version[1]
						return
					}
				}
				p.browser.version = ""
			}
		}
	} else if slen == 1 && len(sections[0].comment) > 1 {
		comment := sections[0].comment
		if comment[0] == "compatible" && strings.HasPrefix(comment[1], "MSIE") {
			p.browser.engine = "Trident"
			p.browser.name = "Internet Explorer"
			// The MSIE version may be reported as the compatibility version.
			// For IE 8 through 10, the Trident token is more accurate.
			// http://msdn.microsoft.com/en-us/library/ie/ms537503(v=vs.85).aspx#VerToken
			for _, v := range comment {
				if strings.HasPrefix(v, "Trident/") {
					switch v[8:] {
					case "4.0":
						p.browser.version = "8.0"
					case "5.0":
						p.browser.version = "9.0"
					case "6.0":
						p.browser.version = "10.0"
					}
					break
				}
			}
			// If the Trident token is not provided, fall back to MSIE token.
			if p.browser.version == "" {
				p.browser.version = strings.TrimSpace(comment[1][4:])
			}
		}
	}
}

// Engine returns two strings. The first string is the name of the engine and the
// second one is the version of the engine.
func (p *UserAgent) Engine() (string, string) {
	return p.browser.engine, p.browser.engineVersion
}

// Browser returns two strings. The first string is the name of the browser and the
// second one is the version of the browser.
func (p *UserAgent) Browser() (string, string) {
	return p.browser.name, p.browser.version
}
