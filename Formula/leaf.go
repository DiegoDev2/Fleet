package main

// Leaf - General purpose reloader for all projects
// Homepage: https://pkg.go.dev/github.com/vrongmeal/leaf

import (
	"fmt"
	
	"os/exec"
)

func installLeaf() {
	// Método 1: Descargar y extraer .tar.gz
	leaf_tar_url := "https://github.com/vrongmeal/leaf/archive/refs/tags/v1.3.0.tar.gz"
	leaf_cmd_tar := exec.Command("curl", "-L", leaf_tar_url, "-o", "package.tar.gz")
	err := leaf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leaf_zip_url := "https://github.com/vrongmeal/leaf/archive/refs/tags/v1.3.0.zip"
	leaf_cmd_zip := exec.Command("curl", "-L", leaf_zip_url, "-o", "package.zip")
	err = leaf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leaf_bin_url := "https://github.com/vrongmeal/leaf/archive/refs/tags/v1.3.0.bin"
	leaf_cmd_bin := exec.Command("curl", "-L", leaf_bin_url, "-o", "binary.bin")
	err = leaf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leaf_src_url := "https://github.com/vrongmeal/leaf/archive/refs/tags/v1.3.0.src.tar.gz"
	leaf_cmd_src := exec.Command("curl", "-L", leaf_src_url, "-o", "source.tar.gz")
	err = leaf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leaf_cmd_direct := exec.Command("./binary")
	err = leaf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
