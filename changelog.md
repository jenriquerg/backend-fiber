# Changelog

Todas las actualizaciones importantes del proyecto se documentarán aquí.

## Version 0.0.1 First Commit 28/06/2025

Se creo el proyecto, se incluyen las librerias, conexion a la base de datos, controller, route y model de el primer crud de la tabla usuarios, archivo README del proyecto y la estructura inicial.

## Version 0.0.2 Apis 28/06/2025
Se crearon los endpoints para el resto de las tablas y las de autenticación, para registro y login, usando jwt, ademas de la implementacion de un rate limiter de solicitudes 10 de una misma ip por minuto, las apis siguen el modelo crud, excepto las de control que no ucentan con api de actualización debido a las reglas de negocio. 

## Version 0.0.3 Middlewares 29/07/2025
Se crearon los middlewares para el JSON  Schema y la proteccion de rutas, asi como la creación del token con permisos en la version "Checkpoint 51"

## Version 0.0.4 Logs 29/07/2025
Se crearon los middlewares para el registro de logs y se implementó