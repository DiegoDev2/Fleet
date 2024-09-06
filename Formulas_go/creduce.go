package main

// Creduce - Reduce a C/C++ program while keeping a property of interest
// Homepage: https://github.com/csmith-project/creduce

import (
	"fmt"
	
	"os/exec"
)

func installCreduce() {
	// Método 1: Descargar y extraer .tar.gz
	creduce_tar_url := "https://github.com/csmith-project/creduce/archive/refs/tags/creduce-2.10.0.tar.gz"
	creduce_cmd_tar := exec.Command("curl", "-L", creduce_tar_url, "-o", "package.tar.gz")
	err := creduce_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	creduce_zip_url := "https://github.com/csmith-project/creduce/archive/refs/tags/creduce-2.10.0.zip"
	creduce_cmd_zip := exec.Command("curl", "-L", creduce_zip_url, "-o", "package.zip")
	err = creduce_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	creduce_bin_url := "https://github.com/csmith-project/creduce/archive/refs/tags/creduce-2.10.0.bin"
	creduce_cmd_bin := exec.Command("curl", "-L", creduce_bin_url, "-o", "binary.bin")
	err = creduce_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	creduce_src_url := "https://github.com/csmith-project/creduce/archive/refs/tags/creduce-2.10.0.src.tar.gz"
	creduce_cmd_src := exec.Command("curl", "-L", creduce_src_url, "-o", "source.tar.gz")
	err = creduce_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	creduce_cmd_direct := exec.Command("./binary")
	err = creduce_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: astyle")
exec.Command("latte", "install", "astyle")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
