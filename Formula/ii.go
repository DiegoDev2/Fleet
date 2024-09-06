package main

// Ii - Minimalist IRC client
// Homepage: https://tools.suckless.org/ii/

import (
	"fmt"
	
	"os/exec"
)

func installIi() {
	// Método 1: Descargar y extraer .tar.gz
	ii_tar_url := "https://dl.suckless.org/tools/ii-2.0.tar.gz"
	ii_cmd_tar := exec.Command("curl", "-L", ii_tar_url, "-o", "package.tar.gz")
	err := ii_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ii_zip_url := "https://dl.suckless.org/tools/ii-2.0.zip"
	ii_cmd_zip := exec.Command("curl", "-L", ii_zip_url, "-o", "package.zip")
	err = ii_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ii_bin_url := "https://dl.suckless.org/tools/ii-2.0.bin"
	ii_cmd_bin := exec.Command("curl", "-L", ii_bin_url, "-o", "binary.bin")
	err = ii_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ii_src_url := "https://dl.suckless.org/tools/ii-2.0.src.tar.gz"
	ii_cmd_src := exec.Command("curl", "-L", ii_src_url, "-o", "source.tar.gz")
	err = ii_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ii_cmd_direct := exec.Command("./binary")
	err = ii_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
