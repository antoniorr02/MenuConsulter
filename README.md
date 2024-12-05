# MenuConsulter
Repositorio para la asignatura Infraestructura Virtual de Antonio Rodríguez Rodríguez. 

## Problema a resolver:

Soy un estudiante que tiene el problema de ser muy exquisito con la comida y quiero saber cuando se va a ofrecer comida que se adapte a mis preferencias, ya que pierdo mucho tiempo teniendo que entrar en la web de comedores universitarios y ver entre todo lo disponible si me gusta algo. Del mismo modo, me gustaría saber la calidad alimenticia de los platos que hay disponibles cada día: macronutrientes, calorías, etc. Ya que tengo como objetivo personal, con mi entrenamiento, subir de peso ganando masa muscular, con lo que debería tener alguna forma de adecuar mi dieta a dicho objetivo.

# Resolución del problema

## Descripción
Este proyecto es un sistema automatizado que ayuda a los estudiantes a encontrar comidas que se adapten a sus preferencias y objetivos nutricionales. El sistema extrae los menús diarios de diferentes comedores disponibles, filtra los platos según las preferencias del usuario y proporciona un análisis nutricional de los platos.
El estudiante evita perder tiempo revisando los menús manualmente y puede seguir un plan nutricional adecuado a sus necesidades de entrenamiento.

## [Journey](https://github.com/antoniorr02/MenuConsulter/docs/journeys.md)

## [Historias de Usuario](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-1/docs/historias_usuario.md)

## [Hitos del Proyecto](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-1/docs/milestones.md)

## [Gestor de tareas](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-3/docs/gestor_tareas.md)

### Resumen de las tareas:
- `just build`: Compila el proyecto y genera el binario en la carpeta bin.
- `just install_deps`: Instala las dependencias del proyecto utilizandgo modules.
- `just clean`: Limpia el binario y otros archivos generados.
- `just check`: Verifica que el código fuente esté correctamente formateado.
- `just test`: Ejecuta los test para comprobar el correcto funcionamiento del proyecto.

## [Gestor de dependencias](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-3/docs/gestor_dependencias.md)

## [Documentación del proceso de desarrollo con test](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-4/docs/documentación_proceso_test.md)

## [Biblioteca de aserciones](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-4/docs/biblioteca_aserciones.md)

## [Test runner](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-4/docs/test_runner.md)

## [Imagen docker](https://github.com/antoniorr02/MenuConsulter/blob/Objetivo-5/docs/imagen.md)

### Generación y ejecución de la imagen
**Construir la imagen:** `docker build -t <nombre_imagen> .`
**Ejecutar la imagen**: `docker run --rm <nombre_imagen>`