package main

// Zlint - X.509 Certificate Linter focused on Web PKI standards and requirements
// Homepage: https://github.com/zmap/zlint

import (
	"fmt"
	
	"os/exec"
)

func installZlint() {
	// Método 1: Descargar y extraer .tar.gz
	zlint_tar_url := "https://github.com/zmap/zlint/archive/refs/tags/v3.6.3.tar.gz"
	zlint_cmd_tar := exec.Command("curl", "-L", zlint_tar_url, "-o", "package.tar.gz")
	err := zlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zlint_zip_url := "https://github.com/zmap/zlint/archive/refs/tags/v3.6.3.zip"
	zlint_cmd_zip := exec.Command("curl", "-L", zlint_zip_url, "-o", "package.zip")
	err = zlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zlint_bin_url := "https://github.com/zmap/zlint/archive/refs/tags/v3.6.3.bin"
	zlint_cmd_bin := exec.Command("curl", "-L", zlint_bin_url, "-o", "binary.bin")
	err = zlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zlint_src_url := "https://github.com/zmap/zlint/archive/refs/tags/v3.6.3.src.tar.gz"
	zlint_cmd_src := exec.Command("curl", "-L", zlint_src_url, "-o", "source.tar.gz")
	err = zlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zlint_cmd_direct := exec.Command("./binary")
	err = zlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
