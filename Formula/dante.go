package main

// Dante - SOCKS server and client, implementing RFC 1928 and related standards
// Homepage: https://www.inet.no/dante/

import (
	"fmt"
	
	"os/exec"
)

func installDante() {
	// Método 1: Descargar y extraer .tar.gz
	dante_tar_url := "https://www.inet.no/dante/files/dante-1.4.3.tar.gz"
	dante_cmd_tar := exec.Command("curl", "-L", dante_tar_url, "-o", "package.tar.gz")
	err := dante_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dante_zip_url := "https://www.inet.no/dante/files/dante-1.4.3.zip"
	dante_cmd_zip := exec.Command("curl", "-L", dante_zip_url, "-o", "package.zip")
	err = dante_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dante_bin_url := "https://www.inet.no/dante/files/dante-1.4.3.bin"
	dante_cmd_bin := exec.Command("curl", "-L", dante_bin_url, "-o", "binary.bin")
	err = dante_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dante_src_url := "https://www.inet.no/dante/files/dante-1.4.3.src.tar.gz"
	dante_cmd_src := exec.Command("curl", "-L", dante_src_url, "-o", "source.tar.gz")
	err = dante_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dante_cmd_direct := exec.Command("./binary")
	err = dante_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
