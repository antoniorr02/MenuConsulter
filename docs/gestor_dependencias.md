# Gestor de dependencias

A partir de `Go 1.11` (lanzado en 2018), se introdujo el concepto de módulos en Go, los cuales son colecciones de paquetes relacionados que se versionan conjuntamente, proporcionando todas las dependencias necesarias para la base de código a través del **Go Module**. Con ello, los gestores de dependencias externos como **Dep** o **Glide**, estos han quedaron obsoletos.

En Go no existe un gestor de dependencias externo como tal. La gestión de dependencias se realiza directamente con **Go Module**, mediante los siguientes archivos:
- `go.mod`: Es el archivo donde se define el path del módulo y se especifican las dependencias necesarias para el proyecto.
- `go.sum`: Este archivo se encarga de gestionar las versiones y garantizar la integridad de las dependencias.