package main

// Levant - Templating and deployment tool for HashiCorp Nomad jobs
// Homepage: https://github.com/hashicorp/levant

import (
	"fmt"
	
	"os/exec"
)

func installLevant() {
	// Método 1: Descargar y extraer .tar.gz
	levant_tar_url := "https://github.com/hashicorp/levant/archive/refs/tags/v0.3.3.tar.gz"
	levant_cmd_tar := exec.Command("curl", "-L", levant_tar_url, "-o", "package.tar.gz")
	err := levant_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	levant_zip_url := "https://github.com/hashicorp/levant/archive/refs/tags/v0.3.3.zip"
	levant_cmd_zip := exec.Command("curl", "-L", levant_zip_url, "-o", "package.zip")
	err = levant_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	levant_bin_url := "https://github.com/hashicorp/levant/archive/refs/tags/v0.3.3.bin"
	levant_cmd_bin := exec.Command("curl", "-L", levant_bin_url, "-o", "binary.bin")
	err = levant_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	levant_src_url := "https://github.com/hashicorp/levant/archive/refs/tags/v0.3.3.src.tar.gz"
	levant_cmd_src := exec.Command("curl", "-L", levant_src_url, "-o", "source.tar.gz")
	err = levant_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	levant_cmd_direct := exec.Command("./binary")
	err = levant_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
