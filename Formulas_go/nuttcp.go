package main

// Nuttcp - Network performance measurement tool
// Homepage: https://www.nuttcp.net/nuttcp/

import (
	"fmt"
	
	"os/exec"
)

func installNuttcp() {
	// Método 1: Descargar y extraer .tar.gz
	nuttcp_tar_url := "https://www.nuttcp.net/nuttcp/nuttcp-8.2.2.tar.bz2"
	nuttcp_cmd_tar := exec.Command("curl", "-L", nuttcp_tar_url, "-o", "package.tar.gz")
	err := nuttcp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nuttcp_zip_url := "https://www.nuttcp.net/nuttcp/nuttcp-8.2.2.tar.bz2"
	nuttcp_cmd_zip := exec.Command("curl", "-L", nuttcp_zip_url, "-o", "package.zip")
	err = nuttcp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nuttcp_bin_url := "https://www.nuttcp.net/nuttcp/nuttcp-8.2.2.tar.bz2"
	nuttcp_cmd_bin := exec.Command("curl", "-L", nuttcp_bin_url, "-o", "binary.bin")
	err = nuttcp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nuttcp_src_url := "https://www.nuttcp.net/nuttcp/nuttcp-8.2.2.tar.bz2"
	nuttcp_cmd_src := exec.Command("curl", "-L", nuttcp_src_url, "-o", "source.tar.gz")
	err = nuttcp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nuttcp_cmd_direct := exec.Command("./binary")
	err = nuttcp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
