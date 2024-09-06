package main

// OsvScanner - Vulnerability scanner which uses the OSV database
// Homepage: https://github.com/google/osv-scanner

import (
	"fmt"
	
	"os/exec"
)

func installOsvScanner() {
	// Método 1: Descargar y extraer .tar.gz
	osvscanner_tar_url := "https://github.com/google/osv-scanner/archive/refs/tags/v1.8.4.tar.gz"
	osvscanner_cmd_tar := exec.Command("curl", "-L", osvscanner_tar_url, "-o", "package.tar.gz")
	err := osvscanner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osvscanner_zip_url := "https://github.com/google/osv-scanner/archive/refs/tags/v1.8.4.zip"
	osvscanner_cmd_zip := exec.Command("curl", "-L", osvscanner_zip_url, "-o", "package.zip")
	err = osvscanner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osvscanner_bin_url := "https://github.com/google/osv-scanner/archive/refs/tags/v1.8.4.bin"
	osvscanner_cmd_bin := exec.Command("curl", "-L", osvscanner_bin_url, "-o", "binary.bin")
	err = osvscanner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osvscanner_src_url := "https://github.com/google/osv-scanner/archive/refs/tags/v1.8.4.src.tar.gz"
	osvscanner_cmd_src := exec.Command("curl", "-L", osvscanner_src_url, "-o", "source.tar.gz")
	err = osvscanner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osvscanner_cmd_direct := exec.Command("./binary")
	err = osvscanner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
