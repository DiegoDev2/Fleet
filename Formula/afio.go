package main

// Afio - Creates cpio-format archives
// Homepage: https://github.com/kholtman/afio

import (
	"fmt"
	
	"os/exec"
)

func installAfio() {
	// Método 1: Descargar y extraer .tar.gz
	afio_tar_url := "https://github.com/kholtman/afio/archive/refs/tags/v2.5.2.tar.gz"
	afio_cmd_tar := exec.Command("curl", "-L", afio_tar_url, "-o", "package.tar.gz")
	err := afio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	afio_zip_url := "https://github.com/kholtman/afio/archive/refs/tags/v2.5.2.zip"
	afio_cmd_zip := exec.Command("curl", "-L", afio_zip_url, "-o", "package.zip")
	err = afio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	afio_bin_url := "https://github.com/kholtman/afio/archive/refs/tags/v2.5.2.bin"
	afio_cmd_bin := exec.Command("curl", "-L", afio_bin_url, "-o", "binary.bin")
	err = afio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	afio_src_url := "https://github.com/kholtman/afio/archive/refs/tags/v2.5.2.src.tar.gz"
	afio_cmd_src := exec.Command("curl", "-L", afio_src_url, "-o", "source.tar.gz")
	err = afio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	afio_cmd_direct := exec.Command("./binary")
	err = afio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
