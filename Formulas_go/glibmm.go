package main

// Glibmm - C++ interface to glib
// Homepage: https://www.gtkmm.org/

import (
	"fmt"
	
	"os/exec"
)

func installGlibmm() {
	// Método 1: Descargar y extraer .tar.gz
	glibmm_tar_url := "https://download.gnome.org/sources/glibmm/2.80/glibmm-2.80.0.tar.xz"
	glibmm_cmd_tar := exec.Command("curl", "-L", glibmm_tar_url, "-o", "package.tar.gz")
	err := glibmm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glibmm_zip_url := "https://download.gnome.org/sources/glibmm/2.80/glibmm-2.80.0.tar.xz"
	glibmm_cmd_zip := exec.Command("curl", "-L", glibmm_zip_url, "-o", "package.zip")
	err = glibmm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glibmm_bin_url := "https://download.gnome.org/sources/glibmm/2.80/glibmm-2.80.0.tar.xz"
	glibmm_cmd_bin := exec.Command("curl", "-L", glibmm_bin_url, "-o", "binary.bin")
	err = glibmm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glibmm_src_url := "https://download.gnome.org/sources/glibmm/2.80/glibmm-2.80.0.tar.xz"
	glibmm_cmd_src := exec.Command("curl", "-L", glibmm_src_url, "-o", "source.tar.gz")
	err = glibmm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glibmm_cmd_direct := exec.Command("./binary")
	err = glibmm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libsigc++")
exec.Command("latte", "install", "libsigc++")
}
