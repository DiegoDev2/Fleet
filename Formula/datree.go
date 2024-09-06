package main

// Datree - CLI tool to run policies against Kubernetes manifests YAML files or Helm charts
// Homepage: https://datree.io/

import (
	"fmt"
	
	"os/exec"
)

func installDatree() {
	// Método 1: Descargar y extraer .tar.gz
	datree_tar_url := "https://github.com/datreeio/datree/archive/refs/tags/1.9.19.tar.gz"
	datree_cmd_tar := exec.Command("curl", "-L", datree_tar_url, "-o", "package.tar.gz")
	err := datree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	datree_zip_url := "https://github.com/datreeio/datree/archive/refs/tags/1.9.19.zip"
	datree_cmd_zip := exec.Command("curl", "-L", datree_zip_url, "-o", "package.zip")
	err = datree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	datree_bin_url := "https://github.com/datreeio/datree/archive/refs/tags/1.9.19.bin"
	datree_cmd_bin := exec.Command("curl", "-L", datree_bin_url, "-o", "binary.bin")
	err = datree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	datree_src_url := "https://github.com/datreeio/datree/archive/refs/tags/1.9.19.src.tar.gz"
	datree_cmd_src := exec.Command("curl", "-L", datree_src_url, "-o", "source.tar.gz")
	err = datree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	datree_cmd_direct := exec.Command("./binary")
	err = datree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
