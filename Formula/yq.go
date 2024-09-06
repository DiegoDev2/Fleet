package main

// Yq - Process YAML, JSON, XML, CSV and properties documents from the CLI
// Homepage: https://github.com/mikefarah/yq

import (
	"fmt"
	
	"os/exec"
)

func installYq() {
	// Método 1: Descargar y extraer .tar.gz
	yq_tar_url := "https://github.com/mikefarah/yq/archive/refs/tags/v4.44.3.tar.gz"
	yq_cmd_tar := exec.Command("curl", "-L", yq_tar_url, "-o", "package.tar.gz")
	err := yq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yq_zip_url := "https://github.com/mikefarah/yq/archive/refs/tags/v4.44.3.zip"
	yq_cmd_zip := exec.Command("curl", "-L", yq_zip_url, "-o", "package.zip")
	err = yq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yq_bin_url := "https://github.com/mikefarah/yq/archive/refs/tags/v4.44.3.bin"
	yq_cmd_bin := exec.Command("curl", "-L", yq_bin_url, "-o", "binary.bin")
	err = yq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yq_src_url := "https://github.com/mikefarah/yq/archive/refs/tags/v4.44.3.src.tar.gz"
	yq_cmd_src := exec.Command("curl", "-L", yq_src_url, "-o", "source.tar.gz")
	err = yq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yq_cmd_direct := exec.Command("./binary")
	err = yq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
}
