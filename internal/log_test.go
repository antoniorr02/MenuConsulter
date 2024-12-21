package internal

import (
	"bytes"
	"log"
	"testing"
)

func TestLoggerInitialization(t *testing.T) {
	var buffer bytes.Buffer

	logger := log.New(&buffer, "TEST: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Printf("Esto es un mensaje de prueba")
	if buffer.Len() == 0 {
		t.Errorf("El logger no registró ningún mensaje")
	}
}
