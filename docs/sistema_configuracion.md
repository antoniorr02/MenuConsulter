# Requisitos buscados para el sistema de configuración
1. Puntuación de [Snyk Advisor](https://snyk.io/) la cuál pone puntuaciones sobre 100 dependiendo de la seguridad, popularidad, mantenimiento y comunidad de diferentes herramientas como lo son por ejemplo los sistemas de configuración.
2. Debe de ser una herramienta actualizada para no aumentar la deuda técnica.

# Opciones disponibles

**[Viper](https://github.com/spf13/viper)**: Tiene una valoración en Snyk de 95 sobre 100 y ha sido actualizada la última semana, con lo que será una opción a tener en cuenta.

**[Koanf](https://github.com/knadh/koanf)**: Tiene una valoración en Snyk de 89 sobre 100, siendo una diferencia casi despreciable respecto a viper y ha sido actualizada hace menos de 4 días.

**[Kong](https://github.com/alecthomas/kong)**: Esta herramienta la he encontrado en reddit como alternativa a viper con la ventaja sobre viper y cobra que tiene muchas menos dependencias, del mismo modo, ciñéndonos a nuestros criterios, su puntuación en Snyk es de 95 sobre 100 igualando a Viper y también ha sido actualizada en los últimos días.

### Justificación de la elección
Finalmente Kong que parecía la más prometedora ha resultado ser la más tediosa de configurar, con lo que finalmente como viper y koanf cumplen todos los requisitos, he optado por quedarme con viper que considero que es la herramienta que más se suele utilizar con golang.

