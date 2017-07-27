package business

import (
	"github.com/testCoverage/person"
	"github.com/testCoverage/work"
)

type Company struct{
	companyName string
	employeeList map[string]person.Employee
	jobList []work.Job
}

func CreateCompany() Company{
	var myCompany Company

	myCompany.companyName = "Splunk"

	return myCompany
}

func (c Company) FillPositions(){
	c.jobList = work.CreateJobs()

	for _,job := range c.jobList {
		c.AddEmployee(job)
	}
}

func (c Company) AddEmployee(job work.Job) {
	newEmployee := person.CreateEmployee(job)
	c.employeeList[newEmployee.GetName()] = newEmployee
}

func (c Company) GetName() string {
	return c.companyName
}