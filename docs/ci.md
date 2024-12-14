# Requisitos buscados para el sistema de integración continua
1. La herramienta debe de ser gratuita (freemium).
2. La herramienta debe de poder conectarse a GitHub Checks API.

# Opciones disponibles

**[GitHub Actions](https://github.com/features/actions)**: Se trata de una herramienta gratuita y con completa integración en el propio Github.

**[Circle CI](https://circleci.com/)**: Se trata de una opción gratuita y que se puede conectar con la GitHub Checks API. Además se ha podido configurar con bastante facilidad y relativamente rápido con lo que será una opción muy a tener en cuenta para la decisión final.

**[TravisCI](https://app.travis-ci.com/)**: Solo tiene un perido de prueba gratuita con lo que se va a descartar al no cumplir con el criterio 1.

**[Semaphore](https://semaphoreci.com/)**:

**[Cirrus CI](https://cirrus-ci.org/)**:

**[Buddy]()**:

### Justificación de la elección
Como es necesario incluir al menos dos herramientas diferentes:

- En primer lugar utilizaremos el propio Github Actions al ser la herramienta estándar.

- 