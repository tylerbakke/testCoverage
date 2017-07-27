package person

import (
	"github.com/tylerbakke/testCoverage/work"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	job := work.CreateJob("Intern")
	employee := CreateEmployee(job)

	if employee.job.GetTitle() != "Intern" {
		t.Errorf("ERROR: Failed assignment")
	}

}