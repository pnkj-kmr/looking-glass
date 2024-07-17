package controllers

import (
	"github.com/pnkj-kmr/looking-glass/repos"
	"github.com/pnkj-kmr/looking-glass/repos/jsondb"
	"github.com/pnkj-kmr/looking-glass/utils"
	"go.uber.org/zap"
)

const (
	ipDetailTable string = "ipinfo"
	auditTable    string = "audit"
)

// ListIPConfig helps to store the src ig detail locally
func ListIPConfig(key string) ([]repos.IPInfo, error) {
	utils.L.Debug("Getting ip list", zap.String("ip", key))
	var ips []repos.IPInfo
	collection, err := jsondb.New(ipDetailTable)
	if err != nil {
		return ips, err
	}
	utils.L.Debug("Collection found")
	if key == "" {
		data := collection.GetAll()
		for _, d := range data {
			var ip repos.IPInfo
			_ = ip.FromJSON(d)
			if newPwd, err := decryptAES(ip.Pwd); err != nil {
				continue
			} else {
				ip.Pwd = maskPasswd(newPwd)
			}
			ips = append(ips, ip)
		}
	} else {
		d, err := collection.Get(key)
		if err != nil {
			return ips, err
		}
		var ip repos.IPInfo
		_ = ip.FromJSON(d)
		newPwd, err := decryptAES(ip.Pwd)
		if err != nil {
			return ips, err
		}
		ip.Pwd = maskPasswd(newPwd)
		ips = append(ips, ip)

	}
	utils.L.Debug("IPs found", zap.Int("ips", len(ips)))
	return ips, nil
}

// AddIPConfig helps to store the src ig detail locally
func AddIPConfig(data []byte) error {
	utils.L.Debug("Adding ip configuration", zap.ByteString("ip", data))
	var ip repos.IPInfo
	if err := ip.FromJSON(data); err != nil {
		return err
	}
	newPwd, err := encryptAES(ip.Pwd)
	if err != nil {
		return err
	}
	ip.Pwd = newPwd
	collection, err := jsondb.New(ipDetailTable)
	if err != nil {
		return err
	}
	d, err := ip.ToJSON()
	if err != nil {
		return err
	}
	if err := collection.Insert(ip.IP, d); err != nil {
		return err
	}
	utils.L.Debug("Record insertion succeed")
	return nil
}

// DeleteIPConfig helps to store the src ig detail locally
func DeleteIPConfig(key string) error {
	utils.L.Debug("Deleting ip configuration", zap.String("ip", key))
	collection, err := jsondb.New(ipDetailTable)
	if err != nil {
		return err
	}
	if err := collection.Delete(key); err != nil {
		return err
	}
	utils.L.Debug("Record deletion succeed")
	return nil
}

// EditIPConfig helps to store the src ig detail locally
func EditIPConfig(key string, data []byte) error {
	utils.L.Debug("Updating ip configuration", zap.String("ip", key), zap.ByteString("ip_edit", data))
	collection, err := jsondb.New(ipDetailTable)
	if err != nil {
		return err
	}
	// existing entry from db
	ex, err := collection.Get(key)
	if err != nil {
		return err
	}
	var exip repos.IPInfo
	_ = exip.FromJSON(ex)
	// delete existing from db
	if err := collection.Delete(key); err != nil {
		return err
	}
	var ip repos.IPInfo
	if err := ip.FromJSON(data); err != nil {
		return err
	}
	newPwd, err := encryptAES(ip.Pwd)
	if err != nil {
		return err
	}
	// cross checking password masking if same
	// using existing password
	mPass := maskPasswd(newPwd)
	if mPass == newPwd {
		ip.Pwd = exip.Pwd
	} else {
		ip.Pwd = newPwd
	}
	d, err := ip.ToJSON()
	if err != nil {
		return err
	}
	if err := collection.Insert(ip.IP, d); err != nil {
		return err
	}
	utils.L.Debug("Record updation succeed")
	return nil
}

// GetAccessData to get the access detailed data
func GetAccessData() (data repos.CountAudit, err error) {
	collection, err := jsondb.New(auditTable)
	if err != nil {
		return data, err
	}
	// fixed key "access" to save or retrive data
	d, err := collection.Get("access")
	if err != nil {
		return data, err
	}
	if err = data.FromJSON(d); err != nil {
		return data, err
	}
	return data, nil
}

// SetAccessData save access audit data into db
func SetAccessData(t int) error {
	collection, err := jsondb.New(auditTable)
	if err != nil {
		return err
	}
	// fixed key "access" to save or retrive data
	data, err := collection.Get("access")
	if err != nil {
		audit := repos.CountAudit{Access: 1, Query: 1}
		data, err = audit.ToJSON()
		if err != nil {
			return err
		}
	} else {
		if err := collection.Delete("access"); err != nil {
			return err
		}
	}
	var audit repos.CountAudit
	err = audit.FromJSON(data)
	if err != nil {
		return err
	}
	// increment by 1
	if t == 1 {
		audit.IncreaseAccess()
	} else {
		audit.IncreaseQuery()
	}

	dataBytes, err := audit.ToJSON()
	if err != nil {
		return err
	}
	if err := collection.Insert("access", dataBytes); err != nil {
		return err
	}
	return nil
}
