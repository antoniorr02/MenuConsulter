# Requisitos para la elección del test runner
1. Debe de cumplir con el estándar del lenguaje (si lo hubiera).
2. Debe de ser una herramienta actualizada para no aumentar la deuda técnica futura.
3. La herramienta debe de permitir realizar pruebas automáticas y agrupar las aserciones en subtests. 

# Opciones adecuadas a los criterios del test runner

**Ginkgo**: Framework avanzado para pruebas en BDD, muy estructurado y adecuado para proyectos grandes. Compatible con herramientas de aserciones como Gomega.
[Documentación oficial](https://github.com/onsi/ginkgo)

**Go test**: Se trata de la herramienta que se usa según el estándar de golang.

### Justificación del test runner.
Si bien puede ser interesante evaluar alternativas como las que ofrece Ginkgo, haciendo que no sea necesaria una biblioteca de aserciones, me he decidido por utilizar el propio `go test` al tratarse de la herramienta estándar que se utiliza en proyectos en Golang.
Ahora si bien no es necesaria una biblioteca de aserciones al poder emplear los errores de go test para dicho propósito, me he decidido a investigar en las diferentes opciones que existen para este fin y trabajar sobre una herramienta diferente, permitiendo así desarrollar la capacidad de buscar herramientas de este tipo y su posterior implementación, lo cuál puede ser bastante interesante de cara al futuro.

# Requisitos para la elección de la biblioteca de aserciones
1. Debe de cumplir con el estándar del lenguaje (si lo hubiera).
2. Debe de ser una herramienta actualizada para no aumentar la deuda técnica futura.
3. La herramienta debe de permitir la separación en directorios para pruebas, de forma que se pueda escribir en algún lugar con permiso, teniendo en cuenta de esta forma el siguiente objetivo en la decisión actual.
4. La herramienta debe de permitir la realización de aserciones.

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