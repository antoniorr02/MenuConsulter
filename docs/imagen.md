# Requisitos buscados para la imagen para el contenedor
1. Imagen actualizada de forma que no se aumente la deuda técnica futura
2. Prefencia por imágenes ligeras (ej. Alpine), disminuyendo el requerimiento de disco.
3. La imagen debe de traer el lenguaje del proyecto configurada por defecto, para reducir la configuración necesaria de la imagen.

# Opciones de imagen

**[Golang](https://hub.docker.com/_/golang)**: Se trata de una imagen que trae go instalado lo que elimina la necesidad de configurarlo a parte, y se encuentra actualizada recientemente lo que disminuye la deuda técnica. Del mismo modo, es una herramienta muy versátil que permite incluso combinar Golang con Alpine.

**[Alpine](https://hub.docker.com/_/alpine)**: Es una imagen muy liviana, pero no es tan sencilla como ubuntu al tener una librería estándar diferente (tal y como se ha comentado en clase), lo que puede complicar su uso. Del mismo modo, exige que instalemos y configuremos go nosotros mismos.

**[DockerSlim](https://github.com/kcq/docker-slim)**: Es un poco más grande que Alpine. Por otro lado, la última actualización tiene más de 7 meses, mientras que alpine y golang se actualizan con mayor frecuencia.

**[Ubuntu](https://hub.docker.com/_/ubuntu)**: Se va a descartar ya que es la imagen más grande de las tres que se han encontrado.

**[Bitnami para Golang](https://hub.docker.com/r/bitnami/golang/)**: La imagen que ofrece Bitnami para docker con Golang funciona sobre Debian, con lo que al igual que ocurre con Ubuntu se va a descartar por el peso de la misma. 

### Justificación de la elección
En base a los criterios me he decantado por utilizar `Golang 1.23` con Alpine. Si entramos [Golang-Alpine](https://hub.docker.com/_/golang/tags?page=1&name=alpine) vemos las diferentes opciones disponibles, concretamente usaremos [Golang 1.23.4:Alpine](https://hub.docker.com/layers/library/golang/1.23.3-alpine3.20/images/sha256-cdb47cf7cc930903987ead22e38dfb607db077bf120e740f7f5f14d1d18e4668?context=explore)

### Generación y ejecución de la imagen
**Construir la imagen:** `docker build -t <nombre_imagen> .`
**Ejecutar la imagen**: `docker run -u 1001 -t -v `pwd`:/app/test <nombre_imagen>`

**Desde docker hub**: `docker run -u 1001 -t -v `pwd`:/app/test antoniorr02/menuconsulter`