package main

// Qp - Command-line (ND)JSON querying
// Homepage: https://github.com/f5io/qp

import (
	"fmt"
	
	"os/exec"
)

func installQp() {
	// Método 1: Descargar y extraer .tar.gz
	qp_tar_url := "https://github.com/f5io/qp/archive/refs/tags/1.0.1.tar.gz"
	qp_cmd_tar := exec.Command("curl", "-L", qp_tar_url, "-o", "package.tar.gz")
	err := qp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qp_zip_url := "https://github.com/f5io/qp/archive/refs/tags/1.0.1.zip"
	qp_cmd_zip := exec.Command("curl", "-L", qp_zip_url, "-o", "package.zip")
	err = qp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qp_bin_url := "https://github.com/f5io/qp/archive/refs/tags/1.0.1.bin"
	qp_cmd_bin := exec.Command("curl", "-L", qp_bin_url, "-o", "binary.bin")
	err = qp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qp_src_url := "https://github.com/f5io/qp/archive/refs/tags/1.0.1.src.tar.gz"
	qp_cmd_src := exec.Command("curl", "-L", qp_src_url, "-o", "source.tar.gz")
	err = qp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qp_cmd_direct := exec.Command("./binary")
	err = qp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: quickjs")
	exec.Command("latte", "install", "quickjs").Run()
}
