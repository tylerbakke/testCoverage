package main

import (
	"github.com/tylerbakke/testCoverage/business"
)

func main() {
	myCompany := business.CreateCompany()
	myCompany.FillPositions()
}

