# Requisitos buscados para el gestor de tareas
1. Lenguaje de programación: Go.
2. Mejores prácticas y herramientas recomendadas para Go.
3. Reducción de la deuda técnica futura.

# Opciones de gestor de tareas

**Mage**: Al estar escrito en Go, Mage permite a los desarrolladores que ya conocen este lenguaje definir tareas sin tener que aprender nuevas sintaxis (como la de Make) o utilizar archivos YAML. Esto facilita la mantenibilidad y reduce la deuda técnica.
[Documentación oficial](https://github.com/magefile/mage)

**Make**: Aunque es una herramienta muy versátil y utilizada en muchos lenguajes, su sintaxis y uso de archivos Makefile no está adaptado a Go y puede ser menos intuitivo para quienes están acostumbrados a trabajar en Go. Además, puede generar cierta deuda técnica si el equipo necesita dedicar tiempo extra a aprender y adaptar la sintaxis para definir tareas básicas.
[Documentación de configuración(no oficial)](https://earthly.dev/blog/golang-makefile/)

**Task**: Aunque Task también está pensado para Go, el uso de archivos YAML puede complicar la flexibilidad en proyectos más complejos, ya que se está añadiendo una capa de abstracción que Mage resuelve de forma directa al programar las tareas en Go.
[Documentación oficial](https://taskfile.dev/)

### Justificación de la elección
Elegir **Mage** en Go es una alternativa muy adecuada pues simplifica la gestión de tareas con sintaxis nativa y aprovechando las mejores prácticas y herramientas del ecosistema Go. Esto, junto con el soporte que la herramienta tiene, reduce la deuda técnica futura al mantener la consistencia en el lenguaje y facilitar el mantenimiento.

### Instalación de mage
- `go install github.com/magefile/mage@latest`
- [Documentación oficial](https://github.com/magefile/mage)
Es posible que necesites ejecutar previamente:
`nano ~/.zshrc`
`export PATH=$PATH:$(go env GOPATH)/bin`
`source ~/.zshrc`
Comprueba la instalación con: `mage -v` 

### Resumen de las tareas:
- `mage build`: Compila el proyecto y genera el binario en la carpeta bin.
- `mage installDeps`: Instala las dependencias del proyecto utilizandgo modules.
- `mage run`: Ejecuta el binario generado del proyecto.
- `mage clean`: Limpia el binario y otros archivos generados.
- `mage check`: Verifica que el código fuente esté correctamente formateado.