package main

// Mill - Scala build tool
// Homepage: https://mill-build.com/mill/Scala_Intro_to_Mill.html

import (
	"fmt"
	
	"os/exec"
)

func installMill() {
	// Método 1: Descargar y extraer .tar.gz
	mill_tar_url := "https://github.com/com-lihaoyi/mill/releases/download/0.11.12/0.11.12-assembly"
	mill_cmd_tar := exec.Command("curl", "-L", mill_tar_url, "-o", "package.tar.gz")
	err := mill_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mill_zip_url := "https://github.com/com-lihaoyi/mill/releases/download/0.11.12/0.11.12-assembly"
	mill_cmd_zip := exec.Command("curl", "-L", mill_zip_url, "-o", "package.zip")
	err = mill_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mill_bin_url := "https://github.com/com-lihaoyi/mill/releases/download/0.11.12/0.11.12-assembly"
	mill_cmd_bin := exec.Command("curl", "-L", mill_bin_url, "-o", "binary.bin")
	err = mill_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mill_src_url := "https://github.com/com-lihaoyi/mill/releases/download/0.11.12/0.11.12-assembly"
	mill_cmd_src := exec.Command("curl", "-L", mill_src_url, "-o", "source.tar.gz")
	err = mill_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mill_cmd_direct := exec.Command("./binary")
	err = mill_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
