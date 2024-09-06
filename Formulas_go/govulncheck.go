package main

// Govulncheck - Database client and tools for the Go vulnerability database
// Homepage: https://github.com/golang/vuln

import (
	"fmt"
	
	"os/exec"
)

func installGovulncheck() {
	// Método 1: Descargar y extraer .tar.gz
	govulncheck_tar_url := "https://github.com/golang/vuln/archive/refs/tags/v1.1.3.tar.gz"
	govulncheck_cmd_tar := exec.Command("curl", "-L", govulncheck_tar_url, "-o", "package.tar.gz")
	err := govulncheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	govulncheck_zip_url := "https://github.com/golang/vuln/archive/refs/tags/v1.1.3.zip"
	govulncheck_cmd_zip := exec.Command("curl", "-L", govulncheck_zip_url, "-o", "package.zip")
	err = govulncheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	govulncheck_bin_url := "https://github.com/golang/vuln/archive/refs/tags/v1.1.3.bin"
	govulncheck_cmd_bin := exec.Command("curl", "-L", govulncheck_bin_url, "-o", "binary.bin")
	err = govulncheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	govulncheck_src_url := "https://github.com/golang/vuln/archive/refs/tags/v1.1.3.src.tar.gz"
	govulncheck_cmd_src := exec.Command("curl", "-L", govulncheck_src_url, "-o", "source.tar.gz")
	err = govulncheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	govulncheck_cmd_direct := exec.Command("./binary")
	err = govulncheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
