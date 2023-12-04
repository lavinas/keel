package domain


// SMTPServer is the struct that contains the SMTP server information
type SMTPServer struct {
	ID		string `json:"id"`
	Host 	string `json:"host"`
	Port 	int `json:"port"`
	User 	string `json:"user"`
	Pass 	string `json:"pass"`
}

//  Sender is the struct that contains the business information
type Sender struct {
	ID		string `json:"id"`
	Name	string `json:"name"`
	Email	string `json:"email"`
}

// Receiver is the struct that contains the client information
type Receiver struct {
	ID 		string `json:"id"`
	Name	string `json:"name"`
	Email	string `json:"email"`
}

// Template is the struct that contains the email template information
type Template struct {
	ID 		string `json:"id"`
	Name	string `json:"name"`
	Subject	string `json:"subject"`
	Body	string `json:"body"`
}

// Email is the struct that contains the email information
type Email struct {
	ID 			string `json:"id"`
	Sender 		Sender `json:"sender"`
	Receiver 	Receiver `json:"receiver"`
	Template 	Template `json:"template"`
	SMTPServer 	SMTPServer `json:"smtp_server"`
	Variables 	map[string]string `json:"variables"`
}

