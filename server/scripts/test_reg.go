package main

import (
	"fmt"
	"regexp"
)

// var output []string = []string{
// 	"Info: The max number of VTY users is 21, the number of current VTY users online is 1, and total number of terminal users online is 1.",
// 	"wait: remote command exited without exit status or exit signal",
// 	"The current login time is 2023-05-09 11:10:15+05:30.",
// 	"Info: The device is not enabled with secure boot, please enable it.",
// 	"<NR-DELHI-NRTCC-IGW-01>",
// 	" ",
// 	" BGP local router ID : 103.120.29.90",
// 	" Local AS number : 132215",
// 	" Paths:   4 available, 1 best, 1 select, 0 best-external, 0 add-path",
// 	" BGP routing table entry information of 1.1.1.0/24:",
// 	" From: 103.77.108.118 (162.158.226.1)  ",
// 	" Route Duration: 0d00h26m10s",
// 	" Direct Out-interface: Eth-Trunk105.15",
// 	" Original nexthop: 103.77.108.118",
// 	" Qos information : 0x0",
// 	" Community: <49378:1>, <59177:59177>",
// 	" AS-path 13335, origin igp, localpref 150, pref-val 0, valid, external, best, select, pre 255, validation valid",
// 	" Aggregator: AS 13335, Aggregator ID 162.158.226.14",
// 	" Advertised to such 3 peers:",
// 	"    172.20.14.57",
// 	"    172.20.14.89",
// 	"    103.120.29.93",
// 	"",
// 	" BGP routing table entry information of 1.1.1.0/24:",
// 	" From: 103.77.108.11 (162.158.226.22)  ",
// 	" Route Duration: 0d00h26m10s",
// 	" Direct Out-interface: Eth-Trunk105.15",
// 	" Original nexthop: 103.77.108.11",
// 	" Qos information : 0x0",
// 	" Community: <49378:1>, <59177:59177>",
// 	" AS-path 13335, origin igp, localpref 150, pref-val 0, valid, external, pre 255, validation valid, not preferred for router ID",
// 	" Aggregator: AS 13335, Aggregator ID 162.158.226.14",
// 	" Not advertised to any peer yet",
// 	"",
// 	"  ---- More ----",
// 	"Info: The max number of VTY users is 21, and the number of current VTY users on line is 0.",
// }

func main() {
	test_reg()

	// for _, x := range output {
	// 	ok, y := checkUnwantedData(x)
	// 	if ok {
	// 		fmt.Println(y)
	// 	}
	// }
}

func test_reg() {
	// helps to filter only hostname and ip address for input value
	pattern := regexp.MustCompile(`^[a-zA-Z0-9:._-]+$`)
	// pattern := regexp.MustCompile("p([A-Za-z0-9]+)")

	fmt.Println(pattern.FindString("Golang"))
	fmt.Println(pattern.FindString("Golang2323"))
	fmt.Println(pattern.FindString("Golangfsdfs"))
	fmt.Println(pattern.FindString("GolangERTDFVHJ"))
	fmt.Println(pattern.FindString("@$#%^"))
	fmt.Println(pattern.FindString("121212.2.2.2:."))
	fmt.Println(pattern.FindString("-sdsd_SDsd_Sd9879*"))
	fmt.Println(pattern.FindString("343:454434::343453:3445344"))

	// func htlps to mask the IP address in give string...
	// Str helps to mask the give ip address in output str

	// str1 := `Proxy Port Last Check Proxy Speed Proxy Country Anonymity 118.99.81.204
	// 118.99.81.204 8080 34 sec Indonesia - Tangerang Transparent 2.184.31.2 8080 58 sec
	// Iran Transparent 93.126.11.189 8080 1 min Iran - Esfahan Transparent 202.118.236.130
	// 7777 1 min China - Harbin Transparent 62.201.207.9 8080 1 min Iraq Transparent`

	// re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	// re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	// fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	// fmt.Println(re.MatchString(str1))        // true

	// submatchall := re.FindAllString(str1, -1)
	// for _, element := range submatchall {
	// 	s := strings.Split(element, ".")
	// 	s1 := s[0] + ".X.X.X"
	// 	fmt.Println(element, strings.Replace(str1, element, s1, -1))
	// }
	// fmt.Println(maskingIP(str1))

}

// func checkUnwantedData(str string) (bool, string) {
// 	if strings.HasPrefix(str, "Info: ") {
// 		return false, ""
// 	}
// 	if strings.HasPrefix(str, "The current login time") {
// 		return false, ""
// 	}
// 	if strings.HasPrefix(str, "<") && strings.HasSuffix(str, ">") {
// 		return false, ""
// 	}
// 	if strings.Contains(str, "exited without exit status or exit signal") {
// 		return false, ""
// 	}
// 	if strings.Contains(str, "Community") {
// 		return true, str[:strings.Index(str, "Community")+10] + " XXXX"
// 	}
// 	if strings.Contains(str, "Direct Out-interface") {
// 		return true, str[:strings.Index(str, "Direct Out-interface")+21] + " XXXX"
// 	}
// 	return true, str
// }

// var ipRegex *regexp.Regexp = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

// func maskingIP(str string) string {
// 	if ipRegex.MatchString(str) {
// 		var newStr string = str
// 		submatchall := ipRegex.FindAllString(str, -1)
// 		for _, element := range submatchall {
// 			_masked := strings.Split(element, ".")[0] + ".X.X.X"
// 			newStr = strings.Replace(newStr, element, _masked, -1)
// 		}
// 		return newStr
// 	}
// 	return str
// }
