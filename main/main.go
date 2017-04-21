package main

import (
	"fmt"

	"github.com/vedicsoft/user_agent"
	"strconv"
)

func main() {
	// The "New" function will create a new UserAgent object and it will parse
	// the given string. If you need to parse more strings, you can re-use
	// this object and call: ua.Parse("another string")
	ua := user_agent.New("Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.23 Mobile Safari/537.36")

	fmt.Println("Device is a Mobile : "+ strconv.FormatBool(ua.Mobile()))   // => false
	fmt.Println("Is this a Bot : "+ strconv.FormatBool(ua.Bot()))      // => false
	fmt.Println("Mozilla version : "+ua.Mozilla())  // => "5.0"

	fmt.Println("Platform : " + ua.Platform()) // => "X11"
	fmt.Println("OS : " + ua.OS())       // => "Linux x86_64"

	name, version := ua.Engine()
	fmt.Println("Browser Engine : " + name)          // => "AppleWebKit"
	fmt.Println("Browser Engine Version: " + version)       // => "537.11"

	name, version = ua.Browser()
	fmt.Println("Browser Name : " +name)          // => "Chrome"
	fmt.Println("Browser Version : " +version)       // => "23.0.1271.97"

	// Let's see an example with a bot.
	ua.Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	fmt.Println("\nFollowing is an example what will return if we provide bot user agent\n")

	fmt.Println("Is this a Bot : "+ strconv.FormatBool(ua.Bot()))      // => true

	name, version = ua.Browser()
	fmt.Println("Browser Name : " +name)          // => Googlebot
	fmt.Println("Browser Version : " +version)       // => 2.1
}
