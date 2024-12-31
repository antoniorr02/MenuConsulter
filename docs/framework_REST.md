# Requisitos buscados para el framework REST
1. Puntuación de [Snyk Advisor](https://snyk.io/) la cuál pone puntuaciones sobre 100 dependiendo de la seguridad, popularidad, mantenimiento y comunidad de diferentes herramientas como lo son por ejemplo los frameworks REST.
2. Debe de ser una herramienta actualizada para no aumentar la deuda técnica.
3. Dada la sencillez del proyecto en el ámbito de una asignatura académica, optaremos con el framework que nos proporcionen la funcionalidad mínima necesaria para nuestro proyecto (manejar rutas), siendo por tanto un criterio la elección de la herramienta menos pesada entre las opciones que cumplan los criterios restantes.

# Opciones disponibles

**[Gin](https://github.com/gin-gonic/gin)**: Cuenta con una valoración en snyk de 93 sobre 100, lleva sin ser actualizada desde el mes pasado.

**[Echo](https://github.com/labstack/echo)**: Tiene una valoración de 97 sobre 100 y ha sido actualizada esta semana.

**[Chi](https://github.com/go-chi/chi)**: Tiene una valoración de 99 sobre 100 en Snyk y fue actualizada la semana pasada.

**[Fasthttp](https://github.com/valyala/fasthttp)**: Tiene una valoración de 97 sobre 100 en Snyk y ha sido actualizada por última vez hace 5 días.

**[Fiber](https://github.com/gofiber/fiber)**: Tiene una valoración de 100 sobre 100 en Snyk y ha sido actualizada en las últimas 24 horas, con lo que será una opción a tener bastante en cuenta.

### Justificación de la elección
Dado que me he centrado en buscar framewors REST que tengan buenas valoraciones en Snyk y que estén lo más actualizados posibles, ya que no tendría sentido incluir herramientas que de primeras sé que voy a eliminar, voy a hacer mi decisión teniendo en cuenta en que estamos en el ámbito de un proyecto pequeño a nivel de una asignatura universitaria con lo que el más básico de los frameworks propuestos sería más que suficiente. Con ello, tras ver la documentación de cada una de las herramientas pienso que **Chi** sería la herramienta más adecuada ya que incluye lo principal para manejar rutas y middleware; lo que nos será más que suficiente para lo que buscamos en nuestro proyecto. 