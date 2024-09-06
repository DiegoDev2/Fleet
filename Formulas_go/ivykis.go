package main

// Ivykis - Async I/O-assisting library
// Homepage: https://sourceforge.net/projects/libivykis/

import (
	"fmt"
	
	"os/exec"
)

func installIvykis() {
	// Método 1: Descargar y extraer .tar.gz
	ivykis_tar_url := "https://github.com/buytenh/ivykis/archive/refs/tags/v0.43.2-trunk.tar.gz"
	ivykis_cmd_tar := exec.Command("curl", "-L", ivykis_tar_url, "-o", "package.tar.gz")
	err := ivykis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ivykis_zip_url := "https://github.com/buytenh/ivykis/archive/refs/tags/v0.43.2-trunk.zip"
	ivykis_cmd_zip := exec.Command("curl", "-L", ivykis_zip_url, "-o", "package.zip")
	err = ivykis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ivykis_bin_url := "https://github.com/buytenh/ivykis/archive/refs/tags/v0.43.2-trunk.bin"
	ivykis_cmd_bin := exec.Command("curl", "-L", ivykis_bin_url, "-o", "binary.bin")
	err = ivykis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ivykis_src_url := "https://github.com/buytenh/ivykis/archive/refs/tags/v0.43.2-trunk.src.tar.gz"
	ivykis_cmd_src := exec.Command("curl", "-L", ivykis_src_url, "-o", "source.tar.gz")
	err = ivykis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ivykis_cmd_direct := exec.Command("./binary")
	err = ivykis_cmd_direct.Run()
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
