package config

import "github.com/joaosoft/golang-learn/52_nsq/common/nsq"

type Configuration struct {
	nsq.NSQ `json:"nsq"`
}
