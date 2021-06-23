package service

import "context"

type GenerateUUIDService interface {
  GenerateUUID(ctx context.Context) string
}

type SendEmailService interface {
  SendEmail(ctx context.Context, req SendEmailServiceRequest) error
}

type SendEmailServiceRequest struct {
  EmailDestination string
  Subject          string
  ContentBody      string
}

type SendEmailActivationService interface {
  SendEmailActivation(ctx context.Context, req SendEmailActivationServiceRequest) error
}

type SendEmailActivationServiceRequest struct {
  Email           string
  ActivationToken string
}

type SendEmailSuccessActivationService interface {
  SendEmailSuccessActivation(ctx context.Context, email string) error
}

type ConstructSuccessActivationMessageService interface {
  ConstructSuccessActivationMessage(ctx context.Context, req ConstructSuccessActivationMessageServiceRequest) (*MessageServiceResponse, error)
}

type ConstructSuccessActivationMessageServiceRequest struct {
  Email string
}

type MessageServiceResponse struct {
  Subject string
  Body    string
}

type ConstructStartActivationMessageService interface {
  ConstructStartActivationMessage(ctx context.Context, req ConstructStartActivationMessageServiceRequest) (*MessageServiceResponse, error)
}

type ConstructStartActivationMessageServiceRequest struct {
  Email           string
  ActivationToken string
}

type GenerateTokenService interface {
  GenerateToken(ctx context.Context, req GenerateTokenServiceRequest) (string, error)
}

type GenerateTokenServiceRequest struct {
  RawContent string
}

type HashPasswordService interface {
  HashPassword(ctx context.Context, plainPassword string) (string, error)
}

type ValidatePasswordService interface {
  ValidatePassword(ctx context.Context, req ValidatePasswordServiceRequest) error
}

type ValidatePasswordServiceRequest struct {
  PlainPassword  string
  HashedPassword string
}
type GenerateRandomStringService interface {
  GenerateRandomString(ctx context.Context) string
}
