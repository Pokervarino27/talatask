# Proyecto de Asignación de Tareas

Este proyecto implementa una API que gestiona la  asignación de tareas a empleados utilizando Golang con el framework Fiber, y siguiendo los principios de la arquitectura hexagonal.

### Descripción del Proyecto

El sistema permite asignar tareas a empleados basándose en sus habilidades(skills) y disponibilidad. También genera reportes de asignaciones para fechas específicas.

## Enfoque de la Solución
### asignación de la tarea:
El enfoque está basado en la pre filtración de los componentes, primero del empleado donde se pre-filtra por su disponibilidad y habilidades, con esto se evita recorrer el listado completo de empleados. Se generan 2 estructuras de mapas de habilidades y disponibilidad, el primero asocia a los empleados disponibles para un día, el segundo asocia los empleados con la habilidad que poseen, si bien hay 2 "for" aninados, se reduce el tiempo al hacerlo con este pre-filtrado a diferencia si se recorriera como una matriz de 2 dimensiones. 
Luego se ordenan las tareas por fecha de forma ascendente para empezar por la más próxima, se valida previamente que la tarea no haya sido asignada, se busca a los empleados disponibles para esa fecha en el Mapa de disponibilidad previamente creado  y luego se llama a un función que filtra por habilidad.
Esta habilidad nuevamente crea otro mapa donde guarda como key el ID del empleado y mantiene un contador incremental por cada habilidad requerida que posee este. Luego se recorre el listado de empleados disponibles y por cada empleado se verifica si la cantidad de habilidades son iguales a las que se requieren, si es así se agrega a la lista de empleados seleccionados finalmente se retorna el listado de empleados seleccionados. Finalmente si se encontraron empleados la tarea se le asigna el primer empleado de la lista filtrada que se retornó y se marca la tarea como asignada y se entrega el listado de asignaciones. 

**Observaciones:** Se abordó un enfoque de construcción de pre-filtrado para hacer búsquedas más acotadas y evitar recorrer bucles completos de forma innecesaria haciendo búsquedas repetitivas. Aún quedan otros enfoques más eficientes como por ejemplo utilización del Algoritmo Húngaro creando matrices de costes o algoritmos en programación de restricciones con herramientras como las de Google OR-Tools ***(mi último trabajo consitió en asignación automática de rutas a conductores con restricciones como distancias, ventanas horarias, capacidad de vehículo, etc)***

### Generación de reporte:
Recibe una fecha como parámetro en la llamada al Endpoint que será el filtro para el reporte que se generará para ese día, nuevamente se utiliza el enfoque de creación de mapas con pre-filtrado en este caso para agrupar las tareas por fecha por el día ingresado como parámetro esto optimizará el listado de tareas según esta fecha específica, luego por cada empleado se crea un reporte con su identificación, tareas y horas asignadas al inicio vacías, se revisan las tareas asignadas y se agrega un nuevo campo para el reporte con las horas totales y restantes para el empleado. Otro enfoque posible sería guardar el formato del reporte en base de datos e ir actualizando cuando hayan actualizaciones en las asignaciones y luego hacer búsqueda por fecha, una base de datos de tipo documentos podría servir para esta información. 

### Características Principales

- Asignación automática de tareas a empleados
- Generación de reportes de asignaciones
- API RESTful construida con Fiber
- Arquitectura hexagonal para una mejor separación de responsabilidades

## Arquitectura Hexagonal

El proyecto sigue la arquitectura hexagonal con la siguiente estructura:

``` bash
TALATASK/
├── cmd/
│   └── routes/
|   |   |__ routes.go
|   |__ main.go
├── internal/
│   ├── domain/
│   │   ├── employee.go
│   │   ├── task.go
│   │   └── assignment.go
│   ├── ports/
│   │   ├── repositories.go
│   │   └── services.go
│   ├── adapters/
│   │   ├── repositories/
│   │   │   ├── employee_repository.go
│   │   │   └── task_repository.go
│   │   └── handlers/
│   │       └── http_handler.go
│   ├── usecases/
│   │   └── task_assignment.go
|   |   └── task_assignment_test.go  
│   └── config/
│       └── config.go
└── go.mod
``` 

- `domain`: Contiene las entidades principales y reglas de negocio.
- `ports`: Define las interfaces para los repositorios y servicios.
- `adapters`: Implementa las interfaces definidas en `ports`.
- `usecases`: Contiene la lógica de negocio principal.
- `infraestructura`: Maneja la configuración de la aplicación y conexiones con otro tipo de componentes, base de datos, mensajería, etc.

## Tecnologías Utilizadas

- Go 1.22
- Fiber v2
- Docker y Docker Compose

## Configuración del Proyecto

### Prerrequisitos

- Go 1.22
- Docker y Docker Compose

### Instalación

1. Clonar el repositorio:
2. Instalar dependencias:
    `go mod tidy`

### Con Docker Compose

Para ejecutar el proyecto usando Docker Compose:
    `docker-compose up --build`

## Uso de la API

### Asignar Tareas

Para asignar tareas, haz una petición POST a `/assign-tasks`

### Generar Reporte

Para generar un reporte de asignaciones, haz una petición GET a `/report`:

GET /report?date=2024-10-21

*utilizar la fecha del día actual o día posterior. 

## Pruebas

Para ejecutar las pruebas:

go test ./...

