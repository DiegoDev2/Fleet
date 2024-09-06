package main

// Metalang99 - C99 preprocessor-based metaprogramming language
// Homepage: https://github.com/Hirrolot/metalang99

import (
	"fmt"
	
	"os/exec"
)

func installMetalang99() {
	// Método 1: Descargar y extraer .tar.gz
	metalang99_tar_url := "https://github.com/Hirrolot/metalang99/archive/refs/tags/v1.13.3.tar.gz"
	metalang99_cmd_tar := exec.Command("curl", "-L", metalang99_tar_url, "-o", "package.tar.gz")
	err := metalang99_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	metalang99_zip_url := "https://github.com/Hirrolot/metalang99/archive/refs/tags/v1.13.3.zip"
	metalang99_cmd_zip := exec.Command("curl", "-L", metalang99_zip_url, "-o", "package.zip")
	err = metalang99_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	metalang99_bin_url := "https://github.com/Hirrolot/metalang99/archive/refs/tags/v1.13.3.bin"
	metalang99_cmd_bin := exec.Command("curl", "-L", metalang99_bin_url, "-o", "binary.bin")
	err = metalang99_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	metalang99_src_url := "https://github.com/Hirrolot/metalang99/archive/refs/tags/v1.13.3.src.tar.gz"
	metalang99_cmd_src := exec.Command("curl", "-L", metalang99_src_url, "-o", "source.tar.gz")
	err = metalang99_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	metalang99_cmd_direct := exec.Command("./binary")
	err = metalang99_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
