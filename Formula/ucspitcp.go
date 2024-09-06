package main

// UcspiTcp - Tools for building TCP client-server applications
// Homepage: https://cr.yp.to/ucspi-tcp.html

import (
	"fmt"
	
	"os/exec"
)

func installUcspiTcp() {
	// Método 1: Descargar y extraer .tar.gz
	ucspitcp_tar_url := "https://cr.yp.to/ucspi-tcp/ucspi-tcp-0.88.tar.gz"
	ucspitcp_cmd_tar := exec.Command("curl", "-L", ucspitcp_tar_url, "-o", "package.tar.gz")
	err := ucspitcp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ucspitcp_zip_url := "https://cr.yp.to/ucspi-tcp/ucspi-tcp-0.88.zip"
	ucspitcp_cmd_zip := exec.Command("curl", "-L", ucspitcp_zip_url, "-o", "package.zip")
	err = ucspitcp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ucspitcp_bin_url := "https://cr.yp.to/ucspi-tcp/ucspi-tcp-0.88.bin"
	ucspitcp_cmd_bin := exec.Command("curl", "-L", ucspitcp_bin_url, "-o", "binary.bin")
	err = ucspitcp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ucspitcp_src_url := "https://cr.yp.to/ucspi-tcp/ucspi-tcp-0.88.src.tar.gz"
	ucspitcp_cmd_src := exec.Command("curl", "-L", ucspitcp_src_url, "-o", "source.tar.gz")
	err = ucspitcp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ucspitcp_cmd_direct := exec.Command("./binary")
	err = ucspitcp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
