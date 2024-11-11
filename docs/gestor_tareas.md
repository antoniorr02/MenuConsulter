# Gestor de tareas

Para la automatización de tareas repetitivas y facilitar el proceso de compilación y ejecución de nuestro proyecto, es necesario elegir un **Task Runner**. A continuación, exploramos algunas de las opciones más interesantes que hemos considerado:

- **Make**: Es una de las herramientas más populares para la automatización de tareas. Utiliza un archivo `Makefile` donde se definen las reglas para ejecutar las tareas. Su ventaja principal es su gran versatilidad, ya que se basa en el uso de la shell, lo que le permite ser utilizado con diferentes lenguajes. Esto lo convierte en una opción bastante flexible.

- **Mage**: Es una herramienta de automatización escrita en Go, que permite definir tareas directamente en este lenguaje. Estas tareas se configuran en el archivo `magefile.go` y se pueden ejecutar desde la línea de comandos.

- **Task**: Otra herramienta para automatizar tareas en Go. En este caso, las tareas se definen en un archivo `Taskfile.yml` y también se ejecutan desde la línea de comandos.

#### Mage
##### Ventajas:

- Al estar escrito en Go, es especialmente útil para proyectos desarrollados en este lenguaje.

- Ofrece un alto grado de mantenibilidad, ya que las tareas se definen en un archivo `magefile.go`, y pueden ser fácilmente adaptadas a distintos sistemas operativos.

- No tiene dependencias más allá de Go, lo que simplifica su configuración.

- Usamos el mismo lenguaje en el que estamos desarrollando el proyecto.

##### Desventajas:

- Requiere conocer Go, pero como nuestro proyecto está en este lenguaje, esta desventaja se minimiza.

- No es tan flexible con lenguajes fuera de Go.

#### Make
##### Ventajas:

- Es ampliamente utilizado y bien conocido en la industria del software.

- Compatible con múltiples plataformas y lenguajes de programación.

##### Desventajas:

- Su sintaxis puede resultar confusa, dependiendo del tipo de proyecto, lo que puede afectar la claridad y mantenibilidad.

- No ofrece tanta modularidad como Mage.

#### Task
##### Ventajas:

- Usa un archivo `.yml` para gestionar las tareas, lo que lo hace fácil de leer y manejar.

- Tiene una amplia documentación y es bastante popular dentro de la comunidad Go.

- Ofrece una gran cantidad de complementos (plugins) que facilitan la integración con Docker.

##### Desventajas:

- La utilización del archivo `.yml` puede hacer que sea menos flexible en términos de personalización de tareas o flujos de trabajo complejos.

- Para quienes ya tienen experiencia con Go, puede resultar menos extensible y configurable que Mage.

Aunque **Make** es la herramienta más comúnmente utilizada en muchos lenguajes, en nuestro caso, dado que trabajamos en un proyecto desarrollado en Go, **Task** y **Mage** son opciones más adecuadas. 

### Elección

 Aunque **Make** sea la opción más popular, considero que **Mage** es la opción más acertada para este proyecto, ya que está diseñado en Go. Esto nos permitirá definir tareas de forma directa en el mismo lenguaje, lo que favorece la integración y facilita el mantenimiento.

### Instalación de mage
- `go install github.com/magefile/mage@latest`
- [Documentación oficial](https://github.com/magefile/mage)

### Resumen de las tareas:
- `mage build`: Compila el proyecto y genera el binario en la carpeta bin.
- `mage installDeps`: Instala las dependencias del proyecto utilizandgo modules.
- `mage run`: Ejecuta el binario generado del proyecto.
- `mage clean`: Limpia el binario y otros archivos generados.
- `mage check`: Verifica que el código fuente esté correctamente formateado.