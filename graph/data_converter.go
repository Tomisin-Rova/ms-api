package graph

import (
	"encoding/json"
	"time"

	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
	"ms.api/protos/pb/types"
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

	dataId := decodedData["id"].(string)
	protoValidation := models.Validation{
		ID:             validation.Id,
		ValidationType: models.ValidationType(validation.ValidationType),
		Applicant: models.Person{
			ID: validation.Applicant,
		},
		Data:         dataId,
		Organisation: validation.Organisation,
		Status:       models.State(validation.Status),
		Approved:     validation.Approved,
		Timestamp:    time.Unix(int64(validation.Ts), 0),
	}

	return &protoValidation, nil
}
