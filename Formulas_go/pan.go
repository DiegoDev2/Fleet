package main

// Pan - Usenet newsreader that's good at both text and binaries
// Homepage: https://gitlab.gnome.org/GNOME/pan

import (
	"fmt"
	
	"os/exec"
)

func installPan() {
	// Método 1: Descargar y extraer .tar.gz
	pan_tar_url := "https://gitlab.gnome.org/GNOME/pan/-/archive/v0.160/pan-v0.160.tar.bz2"
	pan_cmd_tar := exec.Command("curl", "-L", pan_tar_url, "-o", "package.tar.gz")
	err := pan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pan_zip_url := "https://gitlab.gnome.org/GNOME/pan/-/archive/v0.160/pan-v0.160.tar.bz2"
	pan_cmd_zip := exec.Command("curl", "-L", pan_zip_url, "-o", "package.zip")
	err = pan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pan_bin_url := "https://gitlab.gnome.org/GNOME/pan/-/archive/v0.160/pan-v0.160.tar.bz2"
	pan_cmd_bin := exec.Command("curl", "-L", pan_bin_url, "-o", "binary.bin")
	err = pan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pan_src_url := "https://gitlab.gnome.org/GNOME/pan/-/archive/v0.160/pan-v0.160.tar.bz2"
	pan_cmd_src := exec.Command("curl", "-L", pan_src_url, "-o", "source.tar.gz")
	err = pan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pan_cmd_direct := exec.Command("./binary")
	err = pan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: enchant")
exec.Command("latte", "install", "enchant")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gmime")
exec.Command("latte", "install", "gmime")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: gtkspell3")
exec.Command("latte", "install", "gtkspell3")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
}
