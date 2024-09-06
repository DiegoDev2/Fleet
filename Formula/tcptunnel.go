package main

// Tcptunnel - TCP port forwarder
// Homepage: https://github.com/vakuum/tcptunnel

import (
	"fmt"
	
	"os/exec"
)

func installTcptunnel() {
	// Método 1: Descargar y extraer .tar.gz
	tcptunnel_tar_url := "https://github.com/vakuum/tcptunnel/archive/refs/tags/v0.8.tar.gz"
	tcptunnel_cmd_tar := exec.Command("curl", "-L", tcptunnel_tar_url, "-o", "package.tar.gz")
	err := tcptunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcptunnel_zip_url := "https://github.com/vakuum/tcptunnel/archive/refs/tags/v0.8.zip"
	tcptunnel_cmd_zip := exec.Command("curl", "-L", tcptunnel_zip_url, "-o", "package.zip")
	err = tcptunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcptunnel_bin_url := "https://github.com/vakuum/tcptunnel/archive/refs/tags/v0.8.bin"
	tcptunnel_cmd_bin := exec.Command("curl", "-L", tcptunnel_bin_url, "-o", "binary.bin")
	err = tcptunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcptunnel_src_url := "https://github.com/vakuum/tcptunnel/archive/refs/tags/v0.8.src.tar.gz"
	tcptunnel_cmd_src := exec.Command("curl", "-L", tcptunnel_src_url, "-o", "source.tar.gz")
	err = tcptunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcptunnel_cmd_direct := exec.Command("./binary")
	err = tcptunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
