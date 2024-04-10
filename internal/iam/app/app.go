package app

import (
	"kang-edu/iam/app/command"
	"kang-edu/iam/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	// This is a sample command
	DoSomething command.DoSomethingHandler
}

type Queries struct {
	AllApplicableVouchers   query.AllApplicableVouchersHandler
	ApplicableVoucherByCode query.ApplicableVoucherByCodeHandler
}
