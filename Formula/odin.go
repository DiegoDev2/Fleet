package main

// Odin - Programming language with focus on simplicity, performance and modern systems
// Homepage: https://odin-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installOdin() {
	// Método 1: Descargar y extraer .tar.gz
	odin_tar_url := "https://github.com/odin-lang/Odin.git"
	odin_cmd_tar := exec.Command("curl", "-L", odin_tar_url, "-o", "package.tar.gz")
	err := odin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	odin_zip_url := "https://github.com/odin-lang/Odin.git"
	odin_cmd_zip := exec.Command("curl", "-L", odin_zip_url, "-o", "package.zip")
	err = odin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	odin_bin_url := "https://github.com/odin-lang/Odin.git"
	odin_cmd_bin := exec.Command("curl", "-L", odin_bin_url, "-o", "binary.bin")
	err = odin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	odin_src_url := "https://github.com/odin-lang/Odin.git"
	odin_cmd_src := exec.Command("curl", "-L", odin_src_url, "-o", "source.tar.gz")
	err = odin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	odin_cmd_direct := exec.Command("./binary")
	err = odin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glfw")
	exec.Command("latte", "install", "glfw").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: raylib")
	exec.Command("latte", "install", "raylib").Run()
}
