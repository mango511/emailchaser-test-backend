// controllers/utils.go

package controllers

import "strings"

// validateEmailDomain checks if the domain part of two email addresses is the same.
func validateEmailDomain(inviterEmail, inviteeEmail string) bool {
    inviterDomain := strings.Split(inviterEmail, "@")
    inviteeDomain := strings.Split(inviteeEmail, "@")
    // Check if both emails have a domain part and if they are equal
    return len(inviterDomain) == 2 && len(inviteeDomain) == 2 && inviterDomain[1] == inviteeDomain[1]
}
