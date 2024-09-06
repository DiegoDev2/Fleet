package main

// Never - Statically typed, embedded functional programming language
// Homepage: https://never-lang.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installNever() {
	// Método 1: Descargar y extraer .tar.gz
	never_tar_url := "https://github.com/never-lang/never/archive/refs/tags/v2.3.9.tar.gz"
	never_cmd_tar := exec.Command("curl", "-L", never_tar_url, "-o", "package.tar.gz")
	err := never_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	never_zip_url := "https://github.com/never-lang/never/archive/refs/tags/v2.3.9.zip"
	never_cmd_zip := exec.Command("curl", "-L", never_zip_url, "-o", "package.zip")
	err = never_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	never_bin_url := "https://github.com/never-lang/never/archive/refs/tags/v2.3.9.bin"
	never_cmd_bin := exec.Command("curl", "-L", never_bin_url, "-o", "binary.bin")
	err = never_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	never_src_url := "https://github.com/never-lang/never/archive/refs/tags/v2.3.9.src.tar.gz"
	never_cmd_src := exec.Command("curl", "-L", never_src_url, "-o", "source.tar.gz")
	err = never_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	never_cmd_direct := exec.Command("./binary")
	err = never_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
