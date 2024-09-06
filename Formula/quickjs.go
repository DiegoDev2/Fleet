package main

// Quickjs - Small and embeddable JavaScript engine
// Homepage: https://bellard.org/quickjs/

import (
	"fmt"
	
	"os/exec"
)

func installQuickjs() {
	// Método 1: Descargar y extraer .tar.gz
	quickjs_tar_url := "https://bellard.org/quickjs/quickjs-2024-01-13.tar.xz"
	quickjs_cmd_tar := exec.Command("curl", "-L", quickjs_tar_url, "-o", "package.tar.gz")
	err := quickjs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quickjs_zip_url := "https://bellard.org/quickjs/quickjs-2024-01-13.tar.xz"
	quickjs_cmd_zip := exec.Command("curl", "-L", quickjs_zip_url, "-o", "package.zip")
	err = quickjs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quickjs_bin_url := "https://bellard.org/quickjs/quickjs-2024-01-13.tar.xz"
	quickjs_cmd_bin := exec.Command("curl", "-L", quickjs_bin_url, "-o", "binary.bin")
	err = quickjs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quickjs_src_url := "https://bellard.org/quickjs/quickjs-2024-01-13.tar.xz"
	quickjs_cmd_src := exec.Command("curl", "-L", quickjs_src_url, "-o", "source.tar.gz")
	err = quickjs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quickjs_cmd_direct := exec.Command("./binary")
	err = quickjs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
