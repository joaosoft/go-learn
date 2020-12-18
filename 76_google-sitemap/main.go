package main

import (
	"database/sql"
	"encoding/xml"
	"flag"
	"fmt"
	_ "github.com/lib/pq" // postgres driver
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const (
	defaultOutputFile = "sitemap.xml"
	configFile        = "config.yaml"
)

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Text    string   `xml:",chardata"`
	Url     []Url    `xml:"url"`
}

type Url struct {
	Text    string    `xml:",chardata"`
	Loc     string    `xml:"loc"`
	Lastmod time.Time `xml:"lastmod"`
}

type Configuration struct {
	Database struct {
		Dialect    string `yaml:"dialect"`
		DataSource string `yaml:"data_source"`
	} `json:"database"`
	WebsiteUrl string `yaml:"website_url"`
	OutputFile string `yaml:"output_file"`
}

func main() {
	var config Configuration
	body, err := ioutil.ReadFile(fmt.Sprintf("./%s", configFile))
	if err != nil {
		panic(fmt.Sprintf("error reading config file %s: %s", configFile, err.Error()))
	}

	if err = yaml.Unmarshal(body, &config); err != nil {
		panic(fmt.Sprintf("error unmarshall config file %s: %s", configFile, err.Error()))
	}

	if config.OutputFile == "" {
		config.OutputFile = defaultOutputFile
	}

	flag.StringVar(&config.OutputFile, "file", config.OutputFile, "the file name to save")
	flag.Parse()

	if !strings.HasSuffix(config.OutputFile, ".xml") {
		config.OutputFile = fmt.Sprintf("%s.xml", config.OutputFile)
	}

	log.Printf("connecting database with driver [ %s ] and data source [ %s ]", config.Database.Dialect, config.Database.DataSource)
	db, err := sql.Open(config.Database.Dialect, config.Database.DataSource)
	if err != nil {
		panic(fmt.Sprintf("error connecting to database: %s", err.Error()))
	}

	rows, err := db.Query(`
		SELECT $1 || ws.website, now() FROM manufacturer.manufacturer AS m
		JOIN manufacturer.website_settings AS ws ON ws.fk_manufacturer = m.id_manufacturer 
		JOIN auth.entity AS e ON e.external_id = m.id_manufacturer 
		JOIN crm.customer AS c ON c.fk_entity = e.id_entity 
		JOIN crm.subscription AS s ON s.fk_customer = c.id_customer 
		JOIN crm.plan AS p ON p.id_plan = s.fk_plan 
		JOIN crm.plan_limit AS pl ON pl.fk_plan = p.id_plan 
		WHERE e.fk_sub_domain = 2
			AND pl.key = 'access-website'
			AND m.fk_manufacturer_status IN (1,3)
			AND p.active 
			AND pl.active
			AND ws.published
	`, fmt.Sprintf("%s/", config.WebsiteUrl))
	if err != nil {
		panic(fmt.Sprintf("error getting data from database: %s", err.Error()))
	}

	defer rows.Close()
	urlSet := UrlSet{}

	for rows.Next() {
		url := Url{}
		if err := rows.Scan(&url.Loc, &url.Lastmod); err != nil {
			panic(fmt.Sprintf("error getting scanning data from database: %s", err.Error()))
		}
		urlSet.Url = append(urlSet.Url, url)
	}

	body, err = xml.MarshalIndent(urlSet, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshall urlset: %s", err.Error()))
	}

	if err := ioutil.WriteFile(config.OutputFile, body, 0644); err != nil {
		panic(fmt.Sprintf("error saving file %s: %s", config.OutputFile, err.Error()))
	}

	log.Println(fmt.Sprintf("saved %d websites", len(urlSet.Url)))
}
