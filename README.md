# Backend Fiber con Supabase

Sistema backend para gestión de citas y reportes hospitalarios usando GoFiber, GORM y Supabase (PostgreSQL).

---

## Requisitos

- Go 1.24+
- Supabase (proyecto creado con base de datos PostgreSQL)
- [Go Modules](https://blog.golang.org/using-go-modules)

---

## Instalación

Instalación

1. Clona el repositorio:

    git clone https://github.com/jenriquerg/backend-fiber.git
    cd backend-fiber

2. Instala dependencias:

    go mod download

3. Configura variables de entorno

    Crea un archivo .env en la raíz con:

    DATABASE_URL=postgresql://usuario:contraseña@host.supabase.co:6543/postgres
    PORT=3000

4. Corre la aplicación:

    go run main.go

## Estructura del proyecto

  backend-fiber/
  ├── config/ # Configuración de base de datos
  ├── controllers/ # Lógica de negocio (CRUD)
  ├── models/ # Modelos GORM
  ├── routes/ # Definición de rutas
  ├── main.go # Archivo principal
  ├── go.mod # Módulo y dependencias
  └── .env # Variables de entorno (no versionar)



## Seguridad

  Nunca publiques el archivo .env ni compartas la contraseña de tu base de datos.

  Usa HTTPS en producción.

  Implementa validaciones y sanitización de inputs.

## Contacto

  Jesús Enrique Rojas
  github.com/jenriquerg