package security

import (
	"os"
	"testing"
	"time"
)

var testLogFile *os.File

func TestMain(m *testing.M) {
	var err error
	testLogFile, err = os.OpenFile("test_results.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer testLogFile.Close()

	code := m.Run()
	os.Exit(code)
}

func logTestResult(testName string, passed bool) {
	status := "OK"
	if !passed {
		status = "FAIL"
	}
	testLogFile.WriteString(time.Now().Format("2006-01-02 15:04:05") + " | " + testName + " | " + status + "\n")
}

func TestEncryptDecryptAES(t *testing.T) {
	const secret = "clave-secreta-super-segura"
	tests := []struct {
		name      string
		plainText string
	}{
		{"Texto simple", "hola mundo"},
		{"Texto vacío", ""},
		{"Texto con símbolos", "¡Hola, 123! ¿Cómo estás?"},
		{"Texto largo", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passed := true
			encrypted, err := EncryptAES(tt.plainText, secret)
			if err != nil {
				t.Errorf("EncryptAES() error: %v", err)
				passed = false
			} else {
				decrypted, err := DecryptAES(encrypted, secret)
				if err != nil {
					t.Errorf("DecryptAES() error: %v", err)
					passed = false
				} else if decrypted != tt.plainText {
					t.Errorf("DecryptAES() = %v, want %v", decrypted, tt.plainText)
					passed = false
				}
			}
			logTestResult(tt.name, passed)
		})
	}
}
