package main

// Tm - TriggerMesh CLI to work with knative objects
// Homepage: https://triggermesh.com

import (
	"fmt"
	
	"os/exec"
)

func installTm() {
	// Método 1: Descargar y extraer .tar.gz
	tm_tar_url := "https://github.com/triggermesh/tm/archive/refs/tags/v1.21.0.tar.gz"
	tm_cmd_tar := exec.Command("curl", "-L", tm_tar_url, "-o", "package.tar.gz")
	err := tm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tm_zip_url := "https://github.com/triggermesh/tm/archive/refs/tags/v1.21.0.zip"
	tm_cmd_zip := exec.Command("curl", "-L", tm_zip_url, "-o", "package.zip")
	err = tm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tm_bin_url := "https://github.com/triggermesh/tm/archive/refs/tags/v1.21.0.bin"
	tm_cmd_bin := exec.Command("curl", "-L", tm_bin_url, "-o", "binary.bin")
	err = tm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tm_src_url := "https://github.com/triggermesh/tm/archive/refs/tags/v1.21.0.src.tar.gz"
	tm_cmd_src := exec.Command("curl", "-L", tm_src_url, "-o", "source.tar.gz")
	err = tm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tm_cmd_direct := exec.Command("./binary")
	err = tm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
