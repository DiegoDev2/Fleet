package main

// Plank - Framework for generating immutable model objects
// Homepage: https://pinterest.github.io/plank/

import (
	"fmt"
	
	"os/exec"
)

func installPlank() {
	// Método 1: Descargar y extraer .tar.gz
	plank_tar_url := "https://github.com/pinterest/plank/archive/refs/tags/v1.6.tar.gz"
	plank_cmd_tar := exec.Command("curl", "-L", plank_tar_url, "-o", "package.tar.gz")
	err := plank_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plank_zip_url := "https://github.com/pinterest/plank/archive/refs/tags/v1.6.zip"
	plank_cmd_zip := exec.Command("curl", "-L", plank_zip_url, "-o", "package.zip")
	err = plank_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plank_bin_url := "https://github.com/pinterest/plank/archive/refs/tags/v1.6.bin"
	plank_cmd_bin := exec.Command("curl", "-L", plank_bin_url, "-o", "binary.bin")
	err = plank_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plank_src_url := "https://github.com/pinterest/plank/archive/refs/tags/v1.6.src.tar.gz"
	plank_cmd_src := exec.Command("curl", "-L", plank_src_url, "-o", "source.tar.gz")
	err = plank_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plank_cmd_direct := exec.Command("./binary")
	err = plank_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
