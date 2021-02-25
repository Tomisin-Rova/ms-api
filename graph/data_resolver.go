package graph

import (
	"github.com/roava/zebra/models"
	"ms.api/libs/db"
	"ms.api/types"
)

type DataResolver struct {
	store db.DataStore
}

func NewDataResolver(store db.DataStore) *DataResolver {
	return &DataResolver{store: store}
}

func (r *DataResolver) ResolveValidation(v models.Validation) (*types.Validation, error) {
	org, err := r.ResolveOrganization(v.Organisation)
	if err != nil {
		return nil, err
	}
	p, err := r.ResolvePerson(v.Applicant.ID, nil)
	if err != nil {
		return nil, err
	}
	validation := &types.Validation{
		ID:             v.ID,
		ValidationType: types.ValidationType(v.ValidationType),
		Applicant:      p,
		Organisation:   org,
		Status:         types.State(v.Status),
		Approved:       &v.Approved,
		Ts:             Int64(v.Timestamp.UnixNano()),
	}
	if validation.ValidationType == types.ValidationTypeCheck {
		check, err := r.store.GetCheck(v.Data)
		if err != nil {
			return nil, err
		}
		validation.Data = types.Check{
			ID:           check.ID,
			Organisation: org,
			Status:       types.State(check.Status),
			Ts:           Int64(check.Timestamp.UnixNano()),
			Data: &types.CheckData{
				ID:                    check.Data.ID,
				CreatedAt:             String(check.Data.CreatedAt.String()),
				Status:                types.State(check.Data.Status),
				RedirectURI:           &check.Data.ResultsURI,
				Result:                &check.Data.Result,
				Sandbox:               &check.Data.Sandbox,
				ResultsURI:            &check.Data.ResultsURI,
				FormURI:               &check.Data.FormURI,
				Paused:                &check.Data.Paused,
				Version:               &check.Data.Version,
				Href:                  &check.Data.HREF,
				ApplicantID:           &check.Data.ApplicantID,
				ApplicantProvidesData: &check.Data.ApplicantProvidesData,
				Tags:                  nil,
			},
		}
	}
	if validation.ValidationType == types.ValidationTypeScreen {
		screen, err := r.store.GetScreen(v.Data)
		if err != nil {
			return nil, err
		}
		validation.Data = types.Screen{
			ID:           screen.ID,
			Data:         string(screen.Data),
			Organisation: org,
			Status:       types.State(screen.Status),
			Ts:           Int64(screen.Timestamp.UnixNano()),
		}
	}
	//nolint
	if validation.ValidationType == types.ValidationTypeProof {

	}
	return validation, nil
}

func (r *DataResolver) ResolveOrganization(id string) (*types.Organisation, error) {
	org, err := r.store.GetOrganization(id)
	if err != nil {
		return nil, err
	}
	return &types.Organisation{
		ID:          org.ID,
		Name:        &org.Name,
		Keywords:    &org.Keywords,
		Description: &org.Description,
		Domain:      &org.Domain,
		Banner:      &org.Banner,
		Revenue:     &org.Revenue,
		Language:    &org.Language,
		Raised:      &org.Raised,
		Employees:   &org.Employees,
		Email:       &org.Email,
		Ts:          Int64(org.Timestamp.UnixNano()),
		Addresses:   nil,
		Location: &types.OrgLocation{
			Continent:   &org.OrgLocation.Continent,
			Country:     &org.OrgLocation.Country,
			State:       &org.OrgLocation.State,
			City:        &org.OrgLocation.City,
			CountryCode: &org.OrgLocation.CountryCode,
		},
		Industries: nil,
		Social: &types.Social{
			Youtube:    &org.Social.Youtube,
			Github:     &org.Social.Github,
			Facebook:   &org.Social.Facebook,
			Pinterest:  &org.Social.Pinterest,
			Instagram:  &org.Social.Instagram,
			Linkedin:   &org.Social.Linkedin,
			Medium:     &org.Social.Medium,
			Crunchbase: &org.Social.Crunchbase,
			Twitter:    &org.Social.Twitter,
		},
		ImageAssets: nil,
		Identities:  nil,
	}, nil
}

func (r *DataResolver) ResolvePerson(id string, p *models.Person) (*types.Person, error) {
	org, err := r.ResolveOrganization(p.Employer)
	if err != nil {
		return nil, err
	}
	if p != nil {
		return &types.Person{
			ID:               p.ID,
			Title:            &p.Title,
			FirstName:        p.FirstName,
			LastName:         p.LastName,
			MiddleName:       &p.MiddleName,
			Dob:              p.Dob.String(),
			Employer:         org,
			Ts:               p.Timestamp.UnixNano(),
			CountryResidence: &p.CountryResidence,
		}, nil
	}
	person, err := r.store.GetPerson(id)
	if err != nil {
		return nil, err
	}

	return &types.Person{
		ID:               person.ID,
		Title:            &person.Title,
		FirstName:        person.FirstName,
		LastName:         person.LastName,
		MiddleName:       &person.MiddleName,
		Dob:              person.Dob.String(),
		Employer:         org,
		Ts:               person.Timestamp.UnixNano(),
		CountryResidence: &person.CountryResidence,
	}, nil
}
