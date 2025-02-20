```markdown
# Arquitectura Hexagonal

Este proyecto sigue la arquitectura hexagonal. A continuación se muestra la estructura de directorios:

```
/home/manuonda/projects/go-projects/arquitectura-hexagonal/
├── cmd/
│   └── <nombre-del-comando>/
│       └── main.go
├── internal/
│   ├── entities/
│   │   └── <entidad>.go
│   ├── adapters/
│   │   ├── <adaptador>.go
│   └── services/
│       └── <servicio>.go
├── pkg/
│   └── <paquete>.go
├── README.md
└── go.mod
```

- **cmd/**: Contiene los puntos de entrada de la aplicación.
- **internal/**: Contiene el código interno de la aplicación.
    - **entities/**: Contiene las entidades del dominio.
    - **adapters/**: Contiene los adaptadores que interactúan con el mundo exterior.
    - **services/**: Contiene la lógica de negocio.
- **pkg/**: Contiene paquetes reutilizables.
- **README.md**: Este archivo.
- **go.mod**: Archivo de configuración de módulos de Go.
```


//https://carlos.lat/blog/hexagonal-architecture-using-go-fiber/