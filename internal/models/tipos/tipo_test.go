package tipos

import (
	"os"
	"testing"
	"time"

	"github.com/MervanXD/backend_ClubYa/logs"
)

var testLogFile *os.File

func TestMain(m *testing.M) {
	// Abre el archivo en modo truncado (sobrescribe cada vez)
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

func TestEnumString(t *testing.T) {
	logs.InitLogger()
	defer logs.CloseLogger()
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		// TipoSocio
		{"TipoSocio válido", TipoSocio(0).String(), "Socio"},
		{"TipoSocio inválido", TipoSocio(99).String(), "Desconocido"},
		// Estado
		{"Estado válido", Estado(0).String(), "Confirmada"},
		{"Estado inválido", Estado(99).String(), "Desconocido"},
		// Sexo
		{"Sexo válido", Sexo(0).String(), "Femenino"},
		{"Sexo inválido", Sexo(99).String(), "Desconocido"},
		// TipoDocumento
		{"TipoDocumento válido", TipoDocumento(0).String(), "DNI"},
		{"TipoDocumento inválido", TipoDocumento(99).String(), "Desconocido"},
		// TipoTrabajo
		{"TipoTrabajo válido", TipoTrabajo(0).String(), "Independiente"},
		{"TipoTrabajo inválido", TipoTrabajo(99).String(), "Desconocido"},
		// TipoVia
		{"TipoVia válido", TipoVia(0).String(), "Jr."},
		{"TipoVia inválido", TipoVia(99).String(), "Desconocido"},
		// Ubicacion
		{"Ubicacion válido", Ubicacion(0).String(), "Principal"},
		{"Ubicacion inválido", Ubicacion(99).String(), "Desconocido"},
		// EstadoDisponibilidad
		{"EstadoDisponibilidad válido", EstadoDisponibilidad(0).String(), "No disponible"},
		{"EstadoDisponibilidad inválido", EstadoDisponibilidad(99).String(), "Desconocido"},
		// EstadoMembresia
		{"EstadoMembresia válido", EstadoMembresia(0).String(), "Vigente"},
		{"EstadoMembresia inválido", EstadoMembresia(99).String(), "Desconocido"},
		// EstadoSolicitud
		{"EstadoSolicitud válido", EstadoSolicitud(0).String(), "Pendiente"},
		{"EstadoSolicitud inválido", EstadoSolicitud(99).String(), "Desconocido"},
		// TipoMembresia
		{"TipoMembresia válido", TipoMembresia(0).String(), "Regular"},
		{"TipoMembresia inválido", TipoMembresia(99).String(), "Desconocido"},
		// UnidadFrecuenciaSem
		{"UnidadFrecuenciaSem válido", UnidadFrecuenciaSem(0).String(), "Sesiones"},
		{"UnidadFrecuenciaSem inválido", UnidadFrecuenciaSem(99).String(), "Desconocido"},
		// Actividad
		{"Actividad válido", Actividad(0).String(), "Sala de reuniones"},
		{"Actividad inválido", Actividad(99).String(), "Desconocido"},
		// MetodoPago
		{"MetodoPago válido", MetodoPago(0).String(), "Tarjeta"},
		{"MetodoPago inválido", MetodoPago(99).String(), "Desconocido"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passed := true
			if tt.value != tt.expected {
				t.Errorf("String() = %v, want %v", tt.value, tt.expected)
				passed = false
			}
			logTestResult(tt.name, passed)
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
		// Agrega aquí los mismos enums que arriba, ejemplo:
		{"TipoSocio válido", TipoSocio(0), `"Socio"`},
		{"TipoSocio inválido", TipoSocio(99), `"Desconocido"`},
		{"Estado válido", Estado(0), `"Confirmada"`},
		{"Estado inválido", Estado(99), `"Desconocido"`},
		{"Sexo válido", Sexo(0), `"Femenino"`},
		{"Sexo inválido", Sexo(99), `"Desconocido"`},
		{"TipoDocumento válido", TipoDocumento(0), `"DNI"`},
		{"TipoDocumento inválido", TipoDocumento(99), `"Desconocido"`},
		{"TipoTrabajo válido", TipoTrabajo(0), `"Independiente"`},
		{"TipoTrabajo inválido", TipoTrabajo(99), `"Desconocido"`},
		{"TipoVia válido", TipoVia(0), `"Jr."`},
		{"TipoVia inválido", TipoVia(99), `"Desconocido"`},
		{"Ubicacion válido", Ubicacion(0), `"Principal"`},
		{"Ubicacion inválido", Ubicacion(99), `"Desconocido"`},
		{"EstadoDisponibilidad válido", EstadoDisponibilidad(0), `"No disponible"`},
		{"EstadoDisponibilidad inválido", EstadoDisponibilidad(99), `"Desconocido"`},
		{"EstadoMembresia válido", EstadoMembresia(0), `"Vigente"`},
		{"EstadoMembresia inválido", EstadoMembresia(99), `"Desconocido"`},
		{"EstadoSolicitud válido", EstadoSolicitud(0), `"Pendiente"`},
		{"EstadoSolicitud inválido", EstadoSolicitud(99), `"Desconocido"`},
		{"TipoMembresia válido", TipoMembresia(0), `"Regular"`},
		{"TipoMembresia inválido", TipoMembresia(99), `"Desconocido"`},
		{"UnidadFrecuenciaSem válido", UnidadFrecuenciaSem(0), `"Sesiones"`},
		{"UnidadFrecuenciaSem inválido", UnidadFrecuenciaSem(99), `"Desconocido"`},
		{"Actividad válido", Actividad(0), `"Sala de reuniones"`},
		{"Actividad inválido", Actividad(99), `"Desconocido"`},
		{"MetodoPago válido", MetodoPago(0), `"Tarjeta"`},
		{"MetodoPago inválido", MetodoPago(99), `"Desconocido"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passed := true
			got, _ := tt.value.MarshalJSON()
			if string(got) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(got), tt.expected)
				passed = false
			}
			logTestResult(tt.name, passed)
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
		// Ejemplo para TipoSocio y Estado, repite para los demás enums:
		{"TipoSocio válido", []byte(`"Socio"`), new(TipoSocio), "Socio"},
		{"TipoSocio inválido", []byte(`"NoExiste"`), new(TipoSocio), "Desconocido"},
		{"Estado válido", []byte(`"Confirmada"`), new(Estado), "Confirmada"},
		{"Estado inválido", []byte(`"NoExiste"`), new(Estado), "Desconocido"},
		{"Sexo válido", []byte(`"Femenino"`), new(Sexo), "Femenino"},
		{"Sexo inválido", []byte(`"NoExiste"`), new(Sexo), "Desconocido"},
		{"TipoDocumento válido", []byte(`"DNI"`), new(TipoDocumento), "DNI"},
		{"TipoDocumento inválido", []byte(`"NoExiste"`), new(TipoDocumento), "Desconocido"},
		{"TipoTrabajo válido", []byte(`"Independiente"`), new(TipoTrabajo), "Independiente"},
		{"TipoTrabajo inválido", []byte(`"NoExiste"`), new(TipoTrabajo), "Desconocido"},
		{"TipoVia válido", []byte(`"Jr."`), new(TipoVia), "Jr."},
		{"TipoVia inválido", []byte(`"NoExiste"`), new(TipoVia), "Desconocido"},
		{"Ubicacion válido", []byte(`"Principal"`), new(Ubicacion), "Principal"},
		{"Ubicacion inválido", []byte(`"NoExiste"`), new(Ubicacion), "Desconocido"},
		{"EstadoDisponibilidad válido", []byte(`"No disponible"`), new(EstadoDisponibilidad), "No disponible"},
		{"EstadoDisponibilidad inválido", []byte(`"NoExiste"`), new(EstadoDisponibilidad), "Desconocido"},
		{"EstadoMembresia válido", []byte(`"Vigente"`), new(EstadoMembresia), "Vigente"},
		{"EstadoMembresia inválido", []byte(`"NoExiste"`), new(EstadoMembresia), "Desconocido"},
		{"EstadoSolicitud válido", []byte(`"Pendiente"`), new(EstadoSolicitud), "Pendiente"},
		{"EstadoSolicitud inválido", []byte(`"NoExiste"`), new(EstadoSolicitud), "Desconocido"},
		{"TipoMembresia válido", []byte(`"Regular"`), new(TipoMembresia), "Regular"},
		{"TipoMembresia inválido", []byte(`"NoExiste"`), new(TipoMembresia), "Desconocido"},
		{"UnidadFrecuenciaSem válido", []byte(`"Sesiones"`), new(UnidadFrecuenciaSem), "Sesiones"},
		{"UnidadFrecuenciaSem inválido", []byte(`"NoExiste"`), new(UnidadFrecuenciaSem), "Desconocido"},
		{"Actividad válido", []byte(`"Sala de reuniones"`), new(Actividad), "Sala de reuniones"},
		{"Actividad inválido", []byte(`"NoExiste"`), new(Actividad), "Desconocido"},
		{"MetodoPago válido", []byte(`"Tarjeta"`), new(MetodoPago), "Tarjeta"},
		{"MetodoPago inválido", []byte(`"NoExiste"`), new(MetodoPago), "Desconocido"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passed := true
			switch e := tt.enum.(type) {
			case *TipoSocio:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *Estado:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *Sexo:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *TipoDocumento:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *TipoTrabajo:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *TipoVia:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *Ubicacion:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *EstadoDisponibilidad:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *EstadoMembresia:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *EstadoSolicitud:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *TipoMembresia:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *UnidadFrecuenciaSem:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *Actividad:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			case *MetodoPago:
				_ = e.UnmarshalJSON(tt.input)
				if e.String() != tt.expected {
					t.Errorf("UnmarshalJSON() = %v, want %v", e.String(), tt.expected)
					passed = false
				}
			}

			logTestResult(tt.name, passed)
		})
	}
}
