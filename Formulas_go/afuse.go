package main

// Afuse - Automounting file system implemented in userspace with FUSE
// Homepage: https://github.com/pcarrier/afuse/

import (
	"fmt"
	
	"os/exec"
)

func installAfuse() {
	// Método 1: Descargar y extraer .tar.gz
	afuse_tar_url := "https://github.com/pcarrier/afuse/archive/refs/tags/v0.5.0.tar.gz"
	afuse_cmd_tar := exec.Command("curl", "-L", afuse_tar_url, "-o", "package.tar.gz")
	err := afuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	afuse_zip_url := "https://github.com/pcarrier/afuse/archive/refs/tags/v0.5.0.zip"
	afuse_cmd_zip := exec.Command("curl", "-L", afuse_zip_url, "-o", "package.zip")
	err = afuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	afuse_bin_url := "https://github.com/pcarrier/afuse/archive/refs/tags/v0.5.0.bin"
	afuse_cmd_bin := exec.Command("curl", "-L", afuse_bin_url, "-o", "binary.bin")
	err = afuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	afuse_src_url := "https://github.com/pcarrier/afuse/archive/refs/tags/v0.5.0.src.tar.gz"
	afuse_cmd_src := exec.Command("curl", "-L", afuse_src_url, "-o", "source.tar.gz")
	err = afuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	afuse_cmd_direct := exec.Command("./binary")
	err = afuse_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
}
