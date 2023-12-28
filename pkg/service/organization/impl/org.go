package organization_impl

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
	"github.com/quarkloop/quarkloop/pkg/service/organization/store"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
)

type orgService struct {
	store        store.OrgStore
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewOrganizationService(ds store.OrgStore, aclService accesscontrol.Service, quotaService quota.Service) org.Service {
	return &orgService{
		store:        ds,
		aclService:   aclService,
		quotaService: quotaService,
	}
}

func (s *orgService) GetOrganizationList(ctx *gin.Context) ([]org.Organization, error) {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public orgs
		return s.getOrganizationList(ctx, model.PublicVisibility)
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, &accesscontrol.EvaluateFilterParams{
		UserId: user.GetId(),
	})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public orgs
			return s.getOrganizationList(ctx, model.PublicVisibility)
		}
		return nil, err
	}

	// authorized user => return public + private orgs
	return s.getOrganizationList(ctx, model.AllVisibility)
}

func (s *orgService) getOrganizationList(ctx *gin.Context, visibility model.ScopeVisibility) ([]org.Organization, error) {
	orgList, err := s.store.ListOrganizations(ctx, visibility)
	if err != nil {
		return nil, err
	}

	for i := range orgList {
		org := &orgList[i]
		org.GeneratePath()
	}
	return orgList, nil
}

func (s *orgService) GetOrganizationById(ctx *gin.Context, params *org.GetOrganizationByIdParams) (*org.Organization, error) {
	o, err := s.store.GetOrganizationById(ctx, params.OrgId)
	if err != nil {
		return nil, err
	}

	isPrivate := *o.Visibility == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return nil, org.ErrOrgNotFound
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)

		// check permissions
		err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, &accesscontrol.EvaluateFilterParams{
			UserId: user.GetId(),
			OrgId:  params.OrgId,
		})
		if err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return org not found error
				return nil, org.ErrOrgNotFound
			}
			return nil, err
		}
	}

	// anonymous and unauthorized user => return public org
	// authorized user => return public or private org
	o.GeneratePath()
	return o, nil
}

// func (s *orgService) GetOrganization(ctx *gin.Context, params *org.GetOrganizationParams) (*org.Organization, error) {
// 	org, err := s.store.GetOrganization(ctx, &params.Organization)
// 	if err != nil {
// 		return nil, err
// 	}

// 	org.GeneratePath()
// 	return org, nil
// }

func (s *orgService) CreateOrganization(ctx *gin.Context, params *org.CreateOrganizationParams) (*org.Organization, error) {
	if contextdata.IsUserAnonymous(ctx) {
		return nil, errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgCreate, &accesscontrol.EvaluateFilterParams{
		OrgId:  scope.OrgId(),
		UserId: user.GetId(),
	})
	if err != nil {
		return nil, err
	}

	// check quotas
	if err := s.quotaService.CheckCreateOrgQuotaReached(ctx, user.GetId()); err != nil {
		return nil, err
	}

	o, err := s.store.CreateOrganization(ctx, &params.Organization)
	if err != nil {
		return nil, err
	}
	o.GeneratePath()

	return o, nil
}

func (s *orgService) UpdateOrganizationById(ctx *gin.Context, params *org.UpdateOrganizationByIdParams) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgUpdate, &accesscontrol.EvaluateFilterParams{
		OrgId:  params.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return err
	}

	return s.store.UpdateOrganizationById(ctx, params.OrgId, &params.Organization)
}

func (s *orgService) DeleteOrganizationById(ctx *gin.Context, params *org.DeleteOrganizationByIdParams) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgDelete, &accesscontrol.EvaluateFilterParams{
		OrgId:  params.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return err
	}

	return s.store.DeleteOrganizationById(ctx, params.OrgId)
}
