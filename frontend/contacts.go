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

// GetContact will return the contact given its id.
func GetContact(c *gin.Context) {
	id := c.Param("id")

	for _, contact := range contacts {
		if contact.ID == id {
			c.IndentedJSON(http.StatusOK, contact)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Contact not found!"})
}

// Create a new Contact.
func CreateContact(c *gin.Context) {
	var contact Contact
	if err := c.BindJSON(&contact); err != nil {
		return
	}

	contacts = append(contacts, contact)
	c.IndentedJSON(http.StatusOK, contact)
}

// Update a new contact.
func UpdateContact(c *gin.Context) {
	var newContact Contact
	if err := c.BindJSON(&newContact); err != nil {
		return
	}

	for i, contact := range contacts {
		if contact.ID == newContact.ID {
			contacts[i] = newContact
			c.IndentedJSON(http.StatusOK, newContact)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Contact not found!"})
}
