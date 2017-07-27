package business

import (
	"testing"
)
var testCompany Company

func TestCreateCompany(t *testing.T) {
	testCompany = CreateCompany()
	if testCompany.GetName() != "Splunk"{
		t.Errorf("ERROR: Failed assignment")
	}
}

func TestFillPositions(t *testing.T) {
	testCompany.FillPositions()
	if testCompany.employeeList["Programmer"].GetName() != "Mindless"  {
		t.Errorf("ERROR: Failed assignment")
	}
}