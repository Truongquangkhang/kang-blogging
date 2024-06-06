package management

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/policy"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/model"
)

type GetPoliciesParams struct {
}

type GetPoliciesResult struct {
	Policies []model.Policy
}

type GetPoliciesHandler decorator.UsecaseHandler[GetPoliciesParams, GetPoliciesResult]

type getPoliciesHandler struct {
	policyRepo policy.Repository
}

func NewGetPoliciesHandler(
	policyRepo policy.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetPoliciesHandler {
	if policyRepo == nil {
		panic("policyRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetPoliciesParams, GetPoliciesResult](
		getPoliciesHandler{
			policyRepo: policyRepo,
		},
		logger,
		metrics,
	)
}

func (g getPoliciesHandler) Handle(ctx context.Context, param GetPoliciesParams) (GetPoliciesResult, error) {
	err := param.Validate()
	if err != nil {
		return GetPoliciesResult{}, err
	}

	policies, err := g.policyRepo.GetPolicies(ctx)
	if err != nil {
		return GetPoliciesResult{}, err
	}

	return GetPoliciesResult{
		Policies: policies,
	}, err
}

func (p *GetPoliciesParams) Validate() error {
	return nil
}
