package main

// Basedpyright - Pyright fork with various improvements and built-in pylance features
// Homepage: https://github.com/DetachHead/basedpyright

import (
	"fmt"
	
	"os/exec"
)

func installBasedpyright() {
	// Método 1: Descargar y extraer .tar.gz
	basedpyright_tar_url := "https://registry.npmjs.org/basedpyright/-/basedpyright-1.17.2.tgz"
	basedpyright_cmd_tar := exec.Command("curl", "-L", basedpyright_tar_url, "-o", "package.tar.gz")
	err := basedpyright_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	basedpyright_zip_url := "https://registry.npmjs.org/basedpyright/-/basedpyright-1.17.2.tgz"
	basedpyright_cmd_zip := exec.Command("curl", "-L", basedpyright_zip_url, "-o", "package.zip")
	err = basedpyright_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	basedpyright_bin_url := "https://registry.npmjs.org/basedpyright/-/basedpyright-1.17.2.tgz"
	basedpyright_cmd_bin := exec.Command("curl", "-L", basedpyright_bin_url, "-o", "binary.bin")
	err = basedpyright_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	basedpyright_src_url := "https://registry.npmjs.org/basedpyright/-/basedpyright-1.17.2.tgz"
	basedpyright_cmd_src := exec.Command("curl", "-L", basedpyright_src_url, "-o", "source.tar.gz")
	err = basedpyright_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	basedpyright_cmd_direct := exec.Command("./binary")
	err = basedpyright_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
