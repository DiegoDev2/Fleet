package main

// Makeself - Generates a self-extracting compressed tar archive
// Homepage: https://makeself.io/

import (
	"fmt"
	
	"os/exec"
)

func installMakeself() {
	// Método 1: Descargar y extraer .tar.gz
	makeself_tar_url := "https://github.com/megastep/makeself/archive/refs/tags/release-2.5.0.tar.gz"
	makeself_cmd_tar := exec.Command("curl", "-L", makeself_tar_url, "-o", "package.tar.gz")
	err := makeself_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	makeself_zip_url := "https://github.com/megastep/makeself/archive/refs/tags/release-2.5.0.zip"
	makeself_cmd_zip := exec.Command("curl", "-L", makeself_zip_url, "-o", "package.zip")
	err = makeself_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	makeself_bin_url := "https://github.com/megastep/makeself/archive/refs/tags/release-2.5.0.bin"
	makeself_cmd_bin := exec.Command("curl", "-L", makeself_bin_url, "-o", "binary.bin")
	err = makeself_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	makeself_src_url := "https://github.com/megastep/makeself/archive/refs/tags/release-2.5.0.src.tar.gz"
	makeself_cmd_src := exec.Command("curl", "-L", makeself_src_url, "-o", "source.tar.gz")
	err = makeself_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	makeself_cmd_direct := exec.Command("./binary")
	err = makeself_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
