package main

// Janet - Dynamic language and bytecode vm
// Homepage: https://janet-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installJanet() {
	// Método 1: Descargar y extraer .tar.gz
	janet_tar_url := "https://github.com/janet-lang/janet/archive/refs/tags/v1.35.2.tar.gz"
	janet_cmd_tar := exec.Command("curl", "-L", janet_tar_url, "-o", "package.tar.gz")
	err := janet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	janet_zip_url := "https://github.com/janet-lang/janet/archive/refs/tags/v1.35.2.zip"
	janet_cmd_zip := exec.Command("curl", "-L", janet_zip_url, "-o", "package.zip")
	err = janet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	janet_bin_url := "https://github.com/janet-lang/janet/archive/refs/tags/v1.35.2.bin"
	janet_cmd_bin := exec.Command("curl", "-L", janet_bin_url, "-o", "binary.bin")
	err = janet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	janet_src_url := "https://github.com/janet-lang/janet/archive/refs/tags/v1.35.2.src.tar.gz"
	janet_cmd_src := exec.Command("curl", "-L", janet_src_url, "-o", "source.tar.gz")
	err = janet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	janet_cmd_direct := exec.Command("./binary")
	err = janet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
