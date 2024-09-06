package main

// Wren - Small, fast, class-based concurrent scripting language
// Homepage: https://wren.io

import (
	"fmt"
	
	"os/exec"
)

func installWren() {
	// Método 1: Descargar y extraer .tar.gz
	wren_tar_url := "https://github.com/wren-lang/wren/archive/refs/tags/0.4.0.tar.gz"
	wren_cmd_tar := exec.Command("curl", "-L", wren_tar_url, "-o", "package.tar.gz")
	err := wren_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wren_zip_url := "https://github.com/wren-lang/wren/archive/refs/tags/0.4.0.zip"
	wren_cmd_zip := exec.Command("curl", "-L", wren_zip_url, "-o", "package.zip")
	err = wren_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wren_bin_url := "https://github.com/wren-lang/wren/archive/refs/tags/0.4.0.bin"
	wren_cmd_bin := exec.Command("curl", "-L", wren_bin_url, "-o", "binary.bin")
	err = wren_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wren_src_url := "https://github.com/wren-lang/wren/archive/refs/tags/0.4.0.src.tar.gz"
	wren_cmd_src := exec.Command("curl", "-L", wren_src_url, "-o", "source.tar.gz")
	err = wren_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wren_cmd_direct := exec.Command("./binary")
	err = wren_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
