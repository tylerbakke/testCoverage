package person

import (
	"github.com/tylerbakke/testCoverage/work"
)

type Employee struct {
	name string
	job work.Job
}

func CreateEmployee(job work.Job) Employee {
	var newEmployee Employee
	newEmployee.name = "Mindless Drone"
	newEmployee.job = job
	return newEmployee
}

func (e Employee) GetName() string {
	return e.name
}

func (e Employee) GetJob() string {
	return e.job.GetTitle()
}