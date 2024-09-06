package main

// Gti - ASCII-art displaying typo-corrector for commands
// Homepage: https://r-wos.org/hacks/gti

import (
	"fmt"
	
	"os/exec"
)

func installGti() {
	// Método 1: Descargar y extraer .tar.gz
	gti_tar_url := "https://github.com/rwos/gti/archive/refs/tags/v1.9.1.tar.gz"
	gti_cmd_tar := exec.Command("curl", "-L", gti_tar_url, "-o", "package.tar.gz")
	err := gti_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gti_zip_url := "https://github.com/rwos/gti/archive/refs/tags/v1.9.1.zip"
	gti_cmd_zip := exec.Command("curl", "-L", gti_zip_url, "-o", "package.zip")
	err = gti_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gti_bin_url := "https://github.com/rwos/gti/archive/refs/tags/v1.9.1.bin"
	gti_cmd_bin := exec.Command("curl", "-L", gti_bin_url, "-o", "binary.bin")
	err = gti_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gti_src_url := "https://github.com/rwos/gti/archive/refs/tags/v1.9.1.src.tar.gz"
	gti_cmd_src := exec.Command("curl", "-L", gti_src_url, "-o", "source.tar.gz")
	err = gti_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gti_cmd_direct := exec.Command("./binary")
	err = gti_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
