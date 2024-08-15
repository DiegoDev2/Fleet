# Turn

**Turn** es una herramienta de línea de comandos (CLI) para gestionar paquetes con Homebrew en macOS. Esta herramienta permite instalar, desinstalar y actualizar paquetes de manera sencilla y rápida. Además, soporta la selección de idioma entre inglés y español.

## Instalación

Para instalar **Turn**, clona el repositorio y compílalo:

```bash
git clone https://github.com/CodeDiego15/Turn
cd turn
chmod +x install.sh
./install.sh
```

## Uso

**Turn** ofrece los siguientes comandos:

- `install [paquete]`: Instala un paquete utilizando Homebrew.
- `uninstall [paquete]`: Desinstala un paquete utilizando Homebrew.
- `upgrade [paquete]`: Actualiza un paquete utilizando Homebrew.

### Ejemplos

Instalar un paquete:

```bash
turn install <nombre_del_paquete>
```

Desinstalar un paquete:

```bash
turn uninstall <nombre_del_paquete>
```

Actualizar un paquete:

```bash
turn upgrade <nombre_del_paquete>
```

### Selección de Idioma

La primera vez que ejecutes **Turn**, se te pedirá que selecciones un idioma entre inglés (`en`) y español (`es`). La preferencia de idioma se guardará para futuras ejecuciones.

## Agradecimientos

Queremos expresar nuestro agradecimiento a **Homebrew** por proporcionar una excelente plataforma de gestión de paquetes para macOS. Sin Homebrew, **Turn** no sería posible.

## Contribuciones

¡Las contribuciones son bienvenidas! Si encuentras errores o tienes sugerencias, por favor, abre un **issue** o envía un **pull request**.

---

**Turn** es un proyecto independiente y no está asociado oficialmente con Homebrew.
