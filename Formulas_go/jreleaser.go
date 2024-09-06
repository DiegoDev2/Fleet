package main

// Jreleaser - Release projects quickly and easily with JReleaser
// Homepage: https://jreleaser.org/

import (
	"fmt"
	
	"os/exec"
)

func installJreleaser() {
	// Método 1: Descargar y extraer .tar.gz
	jreleaser_tar_url := "https://github.com/jreleaser/jreleaser/releases/download/v1.14.0/jreleaser-1.14.0.zip"
	jreleaser_cmd_tar := exec.Command("curl", "-L", jreleaser_tar_url, "-o", "package.tar.gz")
	err := jreleaser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jreleaser_zip_url := "https://github.com/jreleaser/jreleaser/releases/download/v1.14.0/jreleaser-1.14.0.zip"
	jreleaser_cmd_zip := exec.Command("curl", "-L", jreleaser_zip_url, "-o", "package.zip")
	err = jreleaser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jreleaser_bin_url := "https://github.com/jreleaser/jreleaser/releases/download/v1.14.0/jreleaser-1.14.0.zip"
	jreleaser_cmd_bin := exec.Command("curl", "-L", jreleaser_bin_url, "-o", "binary.bin")
	err = jreleaser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jreleaser_src_url := "https://github.com/jreleaser/jreleaser/releases/download/v1.14.0/jreleaser-1.14.0.zip"
	jreleaser_cmd_src := exec.Command("curl", "-L", jreleaser_src_url, "-o", "source.tar.gz")
	err = jreleaser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jreleaser_cmd_direct := exec.Command("./binary")
	err = jreleaser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
