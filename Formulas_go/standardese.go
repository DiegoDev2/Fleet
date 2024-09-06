package main

// Standardese - Next-gen documentation generator for C++
// Homepage: https://standardese.github.io

import (
	"fmt"
	
	"os/exec"
)

func installStandardese() {
	// Método 1: Descargar y extraer .tar.gz
	standardese_tar_url := "https://github.com/standardese/standardese.git"
	standardese_cmd_tar := exec.Command("curl", "-L", standardese_tar_url, "-o", "package.tar.gz")
	err := standardese_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	standardese_zip_url := "https://github.com/standardese/standardese.git"
	standardese_cmd_zip := exec.Command("curl", "-L", standardese_zip_url, "-o", "package.zip")
	err = standardese_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	standardese_bin_url := "https://github.com/standardese/standardese.git"
	standardese_cmd_bin := exec.Command("curl", "-L", standardese_bin_url, "-o", "binary.bin")
	err = standardese_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	standardese_src_url := "https://github.com/standardese/standardese.git"
	standardese_cmd_src := exec.Command("curl", "-L", standardese_src_url, "-o", "source.tar.gz")
	err = standardese_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	standardese_cmd_direct := exec.Command("./binary")
	err = standardese_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmark-gfm")
exec.Command("latte", "install", "cmark-gfm")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
