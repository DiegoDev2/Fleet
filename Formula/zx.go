package main

// Zx - Tool for writing better scripts
// Homepage: https://github.com/google/zx

import (
	"fmt"
	
	"os/exec"
)

func installZx() {
	// Método 1: Descargar y extraer .tar.gz
	zx_tar_url := "https://registry.npmjs.org/zx/-/zx-8.1.5.tgz"
	zx_cmd_tar := exec.Command("curl", "-L", zx_tar_url, "-o", "package.tar.gz")
	err := zx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zx_zip_url := "https://registry.npmjs.org/zx/-/zx-8.1.5.tgz"
	zx_cmd_zip := exec.Command("curl", "-L", zx_zip_url, "-o", "package.zip")
	err = zx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zx_bin_url := "https://registry.npmjs.org/zx/-/zx-8.1.5.tgz"
	zx_cmd_bin := exec.Command("curl", "-L", zx_bin_url, "-o", "binary.bin")
	err = zx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zx_src_url := "https://registry.npmjs.org/zx/-/zx-8.1.5.tgz"
	zx_cmd_src := exec.Command("curl", "-L", zx_src_url, "-o", "source.tar.gz")
	err = zx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zx_cmd_direct := exec.Command("./binary")
	err = zx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
