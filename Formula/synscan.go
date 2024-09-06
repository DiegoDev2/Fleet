package main

// Synscan - Asynchronous half-open TCP portscanner
// Homepage: https://digit-labs.org/files/tools/synscan/

import (
	"fmt"
	
	"os/exec"
)

func installSynscan() {
	// Método 1: Descargar y extraer .tar.gz
	synscan_tar_url := "https://digit-labs.org/files/tools/synscan/releases/synscan-5.02.tar.gz"
	synscan_cmd_tar := exec.Command("curl", "-L", synscan_tar_url, "-o", "package.tar.gz")
	err := synscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	synscan_zip_url := "https://digit-labs.org/files/tools/synscan/releases/synscan-5.02.zip"
	synscan_cmd_zip := exec.Command("curl", "-L", synscan_zip_url, "-o", "package.zip")
	err = synscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	synscan_bin_url := "https://digit-labs.org/files/tools/synscan/releases/synscan-5.02.bin"
	synscan_cmd_bin := exec.Command("curl", "-L", synscan_bin_url, "-o", "binary.bin")
	err = synscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	synscan_src_url := "https://digit-labs.org/files/tools/synscan/releases/synscan-5.02.src.tar.gz"
	synscan_cmd_src := exec.Command("curl", "-L", synscan_src_url, "-o", "source.tar.gz")
	err = synscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	synscan_cmd_direct := exec.Command("./binary")
	err = synscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpcap")
	exec.Command("latte", "install", "libpcap").Run()
}
