package api

type API struct {
	Job *Job
}

type Job struct {
	Create JobCreate
}
