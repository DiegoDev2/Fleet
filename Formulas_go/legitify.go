package main

// Legitify - Tool to detect/remediate misconfig and security risks of GitHub/GitLab assets
// Homepage: https://legitify.dev/

import (
	"fmt"
	
	"os/exec"
)

func installLegitify() {
	// Método 1: Descargar y extraer .tar.gz
	legitify_tar_url := "https://github.com/Legit-Labs/legitify/archive/refs/tags/v1.0.11.tar.gz"
	legitify_cmd_tar := exec.Command("curl", "-L", legitify_tar_url, "-o", "package.tar.gz")
	err := legitify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	legitify_zip_url := "https://github.com/Legit-Labs/legitify/archive/refs/tags/v1.0.11.zip"
	legitify_cmd_zip := exec.Command("curl", "-L", legitify_zip_url, "-o", "package.zip")
	err = legitify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	legitify_bin_url := "https://github.com/Legit-Labs/legitify/archive/refs/tags/v1.0.11.bin"
	legitify_cmd_bin := exec.Command("curl", "-L", legitify_bin_url, "-o", "binary.bin")
	err = legitify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	legitify_src_url := "https://github.com/Legit-Labs/legitify/archive/refs/tags/v1.0.11.src.tar.gz"
	legitify_cmd_src := exec.Command("curl", "-L", legitify_src_url, "-o", "source.tar.gz")
	err = legitify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	legitify_cmd_direct := exec.Command("./binary")
	err = legitify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
