package controllers

// Protocol contains mapping
var Protocol map[int]string = map[int]string{
	1: "PING IPv4",
	2: "PING IPv6",
	3: "TRACEROUTE IPv4",
	4: "TRACEROUTE IPv6",
	5: "BGP IPv4",
	6: "BGP IPv6",
}

// Vendor contains mapping
var Vendor map[int]string = map[int]string{
	1: "Cisco",
	2: "Huawei",
	0: "Other",
}

// Proto holds the display type
type Proto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProtoSorter defing the sorting
type ProtoSorter []Proto

func (a ProtoSorter) Len() int           { return len(a) }
func (a ProtoSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ProtoSorter) Less(i, j int) bool { return a[i].ID < a[j].ID }
