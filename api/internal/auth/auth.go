package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// Get current user from context
func GetUser(ctx *gin.Context) string {
	authString := ctx.GetHeader("Authorization")
	str2Slice := strings.Split(authString, " ")
	token := str2Slice[1]
	parts := strings.Split(token, ".")
	body := parts[1]

	decoded, err := base64.RawURLEncoding.DecodeString(body)
	if err != nil {
		log.Println("JWT decoding failed:", err)
		return ""
	}

	var m map[string]interface{}
	if err := json.Unmarshal(decoded, &m); err != nil {
		log.Println("json decoding failed:", err)
		return ""
	}

	var subStr string
	if sub, ok := m["sub"]; ok {
		subStr = fmt.Sprintf("%v", sub)
		return subStr
	}

	return ""
}
