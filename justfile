CODE_FOLDERS := "./internal"

# Tarea para construir el proyecto
build:
    echo "Construyendo el proyecto..."
    go build {{CODE_FOLDERS}}

# Tarea para instalar las dependencias
install:
    echo "Instalando las dependencias..."
    go mod download

# Tarea para comprobar la sintaxis del proyecto
check:
    echo "Comprobando sintaxis del proyecto..."
    gofmt -e {{CODE_FOLDERS}} > /dev/null

test:
    echo "Ejecutando tests del proyecto..."
    go test ./internal -v
