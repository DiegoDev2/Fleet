package main

// Erdtree - Multi-threaded file-tree visualizer and disk usage analyzer
// Homepage: https://github.com/solidiquis/erdtree

import (
	"fmt"
	
	"os/exec"
)

func installErdtree() {
	// Método 1: Descargar y extraer .tar.gz
	erdtree_tar_url := "https://github.com/solidiquis/erdtree/archive/refs/tags/v3.1.2.tar.gz"
	erdtree_cmd_tar := exec.Command("curl", "-L", erdtree_tar_url, "-o", "package.tar.gz")
	err := erdtree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	erdtree_zip_url := "https://github.com/solidiquis/erdtree/archive/refs/tags/v3.1.2.zip"
	erdtree_cmd_zip := exec.Command("curl", "-L", erdtree_zip_url, "-o", "package.zip")
	err = erdtree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	erdtree_bin_url := "https://github.com/solidiquis/erdtree/archive/refs/tags/v3.1.2.bin"
	erdtree_cmd_bin := exec.Command("curl", "-L", erdtree_bin_url, "-o", "binary.bin")
	err = erdtree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	erdtree_src_url := "https://github.com/solidiquis/erdtree/archive/refs/tags/v3.1.2.src.tar.gz"
	erdtree_cmd_src := exec.Command("curl", "-L", erdtree_src_url, "-o", "source.tar.gz")
	err = erdtree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	erdtree_cmd_direct := exec.Command("./binary")
	err = erdtree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
