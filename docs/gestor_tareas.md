# Requisitos buscados para el gestor de tareas
1. Lenguaje de programación del proyecto.
- Habrá que tener en cuenta las opciones disponibles para Go.
2. Herramientas que permitan mantener las mejores prácticas para el lenguaje que faciliten su uso y mantenimiento.
3. Reducción de la deuda técnica futura.
- Diseñar el sistema de manera que sea fácil de mantener y escalar, evitando soluciones a corto plazo que generen problemas a largo plazo.

# Opciones de gestor de tareas

**Mage**: Al estar escrito en Go, Mage permite a los desarrolladores que ya conocen este lenguaje definir tareas sin tener que aprender nuevas sintaxis (como la de Make) o utilizar archivos YAML. Esto facilita la mantenibilidad y reduce la deuda técnica.
[Documentación oficial](https://github.com/magefile/mage)

**Make**: Aunque es una herramienta muy versátil y utilizada en muchos lenguajes, su sintaxis y uso de archivos Makefile no está adaptado a Go y puede ser menos intuitivo para quienes están acostumbrados a trabajar en Go. Además, puede generar cierta deuda técnica si el equipo necesita dedicar tiempo extra a aprender y adaptar la sintaxis para definir tareas básicas.
[Documentación de configuración(no oficial)](https://earthly.dev/blog/golang-makefile/)

**Task**: Aunque Task también está pensado para Go, el uso de archivos YAML puede complicar la flexibilidad en proyectos más complejos, ya que se está añadiendo una capa de abstracción que Mage resuelve de forma directa al programar las tareas en Go.
[Documentación oficial](https://taskfile.dev/)

**Just**: Just es una herramienta de automatización escrita en Go que se utiliza para crear "Makefiles" más simples y comprensibles. Funciona de manera similar a Make pero se enfoca en ser más fácil de usar y entender. En lugar de tener que definir un archivo Makefile con comandos complicados, puedes usar un archivo justfile, que es similar a un archivo de script con tareas y dependencias.
[Documentacion oficial](https://github.com/casey/just)

### Justificación de la elección
Just es una herramienta ideal para un gestor de tareas en Go porque, al estar escrita en el mismo lenguaje, se integra perfectamente con el proyecto sin necesidad de dependencias externas. Su sintaxis sencilla facilita el mantenimiento y la automatización de tareas repetitivas como compilación, pruebas y ejecución del servidor, lo que mejora la consistencia y reduce errores humanos. Además, al automatizar estos procesos esenciales, Just ayuda a reducir la deuda técnica, asegurando que el sistema sea fácil de mantener y escalar a largo plazo.

### Resumen de las tareas:
- `just build`: Compila el proyecto y genera el binario en la carpeta bin.
- `just installDeps`: Instala las dependencias del proyecto utilizandgo modules.
- `just run`: Ejecuta el binario generado del proyecto.
- `just clean`: Limpia el binario y otros archivos generados.
- `just check`: Verifica que el código fuente esté correctamente formateado.