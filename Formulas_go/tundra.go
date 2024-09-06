package main

// Tundra - Code build system that tries to be fast for incremental builds
// Homepage: https://github.com/deplinenoise/tundra

import (
	"fmt"
	
	"os/exec"
)

func installTundra() {
	// Método 1: Descargar y extraer .tar.gz
	tundra_tar_url := "https://github.com/deplinenoise/tundra/archive/refs/tags/v2.17.1.tar.gz"
	tundra_cmd_tar := exec.Command("curl", "-L", tundra_tar_url, "-o", "package.tar.gz")
	err := tundra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tundra_zip_url := "https://github.com/deplinenoise/tundra/archive/refs/tags/v2.17.1.zip"
	tundra_cmd_zip := exec.Command("curl", "-L", tundra_zip_url, "-o", "package.zip")
	err = tundra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tundra_bin_url := "https://github.com/deplinenoise/tundra/archive/refs/tags/v2.17.1.bin"
	tundra_cmd_bin := exec.Command("curl", "-L", tundra_bin_url, "-o", "binary.bin")
	err = tundra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tundra_src_url := "https://github.com/deplinenoise/tundra/archive/refs/tags/v2.17.1.src.tar.gz"
	tundra_cmd_src := exec.Command("curl", "-L", tundra_src_url, "-o", "source.tar.gz")
	err = tundra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tundra_cmd_direct := exec.Command("./binary")
	err = tundra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: googletest")
exec.Command("latte", "install", "googletest")
}
