package main

// Tree - Display directories as trees (with optional color/HTML output)
// Homepage: https://oldmanprogrammer.net/source.php?dir=projects/tree

import (
	"fmt"
	
	"os/exec"
)

func installTree() {
	// Método 1: Descargar y extraer .tar.gz
	tree_tar_url := "https://github.com/Old-Man-Programmer/tree/archive/refs/tags/2.1.3.tar.gz"
	tree_cmd_tar := exec.Command("curl", "-L", tree_tar_url, "-o", "package.tar.gz")
	err := tree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tree_zip_url := "https://github.com/Old-Man-Programmer/tree/archive/refs/tags/2.1.3.zip"
	tree_cmd_zip := exec.Command("curl", "-L", tree_zip_url, "-o", "package.zip")
	err = tree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tree_bin_url := "https://github.com/Old-Man-Programmer/tree/archive/refs/tags/2.1.3.bin"
	tree_cmd_bin := exec.Command("curl", "-L", tree_bin_url, "-o", "binary.bin")
	err = tree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tree_src_url := "https://github.com/Old-Man-Programmer/tree/archive/refs/tags/2.1.3.src.tar.gz"
	tree_cmd_src := exec.Command("curl", "-L", tree_src_url, "-o", "source.tar.gz")
	err = tree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tree_cmd_direct := exec.Command("./binary")
	err = tree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
