package main

// C2048 - Console version of 2048
// Homepage: https://github.com/mevdschee/2048.c

import (
	"fmt"
	
	"os/exec"
)

func installC2048() {
	// Método 1: Descargar y extraer .tar.gz
	c2048_tar_url := "https://github.com/mevdschee/2048.c/archive/refs/tags/v1.0.1.tar.gz"
	c2048_cmd_tar := exec.Command("curl", "-L", c2048_tar_url, "-o", "package.tar.gz")
	err := c2048_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	c2048_zip_url := "https://github.com/mevdschee/2048.c/archive/refs/tags/v1.0.1.zip"
	c2048_cmd_zip := exec.Command("curl", "-L", c2048_zip_url, "-o", "package.zip")
	err = c2048_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	c2048_bin_url := "https://github.com/mevdschee/2048.c/archive/refs/tags/v1.0.1.bin"
	c2048_cmd_bin := exec.Command("curl", "-L", c2048_bin_url, "-o", "binary.bin")
	err = c2048_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	c2048_src_url := "https://github.com/mevdschee/2048.c/archive/refs/tags/v1.0.1.src.tar.gz"
	c2048_cmd_src := exec.Command("curl", "-L", c2048_src_url, "-o", "source.tar.gz")
	err = c2048_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	c2048_cmd_direct := exec.Command("./binary")
	err = c2048_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
