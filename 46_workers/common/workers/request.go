package workers

type WorkRequest struct {
	Controller IController `json:"controllers"`
}