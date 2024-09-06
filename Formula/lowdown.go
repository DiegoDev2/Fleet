package main

// Lowdown - Simple markdown translator
// Homepage: https://kristaps.bsd.lv/lowdown

import (
	"fmt"
	
	"os/exec"
)

func installLowdown() {
	// Método 1: Descargar y extraer .tar.gz
	lowdown_tar_url := "https://github.com/kristapsdz/lowdown/archive/refs/tags/VERSION_1_1_0.tar.gz"
	lowdown_cmd_tar := exec.Command("curl", "-L", lowdown_tar_url, "-o", "package.tar.gz")
	err := lowdown_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lowdown_zip_url := "https://github.com/kristapsdz/lowdown/archive/refs/tags/VERSION_1_1_0.zip"
	lowdown_cmd_zip := exec.Command("curl", "-L", lowdown_zip_url, "-o", "package.zip")
	err = lowdown_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lowdown_bin_url := "https://github.com/kristapsdz/lowdown/archive/refs/tags/VERSION_1_1_0.bin"
	lowdown_cmd_bin := exec.Command("curl", "-L", lowdown_bin_url, "-o", "binary.bin")
	err = lowdown_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lowdown_src_url := "https://github.com/kristapsdz/lowdown/archive/refs/tags/VERSION_1_1_0.src.tar.gz"
	lowdown_cmd_src := exec.Command("curl", "-L", lowdown_src_url, "-o", "source.tar.gz")
	err = lowdown_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lowdown_cmd_direct := exec.Command("./binary")
	err = lowdown_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
