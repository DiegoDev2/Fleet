package main

// Gitleaks - Audit git repos for secrets
// Homepage: https://github.com/gitleaks/gitleaks

import (
	"fmt"
	
	"os/exec"
)

func installGitleaks() {
	// Método 1: Descargar y extraer .tar.gz
	gitleaks_tar_url := "https://github.com/gitleaks/gitleaks/archive/refs/tags/v8.18.4.tar.gz"
	gitleaks_cmd_tar := exec.Command("curl", "-L", gitleaks_tar_url, "-o", "package.tar.gz")
	err := gitleaks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitleaks_zip_url := "https://github.com/gitleaks/gitleaks/archive/refs/tags/v8.18.4.zip"
	gitleaks_cmd_zip := exec.Command("curl", "-L", gitleaks_zip_url, "-o", "package.zip")
	err = gitleaks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitleaks_bin_url := "https://github.com/gitleaks/gitleaks/archive/refs/tags/v8.18.4.bin"
	gitleaks_cmd_bin := exec.Command("curl", "-L", gitleaks_bin_url, "-o", "binary.bin")
	err = gitleaks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitleaks_src_url := "https://github.com/gitleaks/gitleaks/archive/refs/tags/v8.18.4.src.tar.gz"
	gitleaks_cmd_src := exec.Command("curl", "-L", gitleaks_src_url, "-o", "source.tar.gz")
	err = gitleaks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitleaks_cmd_direct := exec.Command("./binary")
	err = gitleaks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
