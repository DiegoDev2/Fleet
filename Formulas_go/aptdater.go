package main

// AptDater - Manage package updates on remote hosts using SSH
// Homepage: https://github.com/DE-IBH/apt-dater

import (
	"fmt"
	
	"os/exec"
)

func installAptDater() {
	// Método 1: Descargar y extraer .tar.gz
	aptdater_tar_url := "https://github.com/DE-IBH/apt-dater/archive/refs/tags/v1.0.4.tar.gz"
	aptdater_cmd_tar := exec.Command("curl", "-L", aptdater_tar_url, "-o", "package.tar.gz")
	err := aptdater_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aptdater_zip_url := "https://github.com/DE-IBH/apt-dater/archive/refs/tags/v1.0.4.zip"
	aptdater_cmd_zip := exec.Command("curl", "-L", aptdater_zip_url, "-o", "package.zip")
	err = aptdater_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aptdater_bin_url := "https://github.com/DE-IBH/apt-dater/archive/refs/tags/v1.0.4.bin"
	aptdater_cmd_bin := exec.Command("curl", "-L", aptdater_bin_url, "-o", "binary.bin")
	err = aptdater_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aptdater_src_url := "https://github.com/DE-IBH/apt-dater/archive/refs/tags/v1.0.4.src.tar.gz"
	aptdater_cmd_src := exec.Command("curl", "-L", aptdater_src_url, "-o", "source.tar.gz")
	err = aptdater_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aptdater_cmd_direct := exec.Command("./binary")
	err = aptdater_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
}
