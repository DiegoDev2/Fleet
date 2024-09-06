package main

// Czg - Interactive Commitizen CLI that generate standardized commit messages
// Homepage: https://github.com/Zhengqbbb/cz-git

import (
	"fmt"
	
	"os/exec"
)

func installCzg() {
	// Método 1: Descargar y extraer .tar.gz
	czg_tar_url := "https://registry.npmjs.org/czg/-/czg-1.9.4.tgz"
	czg_cmd_tar := exec.Command("curl", "-L", czg_tar_url, "-o", "package.tar.gz")
	err := czg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	czg_zip_url := "https://registry.npmjs.org/czg/-/czg-1.9.4.tgz"
	czg_cmd_zip := exec.Command("curl", "-L", czg_zip_url, "-o", "package.zip")
	err = czg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	czg_bin_url := "https://registry.npmjs.org/czg/-/czg-1.9.4.tgz"
	czg_cmd_bin := exec.Command("curl", "-L", czg_bin_url, "-o", "binary.bin")
	err = czg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	czg_src_url := "https://registry.npmjs.org/czg/-/czg-1.9.4.tgz"
	czg_cmd_src := exec.Command("curl", "-L", czg_src_url, "-o", "source.tar.gz")
	err = czg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	czg_cmd_direct := exec.Command("./binary")
	err = czg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
