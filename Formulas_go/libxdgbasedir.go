package main

// LibxdgBasedir - C implementation of the XDG Base Directory specifications
// Homepage: https://github.com/devnev/libxdg-basedir

import (
	"fmt"
	
	"os/exec"
)

func installLibxdgBasedir() {
	// Método 1: Descargar y extraer .tar.gz
	libxdgbasedir_tar_url := "https://github.com/devnev/libxdg-basedir/archive/refs/tags/libxdg-basedir-1.2.3.tar.gz"
	libxdgbasedir_cmd_tar := exec.Command("curl", "-L", libxdgbasedir_tar_url, "-o", "package.tar.gz")
	err := libxdgbasedir_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxdgbasedir_zip_url := "https://github.com/devnev/libxdg-basedir/archive/refs/tags/libxdg-basedir-1.2.3.zip"
	libxdgbasedir_cmd_zip := exec.Command("curl", "-L", libxdgbasedir_zip_url, "-o", "package.zip")
	err = libxdgbasedir_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxdgbasedir_bin_url := "https://github.com/devnev/libxdg-basedir/archive/refs/tags/libxdg-basedir-1.2.3.bin"
	libxdgbasedir_cmd_bin := exec.Command("curl", "-L", libxdgbasedir_bin_url, "-o", "binary.bin")
	err = libxdgbasedir_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxdgbasedir_src_url := "https://github.com/devnev/libxdg-basedir/archive/refs/tags/libxdg-basedir-1.2.3.src.tar.gz"
	libxdgbasedir_cmd_src := exec.Command("curl", "-L", libxdgbasedir_src_url, "-o", "source.tar.gz")
	err = libxdgbasedir_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxdgbasedir_cmd_direct := exec.Command("./binary")
	err = libxdgbasedir_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
