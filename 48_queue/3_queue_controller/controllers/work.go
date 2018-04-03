package controllers

type Work struct {
	data interface{}
}

func NewWork(work interface{}) *Work {
	Work := Work{
		data: work,
	}

	return &Work
}

func (work *Work) GetWork() (interface{}, error) {
	return work.data, nil
}
