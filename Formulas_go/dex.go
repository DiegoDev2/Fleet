package main

// Dex - Dextrous text editor
// Homepage: https://github.com/tihirvon/dex

import (
	"fmt"
	
	"os/exec"
)

func installDex() {
	// Método 1: Descargar y extraer .tar.gz
	dex_tar_url := "https://github.com/tihirvon/dex/archive/refs/tags/v1.0.tar.gz"
	dex_cmd_tar := exec.Command("curl", "-L", dex_tar_url, "-o", "package.tar.gz")
	err := dex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dex_zip_url := "https://github.com/tihirvon/dex/archive/refs/tags/v1.0.zip"
	dex_cmd_zip := exec.Command("curl", "-L", dex_zip_url, "-o", "package.zip")
	err = dex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dex_bin_url := "https://github.com/tihirvon/dex/archive/refs/tags/v1.0.bin"
	dex_cmd_bin := exec.Command("curl", "-L", dex_bin_url, "-o", "binary.bin")
	err = dex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dex_src_url := "https://github.com/tihirvon/dex/archive/refs/tags/v1.0.src.tar.gz"
	dex_cmd_src := exec.Command("curl", "-L", dex_src_url, "-o", "source.tar.gz")
	err = dex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dex_cmd_direct := exec.Command("./binary")
	err = dex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
