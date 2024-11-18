# Requisitos buscados para el gestor de tareas
1. Explícito
- Ya que buscamos conocer nuevas herramientas, con lo que utilizar el gestor de tareas implícito de Go, no cumpliría con este objetivo de la asignatura. 

2. Configuración externa
- Requiere un archivo adicional que describa las tareas o dependencias (e.g., Makefile en make o magefile en mage).

3. Imperativo
- Debe de describirse los pasos a seguir para completar una tarea.

4. Genérico
- El gestor de tareas debe de ser genérico sin dejar de estar específicamente orientado a go.

5. Uso de DSLs (Domain-Specific Languages)
- La herramienta debe de ofrecer la ventaja de aprovechar las características del lenguaje mientras se definen las tareas.

6. Facilidad de integración y estandarización
- La herramienta debe de integrarse fácilmente con otros sistemas y herramientas dentro del ecosistema Go.

# Opciones de gestor de tareas

**Mage**: Al estar escrito en Go, Mage permite a los desarrolladores que ya conocen este lenguaje definir tareas sin tener que aprender nuevas sintaxis (como la de Make) o utilizar archivos YAML. Esto facilita la mantenibilidad y reduce la deuda técnica.
[Documentación oficial](https://github.com/magefile/mage)

**Make**: Aunque es una herramienta muy versátil y utilizada en muchos lenguajes, su sintaxis y uso de archivos Makefile no está adaptado a Go y puede ser menos intuitivo para quienes están acostumbrados a trabajar en Go. Además, puede generar cierta deuda técnica si el equipo necesita dedicar tiempo extra a aprender y adaptar la sintaxis para definir tareas básicas.
[Documentación de configuración(no oficial)](https://earthly.dev/blog/golang-makefile/)

**Task**: Aunque Task también está pensado para Go, el uso de archivos YAML puede complicar la flexibilidad en proyectos más complejos, ya que se está añadiendo una capa de abstracción que Mage resuelve de forma directa al programar las tareas en Go.
[Documentación oficial](https://taskfile.dev/)

### Justificación de la elección
La elección de **Mage** como gestor de tareas en un proyecto Go se justifica principalmente por su control y facilidad de integración dentro del ecosistema Go, a pesar de que Go ya ofrece herramientas integradas para tareas simples.
1. Mage, como gestor de tareas explícito, permite un control mucho más granular sobre el flujo de trabajo, a diferencia de las herramientas implícitas de Go, como go build y go test, que son eficientes pero limitadas para tareas complejas o personalizadas.
2. Mage requiere un archivo de configuración externo (magefile.go), lo que facilita la organización y gestión de tareas de forma explícita, mientras que Go depende de su sistema interno para tareas simples, lo que puede volverse limitado a medida que el proyecto crece.
3. Al ser imperativo, Mage permite definir tareas detalladas y secuenciales que se ejecutan con precisión, lo que es ideal cuando se requieren pasos específicos y controlados, a diferencia de los comandos más abstractos de Go.
4. Mage, aunque es un task runner genérico que utiliza Go como su DSL, está bien adaptado al ecosistema Go, aprovechando las ventajas del lenguaje para crear tareas personalizadas sin sacrificar rendimiento ni claridad.
5. Mage usa Go como un DSL para definir tareas, lo que lo que se alinea con las mejores prácticas del lenguaje, mientras que otras herramientas de Go no requieren un DSL pero limitan la capacidad de personalización.
6. Mage se integra fácilmente con las herramientas y bibliotecas de Go, y al estar escrito en Go, se alinea con el flujo de trabajo del equipo, facilitando la adopción y extensión del proyecto.

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