package main

// AsTree - Print a list of paths as a tree of paths
// Homepage: https://github.com/jez/as-tree

import (
	"fmt"
	
	"os/exec"
)

func installAsTree() {
	// Método 1: Descargar y extraer .tar.gz
	astree_tar_url := "https://github.com/jez/as-tree/archive/refs/tags/0.12.0.tar.gz"
	astree_cmd_tar := exec.Command("curl", "-L", astree_tar_url, "-o", "package.tar.gz")
	err := astree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	astree_zip_url := "https://github.com/jez/as-tree/archive/refs/tags/0.12.0.zip"
	astree_cmd_zip := exec.Command("curl", "-L", astree_zip_url, "-o", "package.zip")
	err = astree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	astree_bin_url := "https://github.com/jez/as-tree/archive/refs/tags/0.12.0.bin"
	astree_cmd_bin := exec.Command("curl", "-L", astree_bin_url, "-o", "binary.bin")
	err = astree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	astree_src_url := "https://github.com/jez/as-tree/archive/refs/tags/0.12.0.src.tar.gz"
	astree_cmd_src := exec.Command("curl", "-L", astree_src_url, "-o", "source.tar.gz")
	err = astree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	astree_cmd_direct := exec.Command("./binary")
	err = astree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
