# Requisitos buscados para el gestor de tareas
1. Uso nativo del lenguaje del proyecto para definir tareas.
- Descripción: La herramienta debe ser capaz de integrarse directamente con el entorno de Go, permitiendo la creación y ejecución de tareas sin dependencias adicionales de lenguaje.
- Medición: Evaluar si la herramienta permite definir tareas usando Go de forma nativa o requiere archivos de configuración en otros lenguajes. Medir si tiene integración directa con go modules y si admite tareas en Go sin configuraciones adicionales.
- Importancia: Permitir que el equipo trabaje sin cambiar de lenguaje es crucial para mantener un flujo de trabajo eficiente y evita la necesidad de herramientas externas de traducción o adaptación.
2. Ausencia de dependencias externas.
- Descripción: La herramienta debe instalarse y configurarse sin dependencias adicionales y sin procesos complejos de configuración.
- Medición: Contar la cantidad de pasos de instalación y configuración inicial; número de dependencias externas necesarias (idealmente, ninguna).
- Importancia: Una configuración simple permite que el equipo pueda usar la herramienta rápidamente en nuevos entornos y evita problemas de compatibilidad.
3. Fácil mantenibilidad.
- Descripción: Las tareas deben ser fáciles de leer, modificar y extender en el tiempo sin producir deuda técnica.
- Medición: Evaluar si la sintaxis y estructura de las tareas permite comentarios claros, modularidad, y reutilización. Determinar si el código de las tareas es autodescriptivo y si se puede organizar de manera lógica.
- Importancia: Una estructura de tareas clara facilita el mantenimiento y la colaboración a largo plazo, incluso con cambios en el equipo.

# Opciones de gestor de tareas

**Mage**: Al estar escrito en Go, Mage permite a los desarrolladores que ya conocen este lenguaje definir tareas sin tener que aprender nuevas sintaxis (como la de Make) o utilizar archivos YAML. Esto facilita la mantenibilidad y reduce la deuda técnica.
[Documentación oficial](https://github.com/magefile/mage)
##### Ventajas:

- Al estar escrito en Go, es especialmente útil para proyectos desarrollados en este lenguaje.

- Ofrece un alto grado de mantenibilidad, ya que las tareas se definen en un archivo `magefile.go`, y pueden ser fácilmente adaptadas a distintos sistemas operativos.

- No tiene dependencias más allá de Go, lo que simplifica su configuración.

- Usamos el mismo lenguaje en el que estamos desarrollando el proyecto.

##### Desventajas:

- Requiere conocer Go, pero como nuestro proyecto está en este lenguaje, esta desventaja se minimiza.

- No es tan flexible con lenguajes fuera de Go.

**Make**: Aunque es una herramienta muy versátil y utilizada en muchos lenguajes, su sintaxis y uso de archivos Makefile no está adaptado a Go y puede ser menos intuitivo para quienes están acostumbrados a trabajar en Go. Además, puede generar cierta deuda técnica si el equipo necesita dedicar tiempo extra a aprender y adaptar la sintaxis para definir tareas básicas.
[Documentación de configuración(no oficial)](https://earthly.dev/blog/golang-makefile/)
##### Ventajas:

- Es ampliamente utilizado y bien conocido en la industria del software.

- Compatible con múltiples plataformas y lenguajes de programación.

##### Desventajas:

- Su sintaxis puede resultar confusa, dependiendo del tipo de proyecto, lo que puede afectar la claridad y mantenibilidad.

- No ofrece tanta modularidad como Mage.

**Task**: Aunque Task también está pensado para Go, el uso de archivos YAML puede complicar la flexibilidad en proyectos más complejos, ya que se está añadiendo una capa de abstracción que Mage resuelve de forma directa al programar las tareas en Go.
[Documentación oficial](https://taskfile.dev/)
##### Ventajas:

- Usa un archivo `.yml` para gestionar las tareas, lo que lo hace fácil de leer y manejar.

- Tiene una amplia documentación y es bastante popular dentro de la comunidad Go.

- Ofrece una gran cantidad de complementos (plugins) que facilitan la integración con Docker.

##### Desventajas:

- La utilización del archivo `.yml` puede hacer que sea menos flexible en términos de personalización de tareas o flujos de trabajo complejos.

- Para quienes ya tienen experiencia con Go, puede resultar menos extensible y configurable que Mage.

Aunque **Make** es la herramienta más comúnmente utilizada en muchos lenguajes, en nuestro caso, dado que trabajamos en un proyecto desarrollado en Go, **Task** y **Mage** son opciones más adecuadas. 

# Evaluación de las opciones
**Mage**
- Compatibilidad con Go: Alta. Mage permite escribir tareas directamente en Go, utilizando magefile.go sin dependencias externas.
- Configuración Simple: Alta. Mage se instala con una sola instrucción (go install) y no requiere configuración adicional.
- Mantenibilidad: Alta. Al escribir las tareas en Go, el código es modular y reutilizable, con documentación directa en el archivo de tareas.

**Make**
- Compatibilidad con Go: Baja. Make no está diseñado específicamente para Go y requiere definir las tareas en un archivo Makefile.
- Configuración Simple: Moderada. Make viene preinstalado en sistemas Unix, pero requiere configuración adicional en Windows.
- Mantenibilidad: Baja. La sintaxis de Makefiles es menos intuitiva, especialmente para proyectos en Go.

**Task**
- Compatibilidad con Go: Moderada-Alta. Task permite integrar comandos de Go y usa YAML para definir tareas, lo que requiere conocimientos adicionales de configuración.
- Configuración Simple: Moderada. Task se instala de manera simple, pero requiere un archivo .yml adicional para definir tareas.
- Mantenibilidad: Moderada. La sintaxis de YAML es clara, pero separa el lenguaje del proyecto (Go) de la definición de tareas, lo que puede dificultar la mantenibilidad en entornos avanzados.

### Justificación de la elección

 Aunque **Make** sea la opción más popular, basado en los criterios de decisión, **Mage** se presenta como la herramienta óptima para el proyecto, dada su alta compatibilidad con Go, simplicidad de configuración, y facilidad para mantener las tareas a largo plazo. Mage ofrece una experiencia de desarrollo fluida y específica para Go, alineándose con las necesidades y estructura del equipo técnico en este caso.

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