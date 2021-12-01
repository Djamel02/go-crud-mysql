// Package Documentation CRUD APIs.
//
// Documentation of our crud API .
//
//     Schemes: http
//     Host: localhost:3000
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - JWT:
//
//     SecurityDefinitions:
//     JWT:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package docs

import "bytes"

// swagger:parameters getEmployeeById updateEmployee deleteEmployee
type getEmployeeById struct {
	// Employee ID
	//in:path
	//required: true
	Id int64 `json:"id"`
}

// The body to pass to signup
//swagger:parameters signup
type registerParams struct {
    //in: body
    //required: true
    Body struct {
        Username string `json:"username"`
		//swagger:strfmt email
		Email string 	`json:"email"`
		Password string `json:"password"`
    }
}


// The body to pass to signin
//swagger:parameters signin
type loginParams struct {
    //in: body
    //required: true
    Body struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
}


// Body to pass employee form data
// swagger:parameters createEmployee
type employeeParams struct {
	//in: formData
	//swagger:file
	//in: formData
	Picture *bytes.Buffer `json:"picture"`
	//in: formData
	//required: true
	Name string `json:"name"`
	//in: formData
	Phone string `json:"phone"`
	//in: formData
	Job string `json:"job"`
	//in: formData
	Country string `json:"country"`
	//in: formData
	City string `json:"city"`
	//in: formData
	Postalcode int64 `json:"postalcode"`
	
}

// The body to pass to edit employee
//swagger:parameters editEmployee
type editEmployeeParams struct {
    //in: body
    //required: true
    Body struct {
        Name string `json:"name"`
		Phone string 	`json:"phone"`
		Job string `json:"job"`
		Country string `json:"country"`
		City string `json:"city"`
		Postalcode int64 `json:"postalcode"`
    }
}

// The body to pass transaction infos
//swagger:parameters transaction
type transactionParams struct {
	//in: body
	//required: true
	Body struct {
		SenderID int64 `json:"senderId"`
		ReceiverID int64 `json:"receiverId"`
		Amount float64 `json:"amount"`
	}
}

// swagger:response jsonResponse
type jsonResponse struct {
	// in:body
	Body struct {
		Success bool 		`json:"success"`
		Message string      `json:"message"`
		Data interface{}    `json:"data"`
	}
}
