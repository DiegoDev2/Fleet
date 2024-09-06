package main

// Souffle - Logic Defined Static Analysis
// Homepage: https://souffle-lang.github.io

import (
	"fmt"
	
	"os/exec"
)

func installSouffle() {
	// Método 1: Descargar y extraer .tar.gz
	souffle_tar_url := "https://github.com/souffle-lang/souffle/archive/refs/tags/2.4.1.tar.gz"
	souffle_cmd_tar := exec.Command("curl", "-L", souffle_tar_url, "-o", "package.tar.gz")
	err := souffle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	souffle_zip_url := "https://github.com/souffle-lang/souffle/archive/refs/tags/2.4.1.zip"
	souffle_cmd_zip := exec.Command("curl", "-L", souffle_zip_url, "-o", "package.zip")
	err = souffle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	souffle_bin_url := "https://github.com/souffle-lang/souffle/archive/refs/tags/2.4.1.bin"
	souffle_cmd_bin := exec.Command("curl", "-L", souffle_bin_url, "-o", "binary.bin")
	err = souffle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	souffle_src_url := "https://github.com/souffle-lang/souffle/archive/refs/tags/2.4.1.src.tar.gz"
	souffle_cmd_src := exec.Command("curl", "-L", souffle_src_url, "-o", "source.tar.gz")
	err = souffle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	souffle_cmd_direct := exec.Command("./binary")
	err = souffle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: mcpp")
exec.Command("latte", "install", "mcpp")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
