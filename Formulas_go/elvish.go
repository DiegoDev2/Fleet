package main

// Elvish - Friendly and expressive shell
// Homepage: https://elv.sh/

import (
	"fmt"
	
	"os/exec"
)

func installElvish() {
	// Método 1: Descargar y extraer .tar.gz
	elvish_tar_url := "https://github.com/elves/elvish/archive/refs/tags/v0.21.0.tar.gz"
	elvish_cmd_tar := exec.Command("curl", "-L", elvish_tar_url, "-o", "package.tar.gz")
	err := elvish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elvish_zip_url := "https://github.com/elves/elvish/archive/refs/tags/v0.21.0.zip"
	elvish_cmd_zip := exec.Command("curl", "-L", elvish_zip_url, "-o", "package.zip")
	err = elvish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elvish_bin_url := "https://github.com/elves/elvish/archive/refs/tags/v0.21.0.bin"
	elvish_cmd_bin := exec.Command("curl", "-L", elvish_bin_url, "-o", "binary.bin")
	err = elvish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elvish_src_url := "https://github.com/elves/elvish/archive/refs/tags/v0.21.0.src.tar.gz"
	elvish_cmd_src := exec.Command("curl", "-L", elvish_src_url, "-o", "source.tar.gz")
	err = elvish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elvish_cmd_direct := exec.Command("./binary")
	err = elvish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
