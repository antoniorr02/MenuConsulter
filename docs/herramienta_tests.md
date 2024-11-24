# Requisitos buscados para la herramienta de tests
1. Debe de ser una herramienta actualizada para no aumentar la deuda técnica futura.
2. La herramienta debe de ajustarse a la metodología TDD.
3. La herramienta debe permitir ver el código cubierto por los test.
4. La herramienta debe de permitir la separación en directorios para pruebas, de forma que se pueda escribir en algún lugar con permiso, teniendo en cuenta de esta forma el siguiente objetivo en la decisión actual.
5. La herramienta debe de permitir la realización de aserciones.

# Opciones de herramientas de tests

**Ginkgo**: Framework avanzado para pruebas en BDD, muy estructurado y adecuado para proyectos grandes. Compatible con herramientas de aserciones como Gomega.
[Documentación oficial](https://github.com/onsi/ginkgo)

**GoMock**: Ideal para generar mocks basados en interfaces, facilitando la simulación de dependencias y pruebas unitarias más precisas.
[Documentación oficial](https://github.com/golang/mock)

**Goblin**: Framework inspirado en Mocha, centrado en pruebas de estilo BDD. Ofrece sintaxis limpia y fluida, ideal para escribir pruebas de forma estructurada.
[Documentación oficial](https://github.com/franela/goblin)

**Testify**: Uno de los paquetes más populares en Go. Ofrece aserciones sencillas, mocks integrados y soporte para requisitos como la separación de directorios y pruebas en TDD.
[Documentación oficial](https://github.com/stretchr/testify) 

**GoCheck**: Extiende el paquete testing de go con funcionalidades avanzadas como agrupación de tests y soporte para fixtures, facilitando el trabajo en proyectos más complejos.
[Documentación oficial](https://github.com/go-check/check)

**Gomega**: Librería de aserciones expresiva, diseñada para complementarse con Ginkgo. Permite construir pruebas altamente legibles y adaptadas al enfoque BDD.
[Documentación oficial](https://github.com/onsi/gomega)

**Go-cmp**: Excelente para comparar estructuras de datos complejas, ofreciendo personalización avanzada, pero no incluye soporte para mocks o frameworks más completos.
[Documentación oficial](https://github.com/google/go-cmp)

**Convey**: Framework enfocado en BDD, con un entorno interactivo que permite visualizar resultados y organizar tests de forma legible y modular.
[Documentación oficial](https://github.com/smartystreets/goconvey)

**Snyk Advisor**: Se trata de un herramienta para evaluar dependencias en términos de seguridad, mantenimiento y calidad, útil para decisiones informadas en el ciclo de desarrollo, se ha incluido ya que si bien no es una herramienta que sustituya a las anteriormente expuesta, se comentó en clase durante la explicación del presente objetivo como una herramienta que puede ser interesante a tener en cuenta al tener un rol preventivo, que nos ayudará a seleccionar dependencias robustas y minimizar riesgos de seguridad.
[Documentación oficial](https://snyk.io/advisor/golang)

### Justificación de la elección
La elección de **Testify** como herramienta de pruebas se justifica por su equilibrio entre simplicidad y funcionalidad. Es una herramienta ampliamente utilizada y mantenida en el ecosistema de Go, lo que asegura soporte a largo plazo y una actualización frecuente, siendo la última del mismo de hace menos de una semana. Compatible con TDD, ofrece aserciones claras, integración con mocks, y flexibilidad para ejecutar pruebas desde directorios separados, ajustándose así a los requisitos del presente y del siguiente objetivo. Además, su compatibilidad con la herramienta estándar, `go test` de Go la convierte en una opción confiable y versátil.
Por otro lado, aunque Testify en sí mismo no incluye funcionalidades para medir la cobertura, al usarlo junto con el comando de pruebas nativo de Go, puedes generar informes de cobertura con `go test -cover`.