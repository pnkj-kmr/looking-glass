package repos

import "encoding/json"

// IPInfo - helps to show ip address realted info
type IPInfo struct {
	ID     int    `json:"id",omitempty` // TODO - need to bind uid if need
	IP     string `json:"ip"`
	Host   string `json:"host"`
	Port   int 	  `json:"port"`
	Usr    string `json:"username"`
	Pwd    string `json:"password"`
	Vendor string `json:"vendor"`
}

// ToJSON - helps to convert model to bytes or string
func (ip *IPInfo) ToJSON() ([]byte, error) {
	return json.Marshal(ip)
}

// FromJSON - helps to convert bytes to model object
func (ip *IPInfo) FromJSON(data []byte) error {
	return json.Unmarshal(data, &ip)
}

// CountAudit - helps to record access and query audit
type CountAudit struct {
	Access int `json:"access_count"`
	Query  int `json:"query_count"`
}

// ToJSON - helps to convert model to bytes or string
func (ip *CountAudit) ToJSON() ([]byte, error) {
	return json.Marshal(ip)
}

// FromJSON - helps to convert bytes to model object
func (ip *CountAudit) FromJSON(data []byte) error {
	return json.Unmarshal(data, &ip)
}

// IncreaseAccess incerese the count by 1
func (ip *CountAudit) IncreaseAccess() {
	ip.Access = ip.Access + 1
}

// IncreaseQuery incerese the count by 1
func (ip *CountAudit) IncreaseQuery() {
	ip.Query = ip.Query + 1
}
