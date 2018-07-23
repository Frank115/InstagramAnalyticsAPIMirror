# InstagramAnalyticsAPI
![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

Esta es una API hecha en GO que se encarga de manejar datos de provenientes de Instagram y asi obtener grandes cantidades de informacion de usuarios y publicaciones con el fin de crear metricas e informes utiles para los consumidores de la aplicacion.

## Para comenzar a trabajar

Crear en tu `$HOME/go` la siguiente estructura de directorios

```
  ├──go
      ├── src
          ├── github.com
              ├── ProyectoIT
                  ├── Clonar proyecto aca
```

### Para iniciar el proyecto

Para correr el proyecto

parado en `$HOME/go`/src/github.com/ProyectoIT/InstagramAnalyticsAPI/ correr:

```
go run main.go
```

### Como trabajamos

El proyecto utiliza el patron de diseño MVCS (Model-View-Controller-Service) para eso tiene la siguiente estructura interna de carpetas:

  ```
  ├── app -- Configuracion Inicial del API y mapeo de endpoints.
  ├── config -- Configuracion de los servicios, base de datos y constantes del API.
  ├── controller -- Aca estan las funciones que reciben la informacion de los endpoints
  ├── domain -- Aca estan las estructuras de datos del API, su declaracion y metodos.
  ├── service -- Aca esta la logica de la aplicacion y el principal funcionamiento del API
  ├── dao -- Aca se encuentran los servicios de persistencia del API (Base de Datos)
  ├── main.go -- Punto inicial de la Aplicacion.
```
