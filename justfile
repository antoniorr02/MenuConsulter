BINARY_NAME := "bin"
BIN := "./bin"
CODE_FOLDERS := "./internal"

# Tarea para construir el proyecto
build:
    echo "Construyendo el proyecto..."
    mkdir -p {{BIN}}
    go build -o {{BIN}}/{{BINARY_NAME}} {{CODE_FOLDERS}}

# Tarea para instalar las dependencias
install_deps:
    echo "Instalando las dependencias..."
    go mod tidy

# Tarea para limpiar los binarios
clean:
    echo "Limpiando los binarios..."
    rm -f {{BIN}}/{{BINARY_NAME}}
    go clean {{CODE_FOLDERS}}

# Tarea para comprobar la sintaxis del proyecto
check:
    echo "Comprobando sintaxis del proyecto..."
    gofmt -e {{CODE_FOLDERS}} > /dev/null

test:
    echo "Ejecutando tests del proyecto..."
    go test ./internal -v
