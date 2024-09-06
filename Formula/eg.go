package main

// Eg - Expert Guide. Norton Guide Reader For GNU/Linux
// Homepage: https://github.com/davep/eg

import (
	"fmt"
	
	"os/exec"
)

func installEg() {
	// Método 1: Descargar y extraer .tar.gz
	eg_tar_url := "https://github.com/davep/eg/archive/refs/tags/v1.02.tar.gz"
	eg_cmd_tar := exec.Command("curl", "-L", eg_tar_url, "-o", "package.tar.gz")
	err := eg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eg_zip_url := "https://github.com/davep/eg/archive/refs/tags/v1.02.zip"
	eg_cmd_zip := exec.Command("curl", "-L", eg_zip_url, "-o", "package.zip")
	err = eg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eg_bin_url := "https://github.com/davep/eg/archive/refs/tags/v1.02.bin"
	eg_cmd_bin := exec.Command("curl", "-L", eg_bin_url, "-o", "binary.bin")
	err = eg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eg_src_url := "https://github.com/davep/eg/archive/refs/tags/v1.02.src.tar.gz"
	eg_cmd_src := exec.Command("curl", "-L", eg_src_url, "-o", "source.tar.gz")
	err = eg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eg_cmd_direct := exec.Command("./binary")
	err = eg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: s-lang")
	exec.Command("latte", "install", "s-lang").Run()
}
