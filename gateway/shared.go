package gateway

import (
	"context"
	"fmt"
	"userprofile/application/apperror"
	"userprofile/domain/service"
	"userprofile/infrastructure/log"
	"userprofile/infrastructure/token"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type commonImplementation struct {
	UserToken *token.JWTToken
}

func (r *commonImplementation) GenerateUUID(ctx context.Context) string {
	log.Info(ctx, "called")

	return uuid.NewString()
}

func (r *commonImplementation) GenerateToken(ctx context.Context, req service.GenerateTokenServiceRequest) (string, error) {
	log.Info(ctx, "called")

	result, err := r.UserToken.CreateToken(req.RawContent)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (r *commonImplementation) SendEmail(ctx context.Context, req service.SendEmailServiceRequest) error {
	log.Info(ctx, "called")

	// Email is sending
	fmt.Printf("token:\n%s\n", req.ContentBody)

	return nil
}

func (r *commonImplementation) ConstructStartActivationMessage(ctx context.Context, req service.ConstructStartActivationMessageServiceRequest) (*service.MessageServiceResponse, error) {
	log.Info(ctx, "called")

	result := service.MessageServiceResponse{
		Subject: fmt.Sprintf("Activation for %s", req.Email),
		Body:    fmt.Sprintf("Click this link for activation\n\nhttp://localhost:8888/activation/%s\n\n", req.ActivationToken),
	}

	return &result, nil
}

func (r *commonImplementation) HashPassword(ctx context.Context, plainPassword string) (string, error) {
	log.Info(ctx, "called")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (r *commonImplementation) ConstructSuccessActivationMessage(ctx context.Context, req service.ConstructSuccessActivationMessageServiceRequest) (*service.MessageServiceResponse, error) {
	log.Info(ctx, "called")

	result := service.MessageServiceResponse{
		Subject: fmt.Sprintf("Success Activation for %s", req.Email),
		Body:    "Congratulation!!",
	}

	return &result, nil
}

func (r *commonImplementation) ValidatePassword(ctx context.Context, req service.ValidatePasswordServiceRequest) error {
	log.Info(ctx, "called")

	err := bcrypt.CompareHashAndPassword([]byte(req.HashedPassword), []byte(req.PlainPassword))
	if err != nil {
		return apperror.InvalidEmailOrPassword
	}

	return nil
}

type contextDBType string

const ContextDBValue contextDBType = "DB"

// extractDB is used by other repo to extract the database from context
func extractDB(ctx context.Context) (*gorm.DB, error) {

	db, ok := ctx.Value(ContextDBValue).(*gorm.DB)
	if !ok {
		return nil, apperror.DatabaseNotFoundInContextError
	}

	return db, nil
}
