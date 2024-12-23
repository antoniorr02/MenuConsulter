# Requisitos buscados para la biblioteca de logs
1. Debe de ser un sistema configurable el cual permita tanto personalizar la salida de los logs con el formato que se desee como almacenar los logs en diferentes lugares (terminal o ficheros).
2. Debe de ser una herramienta actualizada para no aumentar la deuda técnica.
3. Si es posible la herramienta debe de ser implementable sin dependencias externas.

# Opciones disponibles

**[Log](https://pkg.go.dev/log)**: Se trata de la biblioteca estándar para el registro de logs en Go, cumple con todos los requisitos exigidos con lo que será una herramienta muy a tener en cuenta, ya que si bien no es la que más variedad ofrece, se trata de la biblioteca la estándar y cumple con lo mínimo que se busca en nuestro proyecto.

**[Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)**: Ofrece una mayor personalización que Log, sin embargo, requiere dependencias externas para su uso.

**[Zerolog](https://pkg.go.dev/github.com/rs/zerolog)**: Ofrece una funcionalidad muy parecida a Logrus, sin embargo, requiere dependencias externas para su uso. Además, he encontrado una sintaxis más amigable en logrus.

**[Zap](https://pkg.go.dev/go.uber.org/zap)**: Cumple con todos los requisitos con una gran personalización de los logs, sin embargo, requiere dependencias externas para su uso.

**[Slog](https://go.dev/blog/slog)**: Cumple con todos los requisitos propuestos, además, se trata del estándar del lenguaje desde Go 1.21, con lo que al estar más actualizado cumple mejor el segundo criterio que Log y supera muchas de sus limitaciones en cuanto a personalización.

### Justificación de la elección
La elección final va a ser Slog, ya que se trata de la biblioteca estándar en Go para login y mejora en personalización y funcionalidad respecto a Log.