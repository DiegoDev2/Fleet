package main

// LibxmlxxAT3 - C++ wrapper for libxml
// Homepage: https://libxmlplusplus.github.io/libxmlplusplus/

import (
	"fmt"
	
	"os/exec"
)

func installLibxmlxxAT3() {
	// Método 1: Descargar y extraer .tar.gz
	libxmlxxat3_tar_url := "https://download.gnome.org/sources/libxml++/3.2/libxml++-3.2.5.tar.xz"
	libxmlxxat3_cmd_tar := exec.Command("curl", "-L", libxmlxxat3_tar_url, "-o", "package.tar.gz")
	err := libxmlxxat3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxmlxxat3_zip_url := "https://download.gnome.org/sources/libxml++/3.2/libxml++-3.2.5.tar.xz"
	libxmlxxat3_cmd_zip := exec.Command("curl", "-L", libxmlxxat3_zip_url, "-o", "package.zip")
	err = libxmlxxat3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxmlxxat3_bin_url := "https://download.gnome.org/sources/libxml++/3.2/libxml++-3.2.5.tar.xz"
	libxmlxxat3_cmd_bin := exec.Command("curl", "-L", libxmlxxat3_bin_url, "-o", "binary.bin")
	err = libxmlxxat3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxmlxxat3_src_url := "https://download.gnome.org/sources/libxml++/3.2/libxml++-3.2.5.tar.xz"
	libxmlxxat3_cmd_src := exec.Command("curl", "-L", libxmlxxat3_src_url, "-o", "source.tar.gz")
	err = libxmlxxat3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxmlxxat3_cmd_direct := exec.Command("./binary")
	err = libxmlxxat3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glibmm@2.66")
	exec.Command("latte", "install", "glibmm@2.66").Run()
}
