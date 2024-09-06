package main

// Skinny - Full-stack web app framework in Scala
// Homepage: https://skinny-framework.github.io

import (
	"fmt"
	
	"os/exec"
)

func installSkinny() {
	// Método 1: Descargar y extraer .tar.gz
	skinny_tar_url := "https://github.com/skinny-framework/skinny-framework/releases/download/4.0.1/skinny-4.0.1.tar.gz"
	skinny_cmd_tar := exec.Command("curl", "-L", skinny_tar_url, "-o", "package.tar.gz")
	err := skinny_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	skinny_zip_url := "https://github.com/skinny-framework/skinny-framework/releases/download/4.0.1/skinny-4.0.1.zip"
	skinny_cmd_zip := exec.Command("curl", "-L", skinny_zip_url, "-o", "package.zip")
	err = skinny_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	skinny_bin_url := "https://github.com/skinny-framework/skinny-framework/releases/download/4.0.1/skinny-4.0.1.bin"
	skinny_cmd_bin := exec.Command("curl", "-L", skinny_bin_url, "-o", "binary.bin")
	err = skinny_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	skinny_src_url := "https://github.com/skinny-framework/skinny-framework/releases/download/4.0.1/skinny-4.0.1.src.tar.gz"
	skinny_cmd_src := exec.Command("curl", "-L", skinny_src_url, "-o", "source.tar.gz")
	err = skinny_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	skinny_cmd_direct := exec.Command("./binary")
	err = skinny_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
