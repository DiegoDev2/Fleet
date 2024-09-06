package main

// Coder - Tool for provisioning self-hosted development environments with Terraform
// Homepage: https://coder.com

import (
	"fmt"
	
	"os/exec"
)

func installCoder() {
	// Método 1: Descargar y extraer .tar.gz
	coder_tar_url := "https://github.com/coder/coder/archive/refs/tags/v2.14.3.tar.gz"
	coder_cmd_tar := exec.Command("curl", "-L", coder_tar_url, "-o", "package.tar.gz")
	err := coder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coder_zip_url := "https://github.com/coder/coder/archive/refs/tags/v2.14.3.zip"
	coder_cmd_zip := exec.Command("curl", "-L", coder_zip_url, "-o", "package.zip")
	err = coder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coder_bin_url := "https://github.com/coder/coder/archive/refs/tags/v2.14.3.bin"
	coder_cmd_bin := exec.Command("curl", "-L", coder_bin_url, "-o", "binary.bin")
	err = coder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coder_src_url := "https://github.com/coder/coder/archive/refs/tags/v2.14.3.src.tar.gz"
	coder_cmd_src := exec.Command("curl", "-L", coder_src_url, "-o", "source.tar.gz")
	err = coder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coder_cmd_direct := exec.Command("./binary")
	err = coder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
