package apperror

const (
  FailUnmarshalResponseBodyError ErrorType = "ER1000 Fail to unmarshal response body"  // used by controller
  ObjectNotFound                 ErrorType = "ER1001 Object %s is not found"           // used by injected repo in interactor
  UnrecognizedEnum               ErrorType = "ER1002 %s is not recognized %s enum"     // used by enum
  DatabaseNotFoundInContextError ErrorType = "ER1003 Database is not found in context" // used by repoimpl
  AddressMustNotEmpty            ErrorType = "ER1000 address must not empty"           //
  IDMustNotEmpty                 ErrorType = "ER1088 id must not empty"                //
  FieldMustNotEmpty              ErrorType = "ER1033 field %s must not empty"          //
  EmailMustNotEmpty              ErrorType = "ER1000 email must not empty"             //
  PasswordMustNotEmpty           ErrorType = "ER1000 password must not empty"          //
  EmailAlreadyUsed               ErrorType = "ER1000 email already used"               //
  EmailIsNotFound                ErrorType = "ER1000 email is not found"               //
  UserAlreadyActiveOrSuspended   ErrorType = "ER1000 user already active or suspended" //
  InvalidActivationToken         ErrorType = "ER1000 invalid activation token"         //
  UserAlreadyActive              ErrorType = "ER1000 user already active"              //
  UserIsNotActive                ErrorType = "ER1000 user is not active"               //
  InvalidEmailOrPassword         ErrorType = "ER1000 invalid email or password"        //
  UserIsNotFound                 ErrorType = "ER1000 user is not found"                //
)
