package main

// Digitemp - Read temperature sensors in a 1-Wire net
// Homepage: https://www.digitemp.com/

import (
	"fmt"
	
	"os/exec"
)

func installDigitemp() {
	// Método 1: Descargar y extraer .tar.gz
	digitemp_tar_url := "https://github.com/bcl/digitemp/archive/refs/tags/v3.7.2.tar.gz"
	digitemp_cmd_tar := exec.Command("curl", "-L", digitemp_tar_url, "-o", "package.tar.gz")
	err := digitemp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	digitemp_zip_url := "https://github.com/bcl/digitemp/archive/refs/tags/v3.7.2.zip"
	digitemp_cmd_zip := exec.Command("curl", "-L", digitemp_zip_url, "-o", "package.zip")
	err = digitemp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	digitemp_bin_url := "https://github.com/bcl/digitemp/archive/refs/tags/v3.7.2.bin"
	digitemp_cmd_bin := exec.Command("curl", "-L", digitemp_bin_url, "-o", "binary.bin")
	err = digitemp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	digitemp_src_url := "https://github.com/bcl/digitemp/archive/refs/tags/v3.7.2.src.tar.gz"
	digitemp_cmd_src := exec.Command("curl", "-L", digitemp_src_url, "-o", "source.tar.gz")
	err = digitemp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	digitemp_cmd_direct := exec.Command("./binary")
	err = digitemp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libusb-compat")
	exec.Command("latte", "install", "libusb-compat").Run()
}
