package registry

import (
	"context"
	"fmt"
	"os"
	"userprofile/application"
	"userprofile/controller/userapi"
	"userprofile/gateway"
	"userprofile/infrastructure/log"
	"userprofile/infrastructure/server"
	"userprofile/infrastructure/token"
	"userprofile/usecase/activation"
	"userprofile/usecase/loginuser"
	"userprofile/usecase/registeruser"
	"userprofile/usecase/showalluser"
	"userprofile/usecase/showuser"
	"userprofile/usecase/updateuser"

	"github.com/spf13/viper"
)

type simplememory struct {
	server.GinHTTPHandler
	userapiController userapi.Controller
	// TODO Another controller will added here ... <<<<<<
}

func NewSimplememory() func() application.RegistryContract {
	return func() application.RegistryContract {

		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			log.Error(context.Background(), "Config Problem %v", err.Error())
			os.Exit(1)
		}

		secretKey := viper.GetString("secretkey")
		userToken, err := token.NewJWTToken(secretKey)
		if err != nil {
			log.Error(context.Background(), "Secret Key Problem %v", err.Error())
			os.Exit(1)
		}

		port := viper.GetInt("port")
		httpHandler, err := server.NewGinHTTPHandler(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Error(context.Background(), "%v", err.Error())
			os.Exit(1)
		}

		datasource, err := gateway.NewInmemoryGateway(userToken)
		if err != nil {
			log.Error(context.Background(), "%v", err.Error())
			os.Exit(1)
		}

		return &simplememory{
			GinHTTPHandler: httpHandler,
			userapiController: userapi.Controller{
				UserToken:          userToken,
				Router:             httpHandler.Router,
				ActivationInport:   activation.NewUsecase(datasource),
				LoginUserInport:    loginuser.NewUsecase(datasource),
				RegisterUserInport: registeruser.NewUsecase(datasource),
				ShowAllUserInport:  showalluser.NewUsecase(datasource),
				ShowUserInport:     showuser.NewUsecase(datasource),
				UpdateUserInport:   updateuser.NewUsecase(datasource),
			},
			// TODO another controller will added here ... <<<<<<
		}

	}
}

func (r *simplememory) SetupController() {
	r.userapiController.RegisterRouter()
	// TODO another router call will added here ... <<<<<<
}
