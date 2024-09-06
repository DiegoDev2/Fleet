package main

// Qrcp - Transfer files to and from your computer by scanning a QR code
// Homepage: https://qrcp.sh

import (
	"fmt"
	
	"os/exec"
)

func installQrcp() {
	// Método 1: Descargar y extraer .tar.gz
	qrcp_tar_url := "https://github.com/claudiodangelis/qrcp/archive/refs/tags/0.11.3.tar.gz"
	qrcp_cmd_tar := exec.Command("curl", "-L", qrcp_tar_url, "-o", "package.tar.gz")
	err := qrcp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qrcp_zip_url := "https://github.com/claudiodangelis/qrcp/archive/refs/tags/0.11.3.zip"
	qrcp_cmd_zip := exec.Command("curl", "-L", qrcp_zip_url, "-o", "package.zip")
	err = qrcp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qrcp_bin_url := "https://github.com/claudiodangelis/qrcp/archive/refs/tags/0.11.3.bin"
	qrcp_cmd_bin := exec.Command("curl", "-L", qrcp_bin_url, "-o", "binary.bin")
	err = qrcp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qrcp_src_url := "https://github.com/claudiodangelis/qrcp/archive/refs/tags/0.11.3.src.tar.gz"
	qrcp_cmd_src := exec.Command("curl", "-L", qrcp_src_url, "-o", "source.tar.gz")
	err = qrcp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qrcp_cmd_direct := exec.Command("./binary")
	err = qrcp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
