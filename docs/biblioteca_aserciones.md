# Requisitos para la elección de la biblioteca de aserciones
1. Debe de ser una herramienta actualizada para no aumentar la deuda técnica.
2. La biblioteca de aserciones no debe de ser BDD, ya que para nuestro proyecto bastaría con aserciones simples, que permitan mensajes personalizados.

# Opciones adecuadas a los criterios de la biblioteca de aserciones

**Gomega**: Se trata de una biblioteca de aserciones con mucho soporte y actualización. Además, ofrece soporte para pruebas basadas en BDD y mensajes detallados para facilitar la depuración.
[Documentación oficial](https://github.com/onsi/gomega)

**Testify**: Ofrece aserciones sencillas y potentes, además de soporte para mensajes personalizados. Del mismo modo, se trata de una herramienta que sigue actualizada y en uso en la actualidad con lo que cumple con el primer requisito.
[Documentación oficial](https://github.com/stretchr/testify) 

**GoCheck**: Presenta una gran deuda técnica al no haber sido actualizada en más de 4 años, no cumpliendo con el primer requisito propuesto. Con ello no merece la pena seguir con su evaluación.
[Documentación oficial](https://github.com/go-check/check)

### Justificación de la biblioteca de aserciones.
La elección de **Testify** como herramienta de pruebas se justifica en que se trata de una herramienta ampliamente utilizada y mantenida en el ecosistema de Go, lo que asegura soporte a largo plazo y una actualización frecuente, siendo la última del mismo de hace menos de una semana. Ofrece aserciones claras sensillas y potentes. Además, su compatibilidad con la herramienta estándar, `go test` de Go la convierte en una opción confiable y versátil.
