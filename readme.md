
estructura del proyecto

backend-project
├── cmd
│   └── main.go                # Punto de entrada de la aplicación
├── internal
│   ├── adapters               # Adaptadores para interactuar con el mundo exterior
│   │   ├── auth               # Adaptador para JWT y Auth0
│   │   ├── profile            # Adaptador para el perfil
│   │   ├── events             # Adaptador para eventos
│   │   ├── dashboard          # Adaptador para el dashboard
│   │   ├── providers          # Adaptador para proveedores
│   │   └── repository         # Adaptador para repositorios
│   ├── domain                 # Lógica de negocio y entidades del dominio
│   │   ├── entities           # Clases y estructuras del dominio
│   │   └── services           # Servicios del dominio
│   ├── infrastructure         # Infraestructura y configuraciones
│   │   ├── database           # Configuración de la base de datos
│   │   ├── logging            # Configuración de logs
│   │   └── config             # Configuración general
│   ├── usecases               # Casos de uso
│   │   ├── ports              # Puertos (interfaces)
│   │   └── interactor         # Implementación de los casos de uso
│   └── api                    # Endpoints y middlewares
│       ├── routes             # Rutas de la API
│       └── middlewares        # Middlewares de la API
├── pkg                        # Paquetes reutilizables
│   └── framework              # Integración con Gin y otras librerías
├── go.mod                     # Archivo de dependencias de Go
└── README.md                  # Documentación del proyecto