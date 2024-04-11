package app

import (
	"kang-blogging/iam/app/command"
	"kang-blogging/iam/app/query"
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
