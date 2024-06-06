package management

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/policy"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type UpdatePoliciesParams struct {
	Policies []model.Policy
}

type UpdatePoliciesResult struct {
	//Policies []model.Policy
}

type UpdatePoliciesHandler decorator.UsecaseHandler[UpdatePoliciesParams, UpdatePoliciesResult]

type updatePoliciesHandler struct {
	policyRepo policy.Repository
}

func NewUpdatePoliciesHandler(
	policyRepo policy.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) UpdatePoliciesHandler {
	if policyRepo == nil {
		panic("policyRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[UpdatePoliciesParams, UpdatePoliciesResult](
		updatePoliciesHandler{
			policyRepo: policyRepo,
		},
		logger,
		metrics,
	)
}

func (g updatePoliciesHandler) Handle(ctx context.Context, param UpdatePoliciesParams) (UpdatePoliciesResult, error) {
	err := param.Validate()
	if err != nil {
		return UpdatePoliciesResult{}, err
	}

	_, err = g.policyRepo.UpdatePolicies(ctx, param.Policies)
	if err != nil {
		return UpdatePoliciesResult{}, err
	}

	return UpdatePoliciesResult{}, err
}

func (p *UpdatePoliciesParams) Validate() error {
	if len(p.Policies) == 0 {
		return errors.NewBadRequestError("Missing policies")
	}
	return nil
}
