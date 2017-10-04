package config

import "golang-learn/52_nsq/common/nsq"

type Configuration struct {
	nsq.NSQ       `json:"nsq"`
}