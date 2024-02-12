package controller

import (
	"main/ent"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

type User struct {
	name     string `form:"name"`
	email    string `form:"email"`
	password string `form:"password"`
}

func CreateUserHandler(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u, err := client.User.
			Create().
			SetName(user.name).
			SetEmail(user.email).
			SetPassword(user.password).
			AddGroups().
			Save(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, u)

	}
}

type Invite struct {
	email string `form:"email"`
}

func InviteUser(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invite Invite
		if err := ctx.ShouldBindJSON(&invite); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Extract inviterEmail and inviteeEmail from the request, assuming they are passed as query parameters or part of the JSON body.
		// You may need to adjust this part based on how you intend to receive these emails.
		inviterEmail := "" // Placeholder: Extract from ctx
		inviteeEmail := invite.email

		// Validate that the inviter and invitee have the same email domain.
		if !validateEmailDomain(inviterEmail, inviteeEmail) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "inviter and invitee email domains do not match"})
			return
		}

		// Look up the inviter's user entity to retrieve their group information
		inviter, err := client.User.
			Query().
			Where(user.EmailEQ(inviterEmail)).
			Only(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to retrieve inviter details: %v", err)})
			return
		}

		// Assuming inviter has been successfully retrieved and you want to proceed with sending an email.
		if err := sendInvitationEmail(inviteeEmail); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send invitation email: %v", err)})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Invitation sent successfully"})
	}
}


func sendInvitationEmail(email string) error {
	// Set up authentication information.
	smtpHost := "smtp.example.com" // SMTP server host
	smtpPort := "587"              // SMTP server port
	senderEmail := "your-email@example.com"
	senderPassword := "your-email-password"
	loginUrl := "your-email@example.com"

	// Define the authentication for the SMTP server
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	// Compose the message
	message := []byte("To: " + email + "\r\n" +
		"Subject: Please join here:" + loginUrl + " \r\n")

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{email}, message)
	if err != nil {
		return err
	}

	return nil
}
