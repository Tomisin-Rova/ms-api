package mapper

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	pb "ms.api/protos/pb/types"
	"ms.api/types"
)

func TestGQLMapper_HydrateProduct(t *testing.T) {
	mapper := &GQLMapper{}

	from := &pb.Product{
		Id:             "id",
		Identification: "identification",
		Details: &pb.ProductDetails{
			Name: "Roava Classic GBP Current Account",
			ProductControl: &pb.ProductControl{
				DormancyPeriodDays: 12,
				OpeningBalance: &pb.OpeningBalance{
					Max: 20,
				},
			},
			OverdraftSetting: &pb.OverdraftSetting{
				AllowTechnicalOverdraft: true,
				InterestSettings: &pb.InterestSettings{
					DaysInYear: "123",
					RateTiers: []*pb.RateTiers{{
						EndingBalance: 12,
					}},
				},
			},
		},
	}

	var product types.Product
	err := mapper.Hydrate(from, &product)
	assert.Nil(t, err)

	assert.NotNil(t, product)
	assert.Equal(t, from.Id, product.ID)
	assert.Equal(t, from.Identification, *product.Identification)
	assert.Equal(t, from.Details.Name, *product.Details.Name)
	assert.Equal(t, from.Details.ProductControl.DormancyPeriodDays, *product.Details.ProductControl.DormancyPeriodDays)
	assert.Equal(t, from.Details.ProductControl.OpeningBalance.Max, *product.Details.ProductControl.OpeningBalance.Max)
	assert.Equal(t, from.Details.OverdraftSetting.InterestSettings.DaysInYear, *product.Details.OverdraftSetting.InterestSettings.DaysInYear)
	assert.Equal(t, from.Details.OverdraftSetting.AllowTechnicalOverdraft, *product.Details.OverdraftSetting.AllowTechnicalOverdraft)
	assert.Equal(t, from.Details.OverdraftSetting.InterestSettings.RateTiers[0].EndingBalance, *product.Details.OverdraftSetting.InterestSettings.RateTiers[0].EndingBalance)
}

func TestGQLMapper_HydrateAccount(t *testing.T) {
	mapper := &GQLMapper{}

	from := &pb.Account{
		Id:    "id",
		Owner: "owner",
		AccountData: &pb.AccountData{
			Name: "name",
		},
	}

	var account types.Account
	err := mapper.Hydrate(from, &account)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, from.Id, account.ID)
}

func TestGQLMapper_HydrateTransaction(t *testing.T) {
	mapper := &GQLMapper{}

	from := &pb.Transaction{
		Id: "id",
		Account: &pb.Account{
			Id:    "accountId",
			Owner: "owner",
		},
		TransactionData: &pb.TransactionData{
			Id:     "id",
			Amount: 123,
			Fees: []*pb.TransactionFee{{
				Name: "name",
			}},
			AffectedAmounts: &pb.AffectedAmounts{
				FeesAmount: 43,
			},
		},
	}

	var transaction types.Transaction
	err := mapper.Hydrate(from, &transaction)
	assert.Nil(t, err)

	assert.Nil(t, err)
	assert.Equal(t, from.Id, transaction.ID)
	assert.Equal(t, from.Account.Id, transaction.Account.ID)
	assert.Equal(t, from.TransactionData.Fees[0].Name, *transaction.TransactionData.Fees[0].Name)
}

func TestGQLMapper_HydrateTag(t *testing.T) {
	mapper := &GQLMapper{}

	from := &pb.Tag{
		Id:   "id",
		Name: "name",
	}

	var tag types.Tag
	err := mapper.Hydrate(from, &tag)
	assert.Nil(t, err)

	assert.NotNil(t, tag)
	assert.Equal(t, from.Id, tag.ID)
	assert.Equal(t, from.Name, *tag.Name)
}

func TestGQLMapper_HydratePayment(t *testing.T) {
	const (
		successPayeeAccountAsTarget = iota
		successAccountAsTarget
		failNilTargetAccount
		failNilTargetPayeeAccount
	)

	ts := time.Now()
	dob := "1994-01-01"
	paymentOwner := &pb.Person{
		Id:          generateID(),
		FirstName:   "First",
		LastName:    "Last",
		Dob:         dob,
		Ts:          ts.Unix(),
		Nationality: []string{"UK", "NG"},
		Emails: []*pb.Email{
			{
				Value:    "firstemail@email.com",
				Verified: true,
			},
			{
				Value:    "secondemail@email.com",
				Verified: true,
			},
		},
		Phones: []*pb.PhoneNumber{
			{
				Number:   "+447911123456",
				Verified: true,
			},
			{
				Number:   "+23410701234",
				Verified: true,
			},
		},
		Identities: []*pb.ExtendedIdentity{
			{
				Id:    generateID(),
				Owner: &pb.Person{Id: generateID()},
				Ts:    ts.Add(time.Second).Unix(),
			},
		},
		Cdd: &pb.Cdd{
			Status:  "ONBOARDED",
			Onboard: true,
			Ts:      int32(ts.Unix()),
		},
	}

	var tests = []struct {
		testType int
		name     string
		args     *pb.Payment
	}{
		{
			testType: successPayeeAccountAsTarget,
			name:     "Succesful hydrate payment with target payee account",
			args: &pb.Payment{
				Id:             generateID(),
				IdempotencyKey: generateID(),
				Owner:          paymentOwner,
				Charge:         0.0,
				Reference:      "test reference",
				Status:         "APPROVED",
				Source: &pb.PaymentAccount{
					Accounts: &pb.PaymentAccount_Account{
						Account: &pb.Account{
							Id:             generateID(),
							AccountData:    new(pb.AccountData),
							AccountDetails: new(pb.AccountDetails),
						},
					},
					Currency: "GBP",
					Amount:   1000.0,
				},
				Target: &pb.PaymentAccount{
					Accounts: &pb.PaymentAccount_PayeeAccount{
						PayeeAccount: &pb.PayeeAccount{
							Id: generateID(),
						},
					},
					Currency: "GBP",
					Amount:   1000.0,
				},
			},
		},
		{
			testType: successAccountAsTarget,
			name:     "Succesful hydrate payment with target account",
			args: &pb.Payment{
				Id:             generateID(),
				IdempotencyKey: generateID(),
				Owner:          paymentOwner,
				Charge:         0.0,
				Reference:      "test reference",
				Status:         "APPROVED",
				Source: &pb.PaymentAccount{
					Accounts: &pb.PaymentAccount_Account{
						Account: &pb.Account{
							Id:             generateID(),
							AccountData:    new(pb.AccountData),
							AccountDetails: new(pb.AccountDetails),
						},
					},
					Currency: "GBP",
					Amount:   1000.0,
				},
				Target: &pb.PaymentAccount{
					Accounts: &pb.PaymentAccount_Account{
						Account: &pb.Account{
							Id:             generateID(),
							AccountData:    new(pb.AccountData),
							AccountDetails: new(pb.AccountDetails),
						},
					},
					Currency: "GBP",
					Amount:   1000.0,
				},
			},
		},
		{
			testType: failNilTargetAccount,
			name:     "Fail if target account is nil",
			args: &pb.Payment{
				Id:             generateID(),
				IdempotencyKey: generateID(),
				Owner:          paymentOwner,
				Charge:         0.0,
				Reference:      "test reference",
				Status:         "APPROVED",
				Source: &pb.PaymentAccount{
					Accounts: &pb.PaymentAccount_Account{
						Account: &pb.Account{
							Id:             generateID(),
							AccountData:    new(pb.AccountData),
							AccountDetails: new(pb.AccountDetails),
						},
					},
					Currency: "GBP",
					Amount:   1000.0,
				},
				Target: &pb.PaymentAccount{
					Accounts: &pb.PaymentAccount_Account{
						Account: nil,
					},
					Currency: "GBP",
					Amount:   1000.0,
				},
			},
		},
		{
			testType: failNilTargetPayeeAccount,
			name:     "Fail if target account is nil",
			args: &pb.Payment{
				Id:             generateID(),
				IdempotencyKey: generateID(),
				Owner:          paymentOwner,
				Charge:         0.0,
				Reference:      "test reference",
				Status:         "APPROVED",
				Source: &pb.PaymentAccount{
					Accounts: &pb.PaymentAccount_Account{
						Account: &pb.Account{
							Id:             generateID(),
							AccountData:    new(pb.AccountData),
							AccountDetails: new(pb.AccountDetails),
						},
					},
					Currency: "GBP",
					Amount:   1000.0,
				},
				Target: &pb.PaymentAccount{
					Accounts: &pb.PaymentAccount_PayeeAccount{
						PayeeAccount: nil,
					},
					Currency: "GBP",
					Amount:   1000.0,
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case successPayeeAccountAsTarget:
				mapper := &GQLMapper{}
				var payment types.Payment
				err := mapper.Hydrate(testCase.args, &payment)
				assert.Nil(t, err)
			case successAccountAsTarget:
				mapper := &GQLMapper{}
				var payment types.Payment
				err := mapper.Hydrate(testCase.args, &payment)
				assert.Nil(t, err)
			case failNilTargetAccount:
				logger := zaptest.NewLogger(t, zaptest.WrapOptions(zap.Hooks(func(e zapcore.Entry) error {
					expectedMessages := "target account decoding"
					if !strings.Contains(expectedMessages, e.Message) {
						t.Fatalf("Log with one of this messages: '%s' should happen", expectedMessages)
					}
					return nil
				})))
				mapper := &GQLMapper{logger: logger}
				var payment types.Payment
				err := mapper.Hydrate(testCase.args, &payment)
				assert.NotNil(t, err)
			case failNilTargetPayeeAccount:
				logger := zaptest.NewLogger(t, zaptest.WrapOptions(zap.Hooks(func(e zapcore.Entry) error {
					expectedMessages := "target payee account decoding"
					if !strings.Contains(expectedMessages, e.Message) {
						t.Fatalf("Log with one of this messages: '%s' should happen", expectedMessages)
					}
					return nil
				})))
				mapper := &GQLMapper{logger: logger}
				var payment types.Payment
				err := mapper.Hydrate(testCase.args, &payment)
				assert.NotNil(t, err)
			}
		})
	}

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateID() string {
	b := make([]byte, 23)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
