# Requisitos para la elección de la biblioteca de aserciones
1. Debe de ser una herramienta actualizada para no aumentar la deuda técnica.
2. Debe permitir emitir mensajes personalizados.

# Opciones adecuadas a los criterios de la biblioteca de aserciones

**Gomega**: Se trata de una biblioteca de aserciones que presenta una muy buena valoración en Snyk Advisor, al tener un gran mantenimiento y comunidad.
[Documentación oficial](https://github.com/onsi/gomega)

**Testify**: Uno de los paquetes más populares en Go. Ofrece aserciones sencillas, mocks integrados y soporte para requisitos como la separación de directorios y pruebas en TDD.
[Documentación oficial](https://github.com/stretchr/testify) 

**GoCheck**: Extiende el paquete testing de go con funcionalidades avanzadas como agrupación de tests y soporte para fixtures, facilitando el trabajo en proyectos más complejos. Si bien presenta una gran deuda técnica al no haber sido actualizada en más de 4 años, no cumpliendo con el segundo requisito propuesto.
[Documentación oficial](https://github.com/go-check/check)

**Convey**: Framework enfocado en BDD, con un entorno interactivo que permite visualizar resultados y organizar tests de forma legible y modular.
[Documentación oficial](https://github.com/smartystreets/goconvey)

### Justificación de la biblioteca de aserciones.
La elección de **Testify** como herramienta de pruebas se justifica por su equilibrio entre simplicidad y funcionalidad. Es una herramienta ampliamente utilizada y mantenida en el ecosistema de Go, lo que asegura soporte a largo plazo y una actualización frecuente, siendo la última del mismo de hace menos de una semana. Compatible con TDD, ofrece aserciones claras, integración con mocks, y flexibilidad para ejecutar pruebas desde directorios separados, ajustándose así a los requisitos del presente y del siguiente objetivo. Además, su compatibilidad con la herramienta estándar, `go test` de Go la convierte en una opción confiable y versátil.
Por otro lado, aunque Testify en sí mismo no incluye funcionalidades para medir la cobertura, al usarlo junto con el comando de pruebas nativo de Go, puedes generar informes de cobertura con `go test -cover`.