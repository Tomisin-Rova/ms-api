package fakes

import (
	"context"
	"google.golang.org/grpc"
	"ms.api/protos/pb/onfidoService"
)

type OnfidoClientFakes struct {
	void *onfidoService.Void
	resp *onfidoService.ApplicantSDKTokenResponse
	err error
}

func NewFakeOnFidoClient(void *onfidoService.Void, resp *onfidoService.ApplicantSDKTokenResponse, err error) *OnfidoClientFakes {
	return &OnfidoClientFakes{
		void: void,
		resp: resp,
		err:  err,
	}
}

func (f *OnfidoClientFakes) GenerateApplicantSDKToken(ctx context.Context, req *onfidoService.ApplicantSDKTokenRequest,
	opts... grpc.CallOption) (*onfidoService.ApplicantSDKTokenResponse, error) {
	return f.resp, f.err
}

func (f *OnfidoClientFakes) WebhookPush(ctx context.Context, req *onfidoService.OnfidoCheckWebhookRequest,
	opts... grpc.CallOption) (*onfidoService.Void, error) {
	return f.void, f.err
}
