package main

// Ringojs - CommonJS-based JavaScript runtime
// Homepage: https://ringojs.org

import (
	"fmt"
	
	"os/exec"
)

func installRingojs() {
	// Método 1: Descargar y extraer .tar.gz
	ringojs_tar_url := "https://github.com/ringo/ringojs/releases/download/v4.0.0/ringojs-4.0.0.tar.gz"
	ringojs_cmd_tar := exec.Command("curl", "-L", ringojs_tar_url, "-o", "package.tar.gz")
	err := ringojs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ringojs_zip_url := "https://github.com/ringo/ringojs/releases/download/v4.0.0/ringojs-4.0.0.zip"
	ringojs_cmd_zip := exec.Command("curl", "-L", ringojs_zip_url, "-o", "package.zip")
	err = ringojs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ringojs_bin_url := "https://github.com/ringo/ringojs/releases/download/v4.0.0/ringojs-4.0.0.bin"
	ringojs_cmd_bin := exec.Command("curl", "-L", ringojs_bin_url, "-o", "binary.bin")
	err = ringojs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ringojs_src_url := "https://github.com/ringo/ringojs/releases/download/v4.0.0/ringojs-4.0.0.src.tar.gz"
	ringojs_cmd_src := exec.Command("curl", "-L", ringojs_src_url, "-o", "source.tar.gz")
	err = ringojs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ringojs_cmd_direct := exec.Command("./binary")
	err = ringojs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
exec.Command("latte", "install", "openjdk@17")
}
