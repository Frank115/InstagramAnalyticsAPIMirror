package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/ProyectoIT/InstagramAnalyticsAPI/config"
)

func getRhxGis() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.instagram.com/", nil)
	req.Header.Add("User-Agent", config.HTTPClientUserAgent)
	res, _ := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)
	r := regexp.MustCompile("(\"rhx_gis\":.[^\"]+)")
	rhxgis := strings.Replace(r.FindString(string(body)), "\"rhx_gis\":\"", "", 1)

	return rhxgis
}

func GetInstagramGisHash(query string) string {
	hash := md5.New()
	hash.Write([]byte(fmt.Sprintf("%v:%v", getRhxGis(), query)))
	return hex.EncodeToString(hash.Sum(nil))
}
