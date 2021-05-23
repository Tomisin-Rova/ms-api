package graph

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
	"ms.api/protos/pb/types"
	apitypes "ms.api/types"
)

type DataConverter struct {
	logger *zap.Logger
}

func NewDataConverter(logger *zap.Logger) *DataConverter {
	return &DataConverter{logger: logger}
}

func decodeValidationData(data *anypb.Any) (map[string]interface{}, error) {
	var output map[string]interface{}
	err := json.Unmarshal(data.GetValue(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *DataConverter) ProtoValidationToModel(validation *types.Validation) (*models.Validation, error) {
	decodedData, err := decodeValidationData(validation.Data)
	if err != nil {
		return nil, err
	}

	dataId, ok := decodedData["id"].(string)
	if !ok {
		return nil, errors.New("could not decode id from data")
	}
	protoValidation := models.Validation{
		ID:             validation.Id,
		ValidationType: models.ValidationType(validation.ValidationType),
		Applicant: models.Person{
			ID: validation.Applicant,
		},
		Data:         dataId,
		Organisation: validation.Organisation.Id,
		Status:       models.State(validation.Status),
		Approved:     validation.Approved,
		Timestamp:    time.Unix(int64(validation.Ts), 0),
	}

	return &protoValidation, nil
}

func (c *DataConverter) StateToStringSlice(states []apitypes.State) []string {
	if len(states) == 0 {
		return nil
	}

	strSlice := make([]string, len(states))
	for i, state := range states {
		strSlice[i] = string(state)
	}
	return strSlice
}

func (c *DataConverter) OrganizationFromProto(org *types.Organisation) *apitypes.Organisation {
	addresses := make([]*apitypes.Address, 0)
	industries := make([]*apitypes.Industry, 0)
	imageAssets := make([]*apitypes.ImageAssets, 0)

	for _, addr := range org.Addresses {
		coordinateLatitude, coordinateLongitutde := 0.0, 0.0
		if len(addr.Coordinate) == 2 {
			coordinateLatitude, coordinateLongitutde = addr.Coordinate[0], addr.Coordinate[1]
		}
		addresses = append(addresses, &apitypes.Address{
			Street:   &addr.Street,
			City:     &addr.City,
			Postcode: &addr.Postcode,
			Country:  &apitypes.Country{CountryName: addr.Country},
			Location: &apitypes.Location{
				Longitude: &coordinateLongitutde,
				Latitude:  &coordinateLatitude,
			},
		})
	}
	for _, indr := range org.Industries {
		score := float64(indr.Score)
		industries = append(industries, &apitypes.Industry{
			Code:        int64(indr.Code),
			Score:       &score,
			Section:     &indr.Section,
			Description: &indr.Description,
		})
	}
	for _, img := range org.ImageAssets {
		imageAssets = append(imageAssets, &apitypes.ImageAssets{
			Safe:  &img.Safe,
			Type:  &img.Type,
			Image: &img.Type,
			Svg:   &img.Svg,
		})
	}

	revenue := float64(org.Revenue)
	raised := float64(org.Raised)
	organization := &apitypes.Organisation{
		ID:          org.Id,
		Name:        &org.Name,
		Keywords:    &org.Keywords,
		Description: &org.Description,
		Domain:      &org.Domain,
		Banner:      &org.Banner,
		Revenue:     &revenue,
		Language:    &org.Language,
		Raised:      &raised,
		Employees:   &org.Employees,
		Email:       &org.Email,
		Addresses:   addresses,
		Industries:  industries,
		ImageAssets: imageAssets,
		Identities:  nil,
	}

	if org.Location != nil {
		organization.Location = &apitypes.OrgLocation{
			Continent:   &org.Location.Continent,
			Country:     &org.Location.Country,
			State:       &org.Location.State,
			City:        &org.Location.City,
			CountryCode: &org.Location.CountryCode,
		}
	}

	if org.Social != nil {
		organization.Social = &apitypes.Social{
			Youtube:    &org.Social.Youtube,
			Github:     &org.Social.Github,
			Facebook:   &org.Social.Facebook,
			Pinterest:  &org.Social.Pinterest,
			Instagram:  &org.Social.Instagram,
			Linkedin:   &org.Social.Linkedin,
			Medium:     &org.Social.Medium,
			Crunchbase: &org.Social.Crunchbase,
			Twitter:    &org.Social.Twitter,
		}
	}

	return organization
}
