
![Logo](./Logo.png)

## Fleet

**Fleet** es un gestor de paquetes creado con Go, diseñado para instalar, gestionar y configurar diferentes herramientas de manera simple y eficiente. Permite descargar e instalar herramientas directamente desde los repositorios.

### Características

- Instalación de herramientas ligera y eficiente.
- Comandos simples para instalar, actualizar y eliminar herramientas.
- Fórmulas personalizables para diferentes herramientas.

### Instalación

1. Clona el repositorio:

```bash
   git clone https://github.com/DiegoDev2/Fleet
   cd Fleet
   go build -o fleet
   sudo mv fleet /usr/local/bin

```
### Uso

**Para instalar una herramienta, usa:**

```bash
fleet install <nombre-herramienta>
```

Por ejemplo, para instalar nmap:

```bash
fleet install nmap
```
### Añadir Nuevas Herramientas
- Abre `lib/structTool.go`,`lib/toolStruc.go` y añade la nueva herramienta a la lista de herramientas disponibles.

- Crea una fórmula para la herramienta en el    directorio `formulas/`.

- Usa el comando install para agregarla a tu sistema.

### Contribuir

Agradecemos contribuciones para mejorar Fleet. Si encuentras un error o tienes una sugerencia, no dudes en enviar un pull request o abrir un issue.
Licencia

Fleet está licenciado bajo la licencia Apache 2.0. Consulta el archivo `LICENSE` para más detalles. 

![imagen](https://github.com/user-attachments/assets/48411519-e46f-4bef-a786-6f49ec94b255)

