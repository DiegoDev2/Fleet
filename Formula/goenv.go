package main

// Goenv - Go version management
// Homepage: https://github.com/go-nv/goenv

import (
	"fmt"
	
	"os/exec"
)

func installGoenv() {
	// Método 1: Descargar y extraer .tar.gz
	goenv_tar_url := "https://github.com/go-nv/goenv/archive/refs/tags/2.2.4.tar.gz"
	goenv_cmd_tar := exec.Command("curl", "-L", goenv_tar_url, "-o", "package.tar.gz")
	err := goenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goenv_zip_url := "https://github.com/go-nv/goenv/archive/refs/tags/2.2.4.zip"
	goenv_cmd_zip := exec.Command("curl", "-L", goenv_zip_url, "-o", "package.zip")
	err = goenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goenv_bin_url := "https://github.com/go-nv/goenv/archive/refs/tags/2.2.4.bin"
	goenv_cmd_bin := exec.Command("curl", "-L", goenv_bin_url, "-o", "binary.bin")
	err = goenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goenv_src_url := "https://github.com/go-nv/goenv/archive/refs/tags/2.2.4.src.tar.gz"
	goenv_cmd_src := exec.Command("curl", "-L", goenv_src_url, "-o", "source.tar.gz")
	err = goenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goenv_cmd_direct := exec.Command("./binary")
	err = goenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
