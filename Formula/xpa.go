package main

// Xpa - Seamless communication between Unix programs
// Homepage: https://hea-www.harvard.edu/RD/xpa/

import (
	"fmt"
	
	"os/exec"
)

func installXpa() {
	// Método 1: Descargar y extraer .tar.gz
	xpa_tar_url := "https://github.com/ericmandel/xpa/archive/refs/tags/2.1.20.tar.gz"
	xpa_cmd_tar := exec.Command("curl", "-L", xpa_tar_url, "-o", "package.tar.gz")
	err := xpa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xpa_zip_url := "https://github.com/ericmandel/xpa/archive/refs/tags/2.1.20.zip"
	xpa_cmd_zip := exec.Command("curl", "-L", xpa_zip_url, "-o", "package.zip")
	err = xpa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xpa_bin_url := "https://github.com/ericmandel/xpa/archive/refs/tags/2.1.20.bin"
	xpa_cmd_bin := exec.Command("curl", "-L", xpa_bin_url, "-o", "binary.bin")
	err = xpa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xpa_src_url := "https://github.com/ericmandel/xpa/archive/refs/tags/2.1.20.src.tar.gz"
	xpa_cmd_src := exec.Command("curl", "-L", xpa_src_url, "-o", "source.tar.gz")
	err = xpa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xpa_cmd_direct := exec.Command("./binary")
	err = xpa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
}
