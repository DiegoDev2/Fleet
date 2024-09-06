package main

// Splint - Secure Programming Lint
// Homepage: https://github.com/splintchecker/splint

import (
	"fmt"
	
	"os/exec"
)

func installSplint() {
	// Método 1: Descargar y extraer .tar.gz
	splint_tar_url := "https://mirrorservice.org/sites/distfiles.macports.org/splint/splint-3.1.2.src.tgz"
	splint_cmd_tar := exec.Command("curl", "-L", splint_tar_url, "-o", "package.tar.gz")
	err := splint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	splint_zip_url := "https://mirrorservice.org/sites/distfiles.macports.org/splint/splint-3.1.2.src.tgz"
	splint_cmd_zip := exec.Command("curl", "-L", splint_zip_url, "-o", "package.zip")
	err = splint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	splint_bin_url := "https://mirrorservice.org/sites/distfiles.macports.org/splint/splint-3.1.2.src.tgz"
	splint_cmd_bin := exec.Command("curl", "-L", splint_bin_url, "-o", "binary.bin")
	err = splint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	splint_src_url := "https://mirrorservice.org/sites/distfiles.macports.org/splint/splint-3.1.2.src.tgz"
	splint_cmd_src := exec.Command("curl", "-L", splint_src_url, "-o", "source.tar.gz")
	err = splint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	splint_cmd_direct := exec.Command("./binary")
	err = splint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
