# Documentación del proceso de desarrollo del objetivo con test

Una vez que hemos realizado la elección de la herramienta de test y teniendo en cuenta que la metodología de desarrollo con test va a ser TDD, y que con ello, en primer lugar se desarrollan los test y posteriormente el código asociado al presente milestone. Vamos a proceder en primer lugar con el planteamiento de los test que se van a realizar.

## Tets que se van a implementar
1. Validación de tipos de menús: Verificar que solo se pueden crear menús con tipos válidos según el archivo tipos.json.
2. Restricción del número de platos: Verificar que no se puedan agregar más de 5 platos a un menú.
3. Confirmar que se descarga correctamente el html sobre el que se extraeran los datos de la página deseada. 
4. Extracción de menús desde el html: Simular la extracción de datos de los menús de los comedores y verificar su transformación a estructuras de datos internas.

Con los test ya planteados, queda su implementación en menu_test.go y extraccion_test.go.
Posteriormente el desarrollo del código del milestone 1; y por último, comprobar que todo funcione correctamente.