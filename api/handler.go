package handlers

import (
    "encoding/json"
    "encoding/base64"
    "net/http"
    "time"

    "github.com/your-org/test-signer/models"
    "github.com/your-org/test-signer/signature"
    "github.com/dgrijalva/jwt-go"
    "os"
)


func SignHandler(w http.ResponseWriter, r *http.Request) {
    
    jwtCookie, err := r.Cookie("token")
    if err != nil {
        if err == http.ErrNoCookie {
            http.Error(w, "Missing cookie: token", http.StatusUnauthorized)
            return
        }

        http.Error(w, "Error accessing cookie: "+err.Error(), http.StatusInternalServerError)
        return
    }


    token, err := jwt.Parse(jwtCookie.Value, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("SECRET_KEY")), nil
    })
    if err != nil || !token.Valid {
    
        if err == jwt.ErrSignatureInvalid {
            http.Error(w, "Invalid signature", http.StatusUnauthorized)
            return
        }
        if err == jwt.ErrTokenExpired {
            http.Error(w, "Token expired", http.StatusUnauthorized)
            return
        }
        // Handle other JWT validation errors
        http.Error(w, "Invalid JWT: "+err.Error(), http.StatusUnauthorized)
        return
    }

    
    var reqData struct {
        Questions []models.Question `json:"questions"`
        Answers []models.Answer   `json:"answers"`
    }
    err = json.NewDecoder(r.Body).Decode(&reqData)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    
    userID := token.Claims.(jwt.MapClaims)["user_id"].(int64)

    
    timestamp := time.Now().Format(time.RFC3339)

    
    salt, err := signature.GenerateSalt(32)
    if err != nil {
        http.Error(w, "Failed to generate salt", http.StatusInternalServerError)
        return
    }


    signature, err := signature.GenerateSignature(userID, timestamp, salt)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Store signature (replace with your actual storage logic)
    err = storage.StoreSignature(signature)
    if err != nil {
        http.Error(w, "Failed to store signature", http.StatusInternalServerError)
        return
    }

    // Send successful response with the generated signature
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"signature": signature})
}
