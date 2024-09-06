package main

// WriteGood - Naive linter for English prose
// Homepage: https://github.com/btford/write-good

import (
	"fmt"
	
	"os/exec"
)

func installWriteGood() {
	// Método 1: Descargar y extraer .tar.gz
	writegood_tar_url := "https://registry.npmjs.org/write-good/-/write-good-1.0.8.tgz"
	writegood_cmd_tar := exec.Command("curl", "-L", writegood_tar_url, "-o", "package.tar.gz")
	err := writegood_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	writegood_zip_url := "https://registry.npmjs.org/write-good/-/write-good-1.0.8.tgz"
	writegood_cmd_zip := exec.Command("curl", "-L", writegood_zip_url, "-o", "package.zip")
	err = writegood_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	writegood_bin_url := "https://registry.npmjs.org/write-good/-/write-good-1.0.8.tgz"
	writegood_cmd_bin := exec.Command("curl", "-L", writegood_bin_url, "-o", "binary.bin")
	err = writegood_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	writegood_src_url := "https://registry.npmjs.org/write-good/-/write-good-1.0.8.tgz"
	writegood_cmd_src := exec.Command("curl", "-L", writegood_src_url, "-o", "source.tar.gz")
	err = writegood_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	writegood_cmd_direct := exec.Command("./binary")
	err = writegood_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
