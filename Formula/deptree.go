package main

// DepTree - Tool for visualizing dependencies between files and enforcing dependency rules
// Homepage: https://github.com/gabotechs/dep-tree

import (
	"fmt"
	
	"os/exec"
)

func installDepTree() {
	// Método 1: Descargar y extraer .tar.gz
	deptree_tar_url := "https://github.com/gabotechs/dep-tree/archive/refs/tags/v0.23.0.tar.gz"
	deptree_cmd_tar := exec.Command("curl", "-L", deptree_tar_url, "-o", "package.tar.gz")
	err := deptree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	deptree_zip_url := "https://github.com/gabotechs/dep-tree/archive/refs/tags/v0.23.0.zip"
	deptree_cmd_zip := exec.Command("curl", "-L", deptree_zip_url, "-o", "package.zip")
	err = deptree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	deptree_bin_url := "https://github.com/gabotechs/dep-tree/archive/refs/tags/v0.23.0.bin"
	deptree_cmd_bin := exec.Command("curl", "-L", deptree_bin_url, "-o", "binary.bin")
	err = deptree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	deptree_src_url := "https://github.com/gabotechs/dep-tree/archive/refs/tags/v0.23.0.src.tar.gz"
	deptree_cmd_src := exec.Command("curl", "-L", deptree_src_url, "-o", "source.tar.gz")
	err = deptree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	deptree_cmd_direct := exec.Command("./binary")
	err = deptree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
