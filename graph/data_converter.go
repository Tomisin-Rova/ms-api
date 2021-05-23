package graph

import (
	"bytes"
	"encoding/base64"
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

func (c *DataConverter) HydrateValidationData(validation *apitypes.Validation, data *anypb.Any, owner apitypes.Owner) error {
	decodedData, err := decodeValidationData(data)
	if err != nil {
		return err
	}
	// TODO - There are some repeated patterns below we should update in the future
	switch validation.ValidationType {
	case apitypes.ValidationTypeCheck:
		check := &apitypes.Check{
			ID:           validation.ID,
			Owner:        owner,
			Organisation: validation.Organisation,
			Status:       apitypes.State(validation.Status),
			Ts:           validation.Ts,
		}
		var checkData apitypes.CheckData
		var dataReports []apitypes.Report
		if decodedData["data"] != nil {
			if dataMap, ok := decodedData["data"].(map[string]interface{}); ok {
				if dataMap["reports"] != nil {
					if reports, ok := dataMap["reports"].([]interface{}); ok {
						for _, reportInterface := range reports {
							reportMap, ok := reportInterface.(map[string]interface{})
							if !ok {
								continue
							}
							var report apitypes.Report
							var review apitypes.ReportReviewStatus
							if reportMap["review"] != nil {
								err = c.transcodeData(reportMap["review"], &review)
								if err != nil {
									c.logger.Debug("failed to trancode report review from validation report", zap.Error(err))
								}
								reportMap["review"] = nil
							}
							var reportData string
							if reportMap["data"] != nil {
								if reportMapData, ok := reportMap["data"].(map[string]interface{}); ok {
									if reportMap64Encoded, ok := reportMapData["Data"].(string); ok {
										b, err := base64.StdEncoding.DecodeString(reportMap64Encoded)
										if err != nil {
											c.logger.Debug("failed to trancode report data from validation", zap.Error(err))
											continue
										}
										reportData = string(b)
									}
								}
								reportMap["data"] = nil
							}
							var reportOrganization apitypes.Organisation
							if orgId, ok := reportMap["organisation"].(string); ok {
								reportOrganization = apitypes.Organisation{
									ID: orgId,
								}
								reportMap["organisation"] = nil
							}
							var reportTs int64
							if tsStr, ok := reportMap["ts"].(string); ok {
								ts, err := time.Parse(time.RFC3339, tsStr)
								if err != nil {
									c.logger.Debug("failed to parse timestamp data from validation report", zap.Error(err))
								}
								reportTs = ts.UnixNano()
								reportMap["ts"] = nil
							}
							err = c.transcodeData(reportMap, &report)
							if err != nil {
								c.logger.Debug("failed to trancode reports from validation", zap.Error(err))
							}
							report.Review = &review
							report.Data = reportData
							report.Organisation = &reportOrganization
							report.Ts = &reportTs
							dataReports = append(dataReports, report)
						}
					}
					dataMap["reports"] = nil
				}
				if dataMap["tags"] != nil {
					dataMap["tags"] = nil
				}
			}
		}
		err := c.transcodeData(decodedData["data"], &checkData)
		if err != nil {
			return err
		}
		checkData.Reports = make([]*apitypes.Report, len(dataReports))
		for i := range dataReports {
			report := dataReports[i]
			checkData.Reports[i] = &report
		}
		check.Data = &checkData
		validation.Data = check
	case apitypes.ValidationTypeScreen:
		var screen apitypes.Screen
		var screenData string
		if decodedData["data"] != nil {
			jsonBytes, err := json.Marshal(decodedData["data"])
			if err != nil {
				c.logger.Debug("failed to decode screen data", zap.Error(err))
			}
			screenData = string(jsonBytes)
			decodedData["data"] = nil
		}
		var screenOrganization apitypes.Organisation
		if orgId, ok := decodedData["organisation"].(string); ok {
			screenOrganization = apitypes.Organisation{
				ID: orgId,
			}
			decodedData["organisation"] = nil
		}
		if decodedData["organisation"] != nil {
			decodedData["organisation"] = nil
		}
		var screenTs int64
		if tsStr, ok := decodedData["ts"].(string); ok {
			ts, err := time.Parse(time.RFC3339, tsStr)
			if err != nil {
				c.logger.Debug("failed to parse timestamp data from validation report", zap.Error(err))
			}
			screenTs = ts.UnixNano()
			decodedData["ts"] = nil
		}
		err := c.transcodeData(decodedData, &screen)
		if err != nil {
			c.logger.Debug("failed to trancode screen from validation", zap.Error(err))
			return err
		}
		validation.Data = &apitypes.Screen{
			ID:           screen.ID,
			Data:         screenData,
			Organisation: &screenOrganization,
			Status:       apitypes.State(screen.Status),
			Ts:           &screenTs,
		}
	case apitypes.ValidationTypeProof:
		var proof models.Proof
		err := c.transcodeData(decodedData, &proof)
		if err != nil {
			return err
		}
		validation.Data = &apitypes.Proof{
			ID:     proof.ID,
			Type:   apitypes.ProofType(proof.Data),
			Data:   string(proof.Data),
			Status: apitypes.State(proof.Status),
			Ts:     Int64(proof.Timestamp.UnixNano()),
		}
	}

	return nil
}

func (c *DataConverter) transcodeData(in, out interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		return err
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		return err
	}
	return nil
}
