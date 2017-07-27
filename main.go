package main

import (
	"github.com/testCoverage/business"
)

func main() {
	myCompany := business.CreateCompany()
	myCompany.FillPositions()
}

