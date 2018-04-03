package config

import "go-learn/52_nsq/common/nsq"

type Configuration struct {
	nsq.NSQ `json:"nsq"`
}
