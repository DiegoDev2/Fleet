package main

// Rollup - Next-generation ES module bundler
// Homepage: https://rollupjs.org/

import (
	"fmt"
	
	"os/exec"
)

func installRollup() {
	// Método 1: Descargar y extraer .tar.gz
	rollup_tar_url := "https://registry.npmjs.org/rollup/-/rollup-4.21.2.tgz"
	rollup_cmd_tar := exec.Command("curl", "-L", rollup_tar_url, "-o", "package.tar.gz")
	err := rollup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rollup_zip_url := "https://registry.npmjs.org/rollup/-/rollup-4.21.2.tgz"
	rollup_cmd_zip := exec.Command("curl", "-L", rollup_zip_url, "-o", "package.zip")
	err = rollup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rollup_bin_url := "https://registry.npmjs.org/rollup/-/rollup-4.21.2.tgz"
	rollup_cmd_bin := exec.Command("curl", "-L", rollup_bin_url, "-o", "binary.bin")
	err = rollup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rollup_src_url := "https://registry.npmjs.org/rollup/-/rollup-4.21.2.tgz"
	rollup_cmd_src := exec.Command("curl", "-L", rollup_src_url, "-o", "source.tar.gz")
	err = rollup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rollup_cmd_direct := exec.Command("./binary")
	err = rollup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
