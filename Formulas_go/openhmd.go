package main

// Openhmd - Free and open source API and drivers for immersive technology
// Homepage: http://openhmd.net

import (
	"fmt"
	
	"os/exec"
)

func installOpenhmd() {
	// Método 1: Descargar y extraer .tar.gz
	openhmd_tar_url := "https://github.com/OpenHMD/OpenHMD/archive/refs/tags/0.3.0.tar.gz"
	openhmd_cmd_tar := exec.Command("curl", "-L", openhmd_tar_url, "-o", "package.tar.gz")
	err := openhmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openhmd_zip_url := "https://github.com/OpenHMD/OpenHMD/archive/refs/tags/0.3.0.zip"
	openhmd_cmd_zip := exec.Command("curl", "-L", openhmd_zip_url, "-o", "package.zip")
	err = openhmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openhmd_bin_url := "https://github.com/OpenHMD/OpenHMD/archive/refs/tags/0.3.0.bin"
	openhmd_cmd_bin := exec.Command("curl", "-L", openhmd_bin_url, "-o", "binary.bin")
	err = openhmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openhmd_src_url := "https://github.com/OpenHMD/OpenHMD/archive/refs/tags/0.3.0.src.tar.gz"
	openhmd_cmd_src := exec.Command("curl", "-L", openhmd_src_url, "-o", "source.tar.gz")
	err = openhmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openhmd_cmd_direct := exec.Command("./binary")
	err = openhmd_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: hidapi")
exec.Command("latte", "install", "hidapi")
}
