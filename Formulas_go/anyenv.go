package main

// Anyenv - All in one for **env
// Homepage: https://anyenv.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installAnyenv() {
	// Método 1: Descargar y extraer .tar.gz
	anyenv_tar_url := "https://github.com/anyenv/anyenv/archive/refs/tags/v1.1.5.tar.gz"
	anyenv_cmd_tar := exec.Command("curl", "-L", anyenv_tar_url, "-o", "package.tar.gz")
	err := anyenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	anyenv_zip_url := "https://github.com/anyenv/anyenv/archive/refs/tags/v1.1.5.zip"
	anyenv_cmd_zip := exec.Command("curl", "-L", anyenv_zip_url, "-o", "package.zip")
	err = anyenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	anyenv_bin_url := "https://github.com/anyenv/anyenv/archive/refs/tags/v1.1.5.bin"
	anyenv_cmd_bin := exec.Command("curl", "-L", anyenv_bin_url, "-o", "binary.bin")
	err = anyenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	anyenv_src_url := "https://github.com/anyenv/anyenv/archive/refs/tags/v1.1.5.src.tar.gz"
	anyenv_cmd_src := exec.Command("curl", "-L", anyenv_src_url, "-o", "source.tar.gz")
	err = anyenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	anyenv_cmd_direct := exec.Command("./binary")
	err = anyenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
