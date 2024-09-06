package main

// Sshs - Graphical command-line client for SSH
// Homepage: https://github.com/quantumsheep/sshs

import (
	"fmt"
	
	"os/exec"
)

func installSshs() {
	// Método 1: Descargar y extraer .tar.gz
	sshs_tar_url := "https://github.com/quantumsheep/sshs/archive/refs/tags/4.5.1.tar.gz"
	sshs_cmd_tar := exec.Command("curl", "-L", sshs_tar_url, "-o", "package.tar.gz")
	err := sshs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshs_zip_url := "https://github.com/quantumsheep/sshs/archive/refs/tags/4.5.1.zip"
	sshs_cmd_zip := exec.Command("curl", "-L", sshs_zip_url, "-o", "package.zip")
	err = sshs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshs_bin_url := "https://github.com/quantumsheep/sshs/archive/refs/tags/4.5.1.bin"
	sshs_cmd_bin := exec.Command("curl", "-L", sshs_bin_url, "-o", "binary.bin")
	err = sshs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshs_src_url := "https://github.com/quantumsheep/sshs/archive/refs/tags/4.5.1.src.tar.gz"
	sshs_cmd_src := exec.Command("curl", "-L", sshs_src_url, "-o", "source.tar.gz")
	err = sshs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshs_cmd_direct := exec.Command("./binary")
	err = sshs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
