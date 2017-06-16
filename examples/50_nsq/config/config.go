package config

import "golang-learn/examples/50_nsq/common/nsq"

type Configuration struct {
	nsq.NSQ       `json:"nsq"`
}