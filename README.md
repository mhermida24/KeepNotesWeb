# 📝 Keep Me Notes

**Keep Me Notes** es una aplicación web diseñada para tomar notas rápidas y organizarlas de forma sencilla.  
Permite **crear, editar y eliminar notas personales**, además de **agruparlas en carpetas** para mantener todo bien estructurado.

---

## 🚀 Introducción

**Keep Me Notes** nace como una herramienta práctica para gestionar tus ideas, tareas y recordatorios en un solo lugar.  
Con una interfaz simple e intuitiva, podrás concentrarte en lo importante: **escribir y organizar tus pensamientos**.

La aplicación te permite:
- Crear nuevas notas.
- Editar notas existentes.
- Eliminar notas que ya no necesites.
- Organizar tus notas dentro de carpetas y subcarpetas.

---

## 📂 Estructura de la información

El sistema se basa en **dos entidades principales**:  
👉 **Notas** y **Carpetas**

### 🗒️ Nota
Cada nota representa un contenido individual y contiene la siguiente información:

- **Título:** breve encabezado que resume el contenido de la nota.  
- **Fecha:** fecha de creación o última modificación.  
- **Cuerpo:** texto principal, que puede incluir texto, imágenes u otros elementos multimedia.

### 📁 Carpeta
Las carpetas permiten organizar y agrupar las notas relacionadas. Cada carpeta incluye:

- **Nombre:** título identificador de la carpeta.  
- **Descripción:** breve resumen del propósito o contenido de la carpeta.  
- **Notas:** una colección de notas que pertenecen a la carpeta.  
- **Sub-Carpetas:** otras carpetas dentro de ella, permitiendo una organización jerárquica.

---

## 🔗 Relación entre entidades

Cada **nota** pertenece a una **carpeta**, y cada **carpeta** puede contener múltiples **notas**.  
Esta relación permite clasificar el contenido fácilmente (por ejemplo, en carpetas como `Trabajo`, `Estudios`, `Personal`, etc.).


# Guía de instalación y ejecución
## Instalar Go
- Descargar el instalador desde la página oficial: https://go.dev/dl/
- Ejecutar el instalador y seguir los pasos.
- Verificar la instalación abriendo una terminal y escribiendo:

```
go version
```
Deberías ver algo como:
```
go version go1.20.5 windows/amd64
```
## Descargar el proyecto
- Clonar el repositorio desde GitHub:
```
git clone https://github.com/asoutrelle/KeepNotesWeb.git
```
- Entrar a la carpeta del proyecto
## Ejecutar la aplicación
Dentro de la carpeta del proyecto, abre una terminal y ejecuta:
```
go run main.go
```
- Esto compila y ejecuta la aplicación.
