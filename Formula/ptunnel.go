package main

// Ptunnel - Tunnel over ICMP
// Homepage: https://www.cs.uit.no/~daniels/PingTunnel/

import (
	"fmt"
	
	"os/exec"
)

func installPtunnel() {
	// Método 1: Descargar y extraer .tar.gz
	ptunnel_tar_url := "https://www.cs.uit.no/~daniels/PingTunnel/PingTunnel-0.72.tar.gz"
	ptunnel_cmd_tar := exec.Command("curl", "-L", ptunnel_tar_url, "-o", "package.tar.gz")
	err := ptunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ptunnel_zip_url := "https://www.cs.uit.no/~daniels/PingTunnel/PingTunnel-0.72.zip"
	ptunnel_cmd_zip := exec.Command("curl", "-L", ptunnel_zip_url, "-o", "package.zip")
	err = ptunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ptunnel_bin_url := "https://www.cs.uit.no/~daniels/PingTunnel/PingTunnel-0.72.bin"
	ptunnel_cmd_bin := exec.Command("curl", "-L", ptunnel_bin_url, "-o", "binary.bin")
	err = ptunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ptunnel_src_url := "https://www.cs.uit.no/~daniels/PingTunnel/PingTunnel-0.72.src.tar.gz"
	ptunnel_cmd_src := exec.Command("curl", "-L", ptunnel_src_url, "-o", "source.tar.gz")
	err = ptunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ptunnel_cmd_direct := exec.Command("./binary")
	err = ptunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
