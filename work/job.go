/***

This package creates jobs (Sorry DT) either individually with a given title
or by reading from a text file.

***/
package work

type Job struct {
	title string
}

func CreateJob(title string) Job{
	var newJob Job
	newJob.title = title
	return newJob
}

func CreateJobs() []Job {
	list := make([]Job, 4)
	list[0] = CreateJob("Programmer")
	list[1] = CreateJob("Marketing")
	list[2] = CreateJob("Sales")
	list[3] = CreateJob("HR")

	return list
}

func (j Job) GetTitle() string {
	return j.title
}
