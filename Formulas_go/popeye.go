package main

// Popeye - Kubernetes cluster resource sanitizer
// Homepage: https://popeyecli.io

import (
	"fmt"
	
	"os/exec"
)

func installPopeye() {
	// Método 1: Descargar y extraer .tar.gz
	popeye_tar_url := "https://github.com/derailed/popeye/archive/refs/tags/v0.21.3.tar.gz"
	popeye_cmd_tar := exec.Command("curl", "-L", popeye_tar_url, "-o", "package.tar.gz")
	err := popeye_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	popeye_zip_url := "https://github.com/derailed/popeye/archive/refs/tags/v0.21.3.zip"
	popeye_cmd_zip := exec.Command("curl", "-L", popeye_zip_url, "-o", "package.zip")
	err = popeye_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	popeye_bin_url := "https://github.com/derailed/popeye/archive/refs/tags/v0.21.3.bin"
	popeye_cmd_bin := exec.Command("curl", "-L", popeye_bin_url, "-o", "binary.bin")
	err = popeye_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	popeye_src_url := "https://github.com/derailed/popeye/archive/refs/tags/v0.21.3.src.tar.gz"
	popeye_cmd_src := exec.Command("curl", "-L", popeye_src_url, "-o", "source.tar.gz")
	err = popeye_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	popeye_cmd_direct := exec.Command("./binary")
	err = popeye_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
