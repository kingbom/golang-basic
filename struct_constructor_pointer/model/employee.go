package model

import (
	"fmt"
)

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

var employeeInstance *employee

func Init(firstName string, lastName string, totalLeave int, leavesTaken int) *employee {
	if employeeInstance == nil {
		employeeInstance = &employee{firstName, lastName, totalLeave, leavesTaken}
	}
	return employeeInstance
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
