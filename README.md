# Proyecto de Asignación de Tareas

Este proyecto implementa una API que gestiona la  asignación de tareas a empleados utilizando Golang con el framework Fiber, y siguiendo los principios de la arquitectura hexagonal.

## Descripción del Proyecto

El sistema permite asignar tareas a empleados basándose en sus habilidades(skills) y disponibilidad. También genera reportes de asignaciones para fechas específicas.

### Características Principales

- Asignación automática de tareas a empleados
- Generación de reportes de asignaciones
- API RESTful construida con Fiber
- Arquitectura hexagonal para una mejor separación de responsabilidades

## Arquitectura Hexagonal

El proyecto sigue la arquitectura hexagonal (también conocida como "Ports and Adapters") con la siguiente estructura:

``` bash
TALATASK/
├── cmd/
│   └── main.go
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
- `config`: Maneja la configuración de la aplicación.

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

GET /report?date=2024-10-19

*utilizar la fecha del día actual o día posterior. 

## Pruebas

Para ejecutar las pruebas:

go test ./...

