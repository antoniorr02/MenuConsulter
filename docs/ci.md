# Requisitos buscados para el sistema de integración continua
1. La herramienta debe de ser gratuita (freemium).
2. La herramienta debe de poder conectarse a GitHub Checks API.

# Opciones disponibles

**[GitHub Actions](https://github.com/features/actions)**: Se trata de una herramienta gratuita y con completa integración en el propio Github.

**[Circle CI](https://circleci.com/)**: Se trata de una opción gratuita y que se puede conectar con la GitHub Checks API. Además se ha podido configurar con bastante facilidad y relativamente rápido con lo que será una opción muy a tener en cuenta para la decisión final.

**[TravisCI](https://app.travis-ci.com/)**: Solo tiene un perido de prueba gratuita con lo que se va a descartar al no cumplir con el criterio 1.

**[Semaphore](https://semaphoreci.com/)**: Al igual que Circle CI cumple los requisitos que se buscan, sin embargo, aún habiendo sido posible configurarlo muy facilmente, lo voy a descartar porque la interfaz donde se muestran los resultados de los test no me parece tan amigable como Circle CI.

**[Cirrus CI](https://cirrus-ci.org/)**: No es valido para nuestro proyecto, ya que para poder utilizarlo con un repositorio nos exige un pago mensual.

**[Buddy](https://buddy.works/)**: 

### Justificación de la elección
Como es necesario incluir al menos dos herramientas diferentes:

- En primer lugar utilizaremos el propio Github Actions al ser la herramienta estándar.

- 