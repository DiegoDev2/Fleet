package main

// Corepack - Package acting as bridge between Node projects and their package managers
// Homepage: https://github.com/nodejs/corepack

import (
	"fmt"
	
	"os/exec"
)

func installCorepack() {
	// Método 1: Descargar y extraer .tar.gz
	corepack_tar_url := "https://registry.npmjs.org/corepack/-/corepack-0.29.3.tgz"
	corepack_cmd_tar := exec.Command("curl", "-L", corepack_tar_url, "-o", "package.tar.gz")
	err := corepack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	corepack_zip_url := "https://registry.npmjs.org/corepack/-/corepack-0.29.3.tgz"
	corepack_cmd_zip := exec.Command("curl", "-L", corepack_zip_url, "-o", "package.zip")
	err = corepack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	corepack_bin_url := "https://registry.npmjs.org/corepack/-/corepack-0.29.3.tgz"
	corepack_cmd_bin := exec.Command("curl", "-L", corepack_bin_url, "-o", "binary.bin")
	err = corepack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	corepack_src_url := "https://registry.npmjs.org/corepack/-/corepack-0.29.3.tgz"
	corepack_cmd_src := exec.Command("curl", "-L", corepack_src_url, "-o", "source.tar.gz")
	err = corepack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	corepack_cmd_direct := exec.Command("./binary")
	err = corepack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
