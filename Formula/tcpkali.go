package main

// Tcpkali - High performance TCP and WebSocket load generator and sink
// Homepage: https://github.com/satori-com/tcpkali

import (
	"fmt"
	
	"os/exec"
)

func installTcpkali() {
	// Método 1: Descargar y extraer .tar.gz
	tcpkali_tar_url := "https://github.com/satori-com/tcpkali/releases/download/v1.1.1/tcpkali-1.1.1.tar.gz"
	tcpkali_cmd_tar := exec.Command("curl", "-L", tcpkali_tar_url, "-o", "package.tar.gz")
	err := tcpkali_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcpkali_zip_url := "https://github.com/satori-com/tcpkali/releases/download/v1.1.1/tcpkali-1.1.1.zip"
	tcpkali_cmd_zip := exec.Command("curl", "-L", tcpkali_zip_url, "-o", "package.zip")
	err = tcpkali_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcpkali_bin_url := "https://github.com/satori-com/tcpkali/releases/download/v1.1.1/tcpkali-1.1.1.bin"
	tcpkali_cmd_bin := exec.Command("curl", "-L", tcpkali_bin_url, "-o", "binary.bin")
	err = tcpkali_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcpkali_src_url := "https://github.com/satori-com/tcpkali/releases/download/v1.1.1/tcpkali-1.1.1.src.tar.gz"
	tcpkali_cmd_src := exec.Command("curl", "-L", tcpkali_src_url, "-o", "source.tar.gz")
	err = tcpkali_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcpkali_cmd_direct := exec.Command("./binary")
	err = tcpkali_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
