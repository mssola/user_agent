// Copyright (C) 2012-2013 Miquel Sabaté Solà
// This file is licensed under the MIT license.
// See the LICENSE file.

// Package user_agent implements an HTTP User Agent string parser. It defines
// the type UserAgent that contains all the information from the parsed string.
// It also implements the Parse function and getters for all the relevant
// information that has been extracted from a parsed User Agent string.
package user_agent

import (
    "strings"
    "regexp"
)


// A "section" of the User-Agent string. A section contains the name of the
// product, its version and an optional comment.
type UASection struct {
    name    string
    version string
    comment []string
}

// The UserAgent struct contains all the info that can be extracted
// from the User-Agent string.
type UserAgent struct {
    mozilla      string
    platform     string
    os           string
    localization string
    browser      Browser
    bot          bool
    mobile       bool
}

// Internal: read from the given string until the given delimiter or the
// end of the string have been reached.
//
// ua        - The User-Agent string.
// index     - A reference to the current index on the User-Agent string,
// delimiter - A byte containing the given delimiter.
// cat       - Determines whether nested '(' should be ignored or not.
//
// Returns an array of bytes containing what has been read.
func readUntil(ua string, index *int, delimiter byte, cat bool) []byte {
    var buffer []byte

    i := *index
    catalan := 0
    for ; i < len(ua); i = i + 1 {
        if ua[i] == delimiter {
            if catalan == 0 {
                *index = i + 1
                return buffer
            }
            catalan--
        } else if cat && ua[i] == '(' {
            catalan++
        }
        buffer = append(buffer, ua[i])
    }
    *index = i + 1
    return buffer
}

// Internal: parse the product, that is, just a name or a string
// formatted as Name/Version.
//
// product - a slice of bytes containing the product.
//
// Returns:
//     - a string containing the name of the product.
//     - a string containing the version of the product.
func parseProduct(product []byte) (string, string) {
    prod := strings.Split(string(product), "/")
    if len(prod) == 2 {
        return prod[0], prod[1]
    }
    return string(product), ""
}

// Internal: parse a section. A section is typically formatted as
// follows "Name/Version (comment)". Both, the comment and the version
// are optional.
//
// ua    - the UserAgent string.
// index - a reference to the index we're using inside the ua string.
//
// Returns a UASection containing the information that we can extract
// from the last parsed section.
func parseSection(ua string, index *int) (section UASection) {
    buffer := readUntil(ua, index, ' ', false)

    section.name, section.version = parseProduct(buffer)
    if *index < len(ua) && ua[*index] == '(' {
        *index++
        buffer = readUntil(ua, index, ')', true)
        section.comment = strings.Split(string(buffer), "; ")
        *index++
    }
    return section
}

// Public: parse a User-Agent string. After calling this function, the
// receiver will be setted up with all the information that we've extracted.
//
// ua - a string containing the User-Agent from the browser (or the bot).
func (p *UserAgent) Parse(ua string) {
    var sections []UASection

    p.mobile = false
    for index, limit := 0, len(ua); index < limit; {
        s := parseSection(ua, &index)
        if !p.mobile && s.name == "Mobile" {
            p.mobile = true
        }
        sections = append(sections, s)
    }

    if len(sections) > 0 {
        p.mozilla = sections[0].version
        if !p.bot {
            p.checkBot(sections[0].comment)
            p.detectBrowser(sections)
            p.detectOS(sections[0])
        }
    }
}

// Internal: check if we're dealing with a Bot.
//
// comment - A string containing the comment from the first section.
func (p *UserAgent) checkBot(comment []string) {
    reg, _ := regexp.Compile("(?i)bot")
    for _, v := range comment {
        if reg.Match([]byte(v)) {
            p.bot = true
            return
        }
    }
}

// Public: get the mozilla version (it's how the User Agent string begins:
// "Mozilla/5.0 ...", unless we're dealing with Opera, of course).
// Returns a string containing the mozilla version.
func (p *UserAgent) Mozilla() string {
    return p.mozilla
}

// Returns true if it's a bot, false otherwise.
func (p *UserAgent) Bot() bool {
    return p.bot
}

// Returns true if it's a mobile device, false otherwise.
func (p *UserAgent) Mobile() bool {
    return p.mobile
}
