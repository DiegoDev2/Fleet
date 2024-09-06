package main

// Okteto - Build better apps by developing and testing code directly in Kubernetes
// Homepage: https://okteto.com

import (
	"fmt"
	
	"os/exec"
)

func installOkteto() {
	// Método 1: Descargar y extraer .tar.gz
	okteto_tar_url := "https://github.com/okteto/okteto/archive/refs/tags/2.31.0.tar.gz"
	okteto_cmd_tar := exec.Command("curl", "-L", okteto_tar_url, "-o", "package.tar.gz")
	err := okteto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	okteto_zip_url := "https://github.com/okteto/okteto/archive/refs/tags/2.31.0.zip"
	okteto_cmd_zip := exec.Command("curl", "-L", okteto_zip_url, "-o", "package.zip")
	err = okteto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	okteto_bin_url := "https://github.com/okteto/okteto/archive/refs/tags/2.31.0.bin"
	okteto_cmd_bin := exec.Command("curl", "-L", okteto_bin_url, "-o", "binary.bin")
	err = okteto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	okteto_src_url := "https://github.com/okteto/okteto/archive/refs/tags/2.31.0.src.tar.gz"
	okteto_cmd_src := exec.Command("curl", "-L", okteto_src_url, "-o", "source.tar.gz")
	err = okteto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	okteto_cmd_direct := exec.Command("./binary")
	err = okteto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
