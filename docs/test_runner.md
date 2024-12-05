# Requisitos para la elección del test runner
1. Debe de cumplir con el estándar del lenguaje (si lo hubiera).
2. Debe de ser una herramienta actualizada para no aumentar la deuda técnica.
3. La herramienta debe de permitir realizar pruebas automáticas y agrupar las aserciones en subtests. 

# Opciones adecuadas a los criterios del test runner

**Ginkgo**: Framework avanzado para pruebas en BDD, muy estructurado y adecuado para proyectos grandes. Se complementa con herramientas de aserciones como Gomega.
[Documentación oficial](https://github.com/onsi/ginkgo)

**Go test**: Se trata de la herramienta que se usa según el estándar de golang.

### Justificación del test runner.
Si bien puede ser interesante evaluar alternativas como las que ofrece Ginkgo, haciendo que no sea necesaria una biblioteca de aserciones al complementarse con bibliotecas de aserciones como Gomega, me he decidido por utilizar el propio `go test` al tratarse de la herramienta estándar que se utiliza en proyectos en Golang.
Ahora si bien no es necesaria una biblioteca de aserciones al poder emplear los errores de go test para dicho propósito, me he decidido a investigar en las diferentes opciones que existen para este fin y trabajar sobre una herramienta diferente, permitiendo así desarrollar la capacidad de buscar herramientas de este tipo y su posterior implementación, lo cuál puede ser bastante interesante de cara al futuro.
