package main

// Clazy - Qt oriented static code analyzer
// Homepage: https://www.kdab.com/

import (
	"fmt"
	
	"os/exec"
)

func installClazy() {
	// Método 1: Descargar y extraer .tar.gz
	clazy_tar_url := "https://download.kde.org/stable/clazy/1.12/src/clazy-1.12.tar.xz"
	clazy_cmd_tar := exec.Command("curl", "-L", clazy_tar_url, "-o", "package.tar.gz")
	err := clazy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clazy_zip_url := "https://download.kde.org/stable/clazy/1.12/src/clazy-1.12.tar.xz"
	clazy_cmd_zip := exec.Command("curl", "-L", clazy_zip_url, "-o", "package.zip")
	err = clazy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clazy_bin_url := "https://download.kde.org/stable/clazy/1.12/src/clazy-1.12.tar.xz"
	clazy_cmd_bin := exec.Command("curl", "-L", clazy_bin_url, "-o", "binary.bin")
	err = clazy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clazy_src_url := "https://download.kde.org/stable/clazy/1.12/src/clazy-1.12.tar.xz"
	clazy_cmd_src := exec.Command("curl", "-L", clazy_src_url, "-o", "source.tar.gz")
	err = clazy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clazy_cmd_direct := exec.Command("./binary")
	err = clazy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
