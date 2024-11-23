# Documentación del proceso de desarrollo del objetivo con test

Una vez que hemos realizado la elección de la herramienta de test y teniendo en cuenta que la metodología de desarrollo con test va a ser TDD, y que con ello, en primer lugar se desarrollan los test y posteriormente el código asociado al presente milestone. Vamos a proceder en primer lugar con el planteamiento de los test que se van a realizar.

## Tets que se van a implementar
1. Validación de tipos de menús: Verificar que solo se pueden crear menús con tipos válidos según el archivo tipos.json.
2. Restricción del número de platos: Verificar que no se puedan agregar más de 5 platos a un menú.
3. Creación de un menú válido: Confirmar que un menú válido se crea correctamente.
4. Extracción de menús desde la web: Simular la extracción de datos de los menús de la página oficial de los comedores y verificar su transformación a estructuras de datos internas.
5. Almacenamiento y recuperación de menús en una base de datos: Probar que los menús extraídos se almacenan correctamente en la base de datos.

Con los test ya planteados, queda su implementación en menu_test.go y extraccion_test.go.
Posteriormente el desarrollo del código del milestone 1; y por último, comprobar que todo funcione correctamente.