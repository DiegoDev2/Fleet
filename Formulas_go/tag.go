package main

// Tag - Manipulate and query tags on macOS files
// Homepage: https://github.com/jdberry/tag/

import (
	"fmt"
	
	"os/exec"
)

func installTag() {
	// Método 1: Descargar y extraer .tar.gz
	tag_tar_url := "https://github.com/jdberry/tag/archive/refs/tags/v0.10.tar.gz"
	tag_cmd_tar := exec.Command("curl", "-L", tag_tar_url, "-o", "package.tar.gz")
	err := tag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tag_zip_url := "https://github.com/jdberry/tag/archive/refs/tags/v0.10.zip"
	tag_cmd_zip := exec.Command("curl", "-L", tag_zip_url, "-o", "package.zip")
	err = tag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tag_bin_url := "https://github.com/jdberry/tag/archive/refs/tags/v0.10.bin"
	tag_cmd_bin := exec.Command("curl", "-L", tag_bin_url, "-o", "binary.bin")
	err = tag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tag_src_url := "https://github.com/jdberry/tag/archive/refs/tags/v0.10.src.tar.gz"
	tag_cmd_src := exec.Command("curl", "-L", tag_src_url, "-o", "source.tar.gz")
	err = tag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tag_cmd_direct := exec.Command("./binary")
	err = tag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
