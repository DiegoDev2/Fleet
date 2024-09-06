package main

// Pstree - Show ps output as a tree
// Homepage: https://github.com/FredHucht/pstree

import (
	"fmt"
	
	"os/exec"
)

func installPstree() {
	// Método 1: Descargar y extraer .tar.gz
	pstree_tar_url := "https://github.com/FredHucht/pstree/archive/refs/tags/v2.40.tar.gz"
	pstree_cmd_tar := exec.Command("curl", "-L", pstree_tar_url, "-o", "package.tar.gz")
	err := pstree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pstree_zip_url := "https://github.com/FredHucht/pstree/archive/refs/tags/v2.40.zip"
	pstree_cmd_zip := exec.Command("curl", "-L", pstree_zip_url, "-o", "package.zip")
	err = pstree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pstree_bin_url := "https://github.com/FredHucht/pstree/archive/refs/tags/v2.40.bin"
	pstree_cmd_bin := exec.Command("curl", "-L", pstree_bin_url, "-o", "binary.bin")
	err = pstree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pstree_src_url := "https://github.com/FredHucht/pstree/archive/refs/tags/v2.40.src.tar.gz"
	pstree_cmd_src := exec.Command("curl", "-L", pstree_src_url, "-o", "source.tar.gz")
	err = pstree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pstree_cmd_direct := exec.Command("./binary")
	err = pstree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
