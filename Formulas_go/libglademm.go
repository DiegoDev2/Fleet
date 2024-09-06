package main

// Libglademm - C++ wrapper around libglade
// Homepage: https://gnome.org

import (
	"fmt"
	
	"os/exec"
)

func installLibglademm() {
	// Método 1: Descargar y extraer .tar.gz
	libglademm_tar_url := "https://download.gnome.org/sources/libglademm/2.6/libglademm-2.6.7.tar.bz2"
	libglademm_cmd_tar := exec.Command("curl", "-L", libglademm_tar_url, "-o", "package.tar.gz")
	err := libglademm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libglademm_zip_url := "https://download.gnome.org/sources/libglademm/2.6/libglademm-2.6.7.tar.bz2"
	libglademm_cmd_zip := exec.Command("curl", "-L", libglademm_zip_url, "-o", "package.zip")
	err = libglademm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libglademm_bin_url := "https://download.gnome.org/sources/libglademm/2.6/libglademm-2.6.7.tar.bz2"
	libglademm_cmd_bin := exec.Command("curl", "-L", libglademm_bin_url, "-o", "binary.bin")
	err = libglademm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libglademm_src_url := "https://download.gnome.org/sources/libglademm/2.6/libglademm-2.6.7.tar.bz2"
	libglademm_cmd_src := exec.Command("curl", "-L", libglademm_src_url, "-o", "source.tar.gz")
	err = libglademm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libglademm_cmd_direct := exec.Command("./binary")
	err = libglademm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gtkmm")
exec.Command("latte", "install", "gtkmm")
	fmt.Println("Instalando dependencia: libglade")
exec.Command("latte", "install", "libglade")
}
