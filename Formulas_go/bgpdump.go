package main

// Bgpdump - C library for analyzing MRT/Zebra/Quagga dump files
// Homepage: https://github.com/RIPE-NCC/bgpdump/wiki

import (
	"fmt"
	
	"os/exec"
)

func installBgpdump() {
	// Método 1: Descargar y extraer .tar.gz
	bgpdump_tar_url := "https://github.com/RIPE-NCC/bgpdump/archive/refs/tags/v1.6.2.tar.gz"
	bgpdump_cmd_tar := exec.Command("curl", "-L", bgpdump_tar_url, "-o", "package.tar.gz")
	err := bgpdump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bgpdump_zip_url := "https://github.com/RIPE-NCC/bgpdump/archive/refs/tags/v1.6.2.zip"
	bgpdump_cmd_zip := exec.Command("curl", "-L", bgpdump_zip_url, "-o", "package.zip")
	err = bgpdump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bgpdump_bin_url := "https://github.com/RIPE-NCC/bgpdump/archive/refs/tags/v1.6.2.bin"
	bgpdump_cmd_bin := exec.Command("curl", "-L", bgpdump_bin_url, "-o", "binary.bin")
	err = bgpdump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bgpdump_src_url := "https://github.com/RIPE-NCC/bgpdump/archive/refs/tags/v1.6.2.src.tar.gz"
	bgpdump_cmd_src := exec.Command("curl", "-L", bgpdump_src_url, "-o", "source.tar.gz")
	err = bgpdump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bgpdump_cmd_direct := exec.Command("./binary")
	err = bgpdump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
}
