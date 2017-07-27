package business

import (
	"testing"
	"github.com/tylerbakke/testCoverage/work"
)
var testCompany Company

func TestCreateCompany(t *testing.T) {
	testCompany = CreateCompany()
	if testCompany.GetName() != "Splunk"{
		t.Errorf("ERROR: Failed assignment")
	}
}
func TestCompany_AddEmployee(t *testing.T) {
	testCompany.AddEmployee(work.CreateJob("Tester"))
	if testCompany.employeeList["Mindless Drone"].GetJob() != "Tester" {
		t.Errorf("ERROR: Failed assignment")
	}
}

func TestFillPositions(t *testing.T) {
	testCompany.FillPositions()
	if testCompany.employeeList["Mindless Drone"].GetName() != "Mindless Drone"  {
		t.Errorf("ERROR: Failed assignment")
	}
}
