package work

import "testing"

func TestCreateJob(t *testing.T) {
	testJob := CreateJob("QA")
	if testJob.title != "QA" {
		t.Errorf("ERROR: Failed assignment")
	}

}

func TestCreateJobs(t *testing.T) {
	list := CreateJobs()
	if len(list) != 4 {
		t.Errorf("ERROR: Failed assignment")
	}
}