package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Contact Represents the data from a Contact
type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Notes string `json:"notes"`
}

// to seed the /contacts API
var contacts = []Contact{
	{
		ID:    "1",
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Phone: "123-456-7890",
		Notes: "This is John Doe's contact information.",
	},
	{
		ID:    "2",
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
		Phone: "456-789-0123",
		Notes: "This is Jane Doe's contact information.",
	},
	{
		ID:    "3",
		Name:  "Cam Doe",
		Email: "Cam.doe@example.com",
		Phone: "123-456-7890",
		Notes: "This is Cam Doe's contact information.",
	},
}

// ListContacts will return the current list of contacts
func ListContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}
