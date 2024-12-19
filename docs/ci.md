# Requisitos buscados para el sistema de integración continua
1. La herramienta debe de ser gratuita (freemium).
2. La herramienta debe de poder conectarse a GitHub Checks API.

# Opciones disponibles

**[GitHub Actions](https://github.com/features/actions)**: Se trata de una herramienta gratuita y con completa integración en el propio Github.

**[Circle CI](https://circleci.com/)**: Se trata de una opción gratuita y que se puede conectar con la GitHub Checks API. Además se ha podido configurar con bastante facilidad y relativamente rápido con lo que será una opción muy a tener en cuenta para la decisión final.

**[TravisCI](https://app.travis-ci.com/)**: Solo tiene un perido de prueba gratuita con lo que se va a descartar al no cumplir con el criterio 1.

**[Semaphore](https://semaphoreci.com/)**: Al igual que Circle CI cumple los requisitos que se buscan, si bien la interfaz donde se muestran los resultados de los test no me parece tan amigable como Circle CI.

**[Cirrus CI](https://cirrus-ci.org/)**: No es valido para nuestro proyecto, ya que para poder utilizarlo con un repositorio privado nos exige un pago mensual.

**[Buddy](https://buddy.works/)**: Al igual que circle cumple nuestros dos requisitos, adicionalmente es el más facil de configurar, además el hecho de que se pueda configurar de forma completamente externa al programa, y sin tener que configurar casi nada del yml (se hace casi todo por interfaz), lo hace a mi parecer la opción más llamativa de todas las que he probado.

### Justificación de la elección
Como es necesario incluir al menos dos herramientas diferentes:

- En primer lugar utilizaremos el propio Github Actions al ser la herramienta estándar. Por otro lado, en la configuración de github actions voy a probar lo que pide el objetivo para comprobar que varias versiones del lenguaje funcionan con nuestro proyecto. En mi caso he utilizado la versión 1.11 que es la primera que implementa módulos en Golang con lo que si funciona con esta versión debería de funcionar con la mayoría y la 1.23.4 que es una de las estables según la página oficial de Golang. 

- En segundo lugar tras haber probado todas las herramientas, me he decido por utilizar Buddy como herramienta de integraciñon continua. 
Adicionalmente, como Buddy tiene el problema de que se configura en la propia página web sin crear un fichero de configuración esto puede resultar en un problema si se está dearrollando el proyecto en un equipo de desarrollo, con ello, se va a también utilizar Circle CI como herramienta de integración continua. 