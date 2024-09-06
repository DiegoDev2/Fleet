package main

// Evince - GNOME document viewer
// Homepage: https://apps.gnome.org/Evince/

import (
	"fmt"
	
	"os/exec"
)

func installEvince() {
	// Método 1: Descargar y extraer .tar.gz
	evince_tar_url := "https://download.gnome.org/sources/evince/46/evince-46.3.1.tar.xz"
	evince_cmd_tar := exec.Command("curl", "-L", evince_tar_url, "-o", "package.tar.gz")
	err := evince_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	evince_zip_url := "https://download.gnome.org/sources/evince/46/evince-46.3.1.tar.xz"
	evince_cmd_zip := exec.Command("curl", "-L", evince_zip_url, "-o", "package.zip")
	err = evince_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	evince_bin_url := "https://download.gnome.org/sources/evince/46/evince-46.3.1.tar.xz"
	evince_cmd_bin := exec.Command("curl", "-L", evince_bin_url, "-o", "binary.bin")
	err = evince_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	evince_src_url := "https://download.gnome.org/sources/evince/46/evince-46.3.1.tar.xz"
	evince_cmd_src := exec.Command("curl", "-L", evince_src_url, "-o", "source.tar.gz")
	err = evince_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	evince_cmd_direct := exec.Command("./binary")
	err = evince_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: desktop-file-utils")
exec.Command("latte", "install", "desktop-file-utils")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: itstool")
exec.Command("latte", "install", "itstool")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: djvulibre")
exec.Command("latte", "install", "djvulibre")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gspell")
exec.Command("latte", "install", "gspell")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
exec.Command("latte", "install", "hicolor-icon-theme")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libgxps")
exec.Command("latte", "install", "libgxps")
	fmt.Println("Instalando dependencia: libhandy")
exec.Command("latte", "install", "libhandy")
	fmt.Println("Instalando dependencia: libsecret")
exec.Command("latte", "install", "libsecret")
	fmt.Println("Instalando dependencia: libspectre")
exec.Command("latte", "install", "libspectre")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
