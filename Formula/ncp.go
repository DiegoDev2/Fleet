package main

// Ncp - File copy tool for LANs
// Homepage: https://www.fefe.de/ncp/

import (
	"fmt"
	
	"os/exec"
)

func installNcp() {
	// Método 1: Descargar y extraer .tar.gz
	ncp_tar_url := "https://dl.fefe.de/ncp-1.2.4.tar.bz2"
	ncp_cmd_tar := exec.Command("curl", "-L", ncp_tar_url, "-o", "package.tar.gz")
	err := ncp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncp_zip_url := "https://dl.fefe.de/ncp-1.2.4.tar.bz2"
	ncp_cmd_zip := exec.Command("curl", "-L", ncp_zip_url, "-o", "package.zip")
	err = ncp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncp_bin_url := "https://dl.fefe.de/ncp-1.2.4.tar.bz2"
	ncp_cmd_bin := exec.Command("curl", "-L", ncp_bin_url, "-o", "binary.bin")
	err = ncp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncp_src_url := "https://dl.fefe.de/ncp-1.2.4.tar.bz2"
	ncp_cmd_src := exec.Command("curl", "-L", ncp_src_url, "-o", "source.tar.gz")
	err = ncp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncp_cmd_direct := exec.Command("./binary")
	err = ncp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libowfat")
	exec.Command("latte", "install", "libowfat").Run()
}
