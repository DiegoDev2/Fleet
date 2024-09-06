package main

// Slugify - Convert filenames and directories to a web friendly format
// Homepage: https://github.com/benlinton/slugify

import (
	"fmt"
	
	"os/exec"
)

func installSlugify() {
	// Método 1: Descargar y extraer .tar.gz
	slugify_tar_url := "https://github.com/benlinton/slugify/archive/refs/tags/v1.0.1.tar.gz"
	slugify_cmd_tar := exec.Command("curl", "-L", slugify_tar_url, "-o", "package.tar.gz")
	err := slugify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slugify_zip_url := "https://github.com/benlinton/slugify/archive/refs/tags/v1.0.1.zip"
	slugify_cmd_zip := exec.Command("curl", "-L", slugify_zip_url, "-o", "package.zip")
	err = slugify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slugify_bin_url := "https://github.com/benlinton/slugify/archive/refs/tags/v1.0.1.bin"
	slugify_cmd_bin := exec.Command("curl", "-L", slugify_bin_url, "-o", "binary.bin")
	err = slugify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slugify_src_url := "https://github.com/benlinton/slugify/archive/refs/tags/v1.0.1.src.tar.gz"
	slugify_cmd_src := exec.Command("curl", "-L", slugify_src_url, "-o", "source.tar.gz")
	err = slugify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slugify_cmd_direct := exec.Command("./binary")
	err = slugify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
