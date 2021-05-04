package graph

import (
	"github.com/pkg/errors"
	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"ms.api/libs/db"
	"ms.api/types"
)

type DataResolver struct {
	store  db.DataStore
	logger *zap.Logger
}

func NewDataResolver(store db.DataStore, logger *zap.Logger) *DataResolver {
	return &DataResolver{store: store, logger: logger}
}

func (r *DataResolver) ResolveValidation(v models.Validation) (*types.Validation, error) {
	org, err := r.ResolveOrganization(v.Organisation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve organisation")
	}
	p, err := r.ResolvePerson(v.Applicant.ID, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve person")
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
			return nil, errors.Wrap(err, "failed to get check")
		}
		checkOrg, err := r.ResolveOrganization(check.Organisation)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolve check's organisation")
		}
		reports := make([]*types.Report, 0)
		for _, r := range check.Data.Reports {
			reports = append(reports, &types.Report{
				ID:           r.ID,
				Data:         string(r.Data),
				Status:       types.State(r.Status),
				Organisation: checkOrg,
				Ts:           Int64(r.Timestamp.UnixNano()),
				Review: &types.ReportReviewStatus{
					Resubmit: &r.Review.Resubmit,
					Message:  &r.Review.Message,
				},
			})
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
				Reports:               reports,
			},
		}
	}
	if validation.ValidationType == types.ValidationTypeScreen {
		screen, err := r.store.GetScreen(v.Data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get validation screen")
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
		// TODO: Please fix this
	}
	return validation, nil
}

func (r *DataResolver) ResolveOrganization(id string) (*types.Organisation, error) {
	org, err := r.store.GetOrganization(id)
	if err != nil {
		return nil, err
	}

	addresses := make([]*types.Address, 0)
	industries := make([]*types.Industry, 0)
	imageAssets := make([]*types.ImageAssets, 0)

	for _, addr := range org.Addresses {
		coordinateLatitude, coordinateLongitutde := 0.0, 0.0
		if len(addr.Coordinate) > 1 {
			coordinateLatitude, coordinateLongitutde = addr.Coordinate[0], addr.Coordinate[1]
		}
		addresses = append(addresses, &types.Address{
			Street:   &addr.Street,
			City:     &addr.City,
			Postcode: &addr.PostCode,
			Country:  &types.Country{CountryName: addr.Country},
			Location: &types.Location{
				Longitude: &coordinateLongitutde,
				Latitude:  &coordinateLatitude,
			},
		})
	}
	for _, indr := range org.Industry {
		industries = append(industries, &types.Industry{
			Code:        int64(indr.Code),
			Score:       &indr.Score,
			Section:     &indr.Section,
			Description: &indr.Description,
		})
	}
	for _, img := range org.ImageAssets {
		imageAssets = append(imageAssets, &types.ImageAssets{
			Safe:  &img.Safe,
			Type:  &img.Type,
			Image: &img.Type,
			Svg:   &img.Svg,
		})
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
		Addresses:   addresses,
		Location: &types.OrgLocation{
			Continent:   &org.OrgLocation.Continent,
			Country:     &org.OrgLocation.Country,
			State:       &org.OrgLocation.State,
			City:        &org.OrgLocation.City,
			CountryCode: &org.OrgLocation.CountryCode,
		},
		Industries: industries,
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
		ImageAssets: imageAssets,
		Identities:  nil,
	}, nil
}

func (r *DataResolver) ResolvePerson(id string, p *models.Person) (*types.Person, error) {
	if p != nil {
		org, err := r.ResolveOrganization(p.Employer)
		if err != nil {
			return nil, errors.Wrap(err, "failed to resolver person's organisation")
		}
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
		return nil, errors.Wrap(err, "failed to resolver person")
	}
	org, err := r.ResolveOrganization(person.Employer)
	if err != nil {
		r.logger.With(zap.Error(err)).Error("failed to resolver person's organisation")
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
