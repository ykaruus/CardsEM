package utils

const (
	// Usuário / Autenticação
	USER_ALREADY_EXISTS = "USER_ALREADY_EXISTS"
	USER_NOT_FOUND      = "USER_NOT_FOUND"
	INVALID_CREDENTIALS = "INVALID_CREDENTIALS"
	INVALID_TOKEN       = "INVALID_TOKEN"
	TOKEN_EXPIRED       = "TOKEN_EXPIRED"
	UNAUTHORIZED_ACCESS = "UNAUTHORIZED_ACCESS"
	AUTHORIZED_ACCESS   = "AUTHORIZED_ACCESS"

	// Validação / Requisição
	INVALID_REQUEST_BODY    = "INVALID_REQUEST_BODY"
	MISSING_REQUIRED_FIELDS = "MISSING_REQUIRED_FIELDS"
	VALIDATION_ERROR        = "VALIDATION_ERROR"

	// Recurso
	RESOURCE_NOT_FOUND      = "RESOURCE_NOT_FOUND"
	RESOURCE_ALREADY_EXISTS = "RESOURCE_ALREADY_EXISTS"
	RESOURCE_CONFLICT       = "RESOURCE_CONFLICT"

	// Permissões / Acesso
	ACCESS_DENIED         = "ACCESS_DENIED"
	OPERATION_NOT_ALLOWED = "OPERATION_NOT_ALLOWED"

	// Sistema / Infraestrutura
	INTERNAL_ERROR      = "INTERNAL_ERROR"
	SERVICE_UNAVAILABLE = "SERVICE_UNAVAILABLE"
	DATABASE_ERROR      = "DATABASE_ERROR"
	TIMEOUT_OCCURRED    = "TIMEOUT_OCCURRED"

	// CARDS - USER
	CARD_ALREADY_EXISTS = "CARD_ALREADY_EXISTS"
	CARD_NOT_FOUND      = "CARD_NOT_FOUND"

	RESOURCE_SUCCESS_CREATED = "RESOURCE_SUCCESS_CREATED"
	SUCCESS                  = "SUCCESS"
)

var StatusCodes = struct {
	Ok, Created, NoContent                                  int
	BadRequest, Unauthorized, Forbidden, NotFound, Conflict int
	UnprocessableEntity, InternalError                      int
}{
	Ok:                  200,
	Created:             201,
	NoContent:           204,
	BadRequest:          400,
	Unauthorized:        401,
	Forbidden:           403,
	NotFound:            404,
	Conflict:            409,
	UnprocessableEntity: 422,
	InternalError:       500,
}

var StatusMessages = struct {
	Success, Created, NoContent                                                                 string
	BadRequest, Unauthorized, Forbidden, NotFound, Conflict                                     string
	UnprocessableEntity, InternalError, Authorized, ResourceUpdated, InvalidId, Token_malformed string
}{
	Success:             "Requisição bem-sucedida",
	Created:             "Recurso criado com sucesso",
	NoContent:           "Nenhum conteúdo",
	BadRequest:          "Corpo da requisição inválido ou malformado",
	Unauthorized:        "Não autorizado",
	Authorized:          "Autorizado com sucesso",
	Forbidden:           "Acesso negado",
	NotFound:            "Recurso não encontrado",
	Conflict:            "Conflito de dados",
	UnprocessableEntity: "Entidade não processável",
	InternalError:       "Erro interno no servidor",
	ResourceUpdated:     "Recurso atualizado com sucesso",
	InvalidId:           "Falha na validação do ID. Verifique o parâmetro e tente novamente",
	Token_malformed:     "O token fornecido está malformado ou invalído",
}
