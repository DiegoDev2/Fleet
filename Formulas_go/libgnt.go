package main

// Libgnt - NCurses toolkit for creating text-mode graphical user interfaces
// Homepage: https://keep.imfreedom.org/libgnt/libgnt

import (
	"fmt"
	
	"os/exec"
)

func installLibgnt() {
	// Método 1: Descargar y extraer .tar.gz
	libgnt_tar_url := "https://downloads.sourceforge.net/project/pidgin/libgnt/2.14.3/libgnt-2.14.3.tar.xz"
	libgnt_cmd_tar := exec.Command("curl", "-L", libgnt_tar_url, "-o", "package.tar.gz")
	err := libgnt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgnt_zip_url := "https://downloads.sourceforge.net/project/pidgin/libgnt/2.14.3/libgnt-2.14.3.tar.xz"
	libgnt_cmd_zip := exec.Command("curl", "-L", libgnt_zip_url, "-o", "package.zip")
	err = libgnt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgnt_bin_url := "https://downloads.sourceforge.net/project/pidgin/libgnt/2.14.3/libgnt-2.14.3.tar.xz"
	libgnt_cmd_bin := exec.Command("curl", "-L", libgnt_bin_url, "-o", "binary.bin")
	err = libgnt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgnt_src_url := "https://downloads.sourceforge.net/project/pidgin/libgnt/2.14.3/libgnt-2.14.3.tar.xz"
	libgnt_cmd_src := exec.Command("curl", "-L", libgnt_src_url, "-o", "source.tar.gz")
	err = libgnt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgnt_cmd_direct := exec.Command("./binary")
	err = libgnt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gtk-doc")
exec.Command("latte", "install", "gtk-doc")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
