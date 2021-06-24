package registry

import (
	"context"
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
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type usingdb struct {
	server.GinHTTPHandler
	userapiController userapi.Controller
	// TODO Another controller will added here ... <<<<<<
}

func NewUsingdb() func() application.RegistryContract {
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

		databaseConnectionString := viper.GetString("database")

		db, err := gorm.Open(sqlite.Open(databaseConnectionString), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		httpHandler := server.NewGinHTTPHandler(":8080")
		datasource := gateway.NewIndatabase2Gateway(userToken, db)

		return &usingdb{
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
				// TODO another Inport will added here ... <<<<<<
			},
			// TODO another controller will added here ... <<<<<<
		}

	}
}

func (r *usingdb) SetupController() {
	r.userapiController.RegisterRouter()
	// TODO another router call will added here ... <<<<<<
}
