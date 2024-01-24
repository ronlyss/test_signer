// storage/storage.go
package storage

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/yourusername/test-signer/model"
)

var (
	ErrTestNotFound    = errors.New("test not found")
	ErrSignatureInvalid = errors.New("signature invalid")
)

type Storage interface {
	SignTest(test *model.Test) (*model.Signature, error)
	VerifySignature(signature *model.Signature) (*model.Test, error)
}

type MemoryStorage struct {
	tests      map[string]*model.Test
	signatures map[string]*model.Signature
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tests:      make(map[string]*model.Test),
		signatures: make(map[string]*model.Signature),
	}
}

func (s *MemoryStorage) SignTest(test *model.Test) (*model.Signature, error) {
	
	if _, exists := s.tests[test.UserID]; exists {
		return nil, errors.New("test for user already exists")
	}


	s.tests[test.UserID] = test

	/
	testHash := hashTest(test)

	signature := &model.Signature{
		UserID:      test.UserID,
		Signature:   "some_generated_signature", 
		Completion:  time.Now(),
		TestHash:    testHash,
	}

	
	s.signatures[test.UserID] = signature

	return signature, nil
}

func (s *MemoryStorage) VerifySignature(signature *model.Signature) (*model.Test, error) {
	
	storedTest, testExists := s.tests[signature.UserID]
	storedSignature, signatureExists := s.signatures[signature.UserID]

	if !testExists || !signatureExists {
		return nil, ErrTestNotFound
	}

	
	if storedSignature.TestHash != signature.TestHash {
		return nil, ErrSignatureInvalid
	}

	return storedTest, nil
}

func hashTest(test *model.Test) string {
	
	data := fmt.Sprintf("%s|%v|%v", test.UserID, test.Questions, test.Answers)
	hash := sha256.Sum256([]byte(data))
	return base64.StdEncoding.EncodeToString(hash[:])
}
