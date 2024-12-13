# Requisitos buscados para la imagen para el contenedor
1. Imagen actualizada de forma que no se aumente la deuda técnica futura
2. Preferencia por imágenes ligeras (ej. Alpine), disminuyendo el requerimiento de disco.

# Opciones de imagen

**[Golang](https://hub.docker.com/_/golang)**: Se trata de una imagen que trae go instalado lo que elimina la necesidad de configurarlo a parte, y se encuentra actualizada recientemente lo que disminuye la deuda técnica. Del mismo modo, es una herramienta muy versátil que permite incluso combinar Golang con Alpine.

**[Alpine](https://hub.docker.com/_/alpine)**: Es una imagen muy liviana, pero no es tan sencilla como ubuntu al tener una librería estándar diferente (tal y como se ha comentado en clase), lo que puede complicar su uso. Del mismo modo, exige que instalemos y configuremos go nosotros mismos.

**[DockerSlim](https://github.com/kcq/docker-slim)**: Es un poco más grande que Alpine, si bien no deja de ser muy ligera y es una diferencia despreciable. Por otro lado, la última actualización tiene más de 7 meses, mientras que alpine y golang se actualizan con mayor frecuencia.

**[Ubuntu](https://hub.docker.com/_/ubuntu)**: Se va a descartar ya que es la imagen más grande de las tres que se han encontrado.

**[Bitnami Ubuntu](https://hub.docker.com/r/bitnami/ubuntu-base-buildpack)**: Esta imagen pesa más que otras alternativas, con lo que se va a descartar.

**[Okteto golang](https://github.com/okteto/okteto)**: Esta imagen pesa mucho más que otras alternativas, casi tanto como la alternativa de golang, con lo que se va a descartar.

**[Circleci golang](https://github.com/CircleCI-Archived/circleci-dockerfiles/blob/master/golang/images/1.10.0-stretch/Dockerfile)**: Ocupa incluso más que okteto, descartada.

### Justificación de la elección
En base a los criterios y las pruebas realizadas en la carpeta [pruebas_imagenes](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-5/docs/pruebas_imagenes) me he decantado por utilizar `Golang`, con Alpine por su ligereza y velocidad de ejecución en comparación a las otras imagenes probadas. Si entramos [Golang-Alpine](https://hub.docker.com/_/golang/tags?page=1&name=alpine) vemos las diferentes opciones disponibles.

### Generación y ejecución de la imagen
**Construir la imagen:** `docker build -t <nombre_imagen> .`
**Ejecutar la imagen**: `docker run -u 1001 -t -v `pwd`:/app/test <nombre_imagen>`

**Desde docker hub**: `docker run -u 1001 -t -v `pwd`:/app/test antoniorr02/menuconsulter`