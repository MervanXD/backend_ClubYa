package tipos

import (
	"testing"

	"github.com/MervanXD/backend_ClubYa/logs"
)

func TestEnumString(t *testing.T) {
	logs.InitLogger()
	defer logs.CloseLogger()
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"TipoSocio válido", TipoSocio(0).String(), "Socio"},
		{"TipoSocio inválido", TipoSocio(99).String(), "Desconocido"},
		{"Estado válido", Estado(0).String(), "Confirmada"},
		{"Estado inválido", Estado(99).String(), "Desconocido"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("String() = %v, want %v", tt.value, tt.expected)
			}
		})
	}
}

func TestEnumMarshalJSON(t *testing.T) {
	logs.InitLogger()
	defer logs.CloseLogger()
	tests := []struct {
		name     string
		value    interface{ MarshalJSON() ([]byte, error) }
		expected string
	}{
		{"TipoSocio válido", TipoSocio(0), `"Socio"`},
		{"TipoSocio inválido", TipoSocio(99), `"Desconocido"`},
		{"Estado válido", Estado(0), `"Confirmada"`},
		{"Estado inválido", Estado(99), `"Desconocido"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.value.MarshalJSON()
			if string(got) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(got), tt.expected)
			}
		})
	}
}

func TestEnumUnmarshalJSON(t *testing.T) {
	logs.InitLogger()
	defer logs.CloseLogger()
	tests := []struct {
		name     string
		input    []byte
		enum     interface{}
		expected string
	}{
		{"TipoSocio válido", []byte(`"Socio"`), new(TipoSocio), "Socio"},
		{"TipoSocio inválido", []byte(`"NoExiste"`), new(TipoSocio), "Desconocido"},
		{"Estado válido", []byte(`"Confirmada"`), new(Estado), "Confirmada"},
		{"Estado inválido", []byte(`"NoExiste"`), new(Estado), "Desconocido"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch e := tt.enum.(type) {
			case *TipoSocio:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
				}
			case *Estado:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
				}
			}
		})
	}
}
