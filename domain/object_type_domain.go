package domain

var (
	MessageObjectTypeSuccessRetrieve  = "Object types retrieved successfully"
	MessageObjectTypeSuccessCreate    = "Object type created successfully"
	MessageObjectTypeSuccessGetDetail = "Object type detail retrieved successfully"
	MessageObjectTypeSuccessUpdate    = "Object type updated successfully"
	MessageObjectTypeSuccessDelete    = "Object type deleted successfully"
	MessageObjectTypeFailedRetrieve   = "Object types retrieved failed"
	MessageObjectTypeFailedCreate     = "Object type created failed"
	MessageObjectTypeFailedGet        = "Object type detail retrieved failed"
	MessageObjectTypeFailedUpdate     = "Object type updated failed"
	MessageObjectTypeFailedDelete     = "Object type deleted failed"

	ErrObjectTypeNotFound     = "object type not found"
	ErrObjectTypeAlreadyExist = "object type already exists"
)

type (
	ObjectTypeCreateRequest struct {
		Name string `json:"name" validate:"required,min=4"`
	}

	ObjectTypeUpdateRequest struct {
		Name string `json:"name,omitempty" validate:"min=4"`
	}
)
