package main

// Grepip - Filters IPv4 & IPv6 addresses with a grep-compatible interface
// Homepage: https://ipinfo.io

import (
	"fmt"
	
	"os/exec"
)

func installGrepip() {
	// Método 1: Descargar y extraer .tar.gz
	grepip_tar_url := "https://github.com/ipinfo/cli/archive/refs/tags/grepip-1.2.2.tar.gz"
	grepip_cmd_tar := exec.Command("curl", "-L", grepip_tar_url, "-o", "package.tar.gz")
	err := grepip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grepip_zip_url := "https://github.com/ipinfo/cli/archive/refs/tags/grepip-1.2.2.zip"
	grepip_cmd_zip := exec.Command("curl", "-L", grepip_zip_url, "-o", "package.zip")
	err = grepip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grepip_bin_url := "https://github.com/ipinfo/cli/archive/refs/tags/grepip-1.2.2.bin"
	grepip_cmd_bin := exec.Command("curl", "-L", grepip_bin_url, "-o", "binary.bin")
	err = grepip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grepip_src_url := "https://github.com/ipinfo/cli/archive/refs/tags/grepip-1.2.2.src.tar.gz"
	grepip_cmd_src := exec.Command("curl", "-L", grepip_src_url, "-o", "source.tar.gz")
	err = grepip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grepip_cmd_direct := exec.Command("./binary")
	err = grepip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
