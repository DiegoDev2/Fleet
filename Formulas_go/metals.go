package main

// Metals - Scala language server
// Homepage: https://github.com/scalameta/metals

import (
	"fmt"
	
	"os/exec"
)

func installMetals() {
	// Método 1: Descargar y extraer .tar.gz
	metals_tar_url := "https://github.com/scalameta/metals/archive/refs/tags/v1.3.5.tar.gz"
	metals_cmd_tar := exec.Command("curl", "-L", metals_tar_url, "-o", "package.tar.gz")
	err := metals_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	metals_zip_url := "https://github.com/scalameta/metals/archive/refs/tags/v1.3.5.zip"
	metals_cmd_zip := exec.Command("curl", "-L", metals_zip_url, "-o", "package.zip")
	err = metals_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	metals_bin_url := "https://github.com/scalameta/metals/archive/refs/tags/v1.3.5.bin"
	metals_cmd_bin := exec.Command("curl", "-L", metals_bin_url, "-o", "binary.bin")
	err = metals_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	metals_src_url := "https://github.com/scalameta/metals/archive/refs/tags/v1.3.5.src.tar.gz"
	metals_cmd_src := exec.Command("curl", "-L", metals_src_url, "-o", "source.tar.gz")
	err = metals_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	metals_cmd_direct := exec.Command("./binary")
	err = metals_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbt")
exec.Command("latte", "install", "sbt")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
