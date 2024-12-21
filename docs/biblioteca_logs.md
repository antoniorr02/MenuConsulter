# Requisitos buscados para la biblioteca de logs
1. Debe de ser un sistema configurable el cual permita tanto personalizar la salida de los logs con el formato que se desee como almacenar los logs en diferentes lugares (terminal o ficheros).
2. Debe de ser una herramienta actualizada para no aumentar la deuda técnica.

# Opciones disponibles

**[Log](https://pkg.go.dev/log)**: Se trata de la biblioteca estándar para el registro de logs en Go, cumple con todos los requisitos exigidos con lo que será una herramienta muy a tener en cuenta, ya que si bien no es la que más variedad ofrece, se trata de la biblioteca la estándar y cumple con lo mínimo que se busca en nuestro proyecto.

**[Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)**:

**[Zerolog](https://pkg.go.dev/github.com/rs/zerolog)**:

**[Zap](https://pkg.go.dev/go.uber.org/zap)**:

**[Slog](https://go.dev/blog/slog)**:

### Justificación de la elección