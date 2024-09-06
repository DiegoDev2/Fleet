package main

// Qbe - Compiler Backend
// Homepage: https://c9x.me/compile/

import (
	"fmt"
	
	"os/exec"
)

func installQbe() {
	// Método 1: Descargar y extraer .tar.gz
	qbe_tar_url := "https://c9x.me/compile/release/qbe-1.2.tar.xz"
	qbe_cmd_tar := exec.Command("curl", "-L", qbe_tar_url, "-o", "package.tar.gz")
	err := qbe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qbe_zip_url := "https://c9x.me/compile/release/qbe-1.2.tar.xz"
	qbe_cmd_zip := exec.Command("curl", "-L", qbe_zip_url, "-o", "package.zip")
	err = qbe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qbe_bin_url := "https://c9x.me/compile/release/qbe-1.2.tar.xz"
	qbe_cmd_bin := exec.Command("curl", "-L", qbe_bin_url, "-o", "binary.bin")
	err = qbe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qbe_src_url := "https://c9x.me/compile/release/qbe-1.2.tar.xz"
	qbe_cmd_src := exec.Command("curl", "-L", qbe_src_url, "-o", "source.tar.gz")
	err = qbe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qbe_cmd_direct := exec.Command("./binary")
	err = qbe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
