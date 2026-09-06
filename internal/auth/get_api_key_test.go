package auth

import (
	"net/http"
	"testing"
)

func Test_GetApiKey(t *testing.T) {
	_, err := GetAPIKey(nil)

	if err != nil {
		t.Errorf("No error, expected error")
	}

	_, err = GetAPIKey(http.Header{
		"Authorization": []string{},
	})

	if err == nil {
		t.Errorf("No error, expected error")
	}

}
