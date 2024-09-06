package main

// Wartremover - Flexible Scala code linting tool
// Homepage: https://github.com/wartremover/wartremover

import (
	"fmt"
	
	"os/exec"
)

func installWartremover() {
	// Método 1: Descargar y extraer .tar.gz
	wartremover_tar_url := "https://github.com/wartremover/wartremover/archive/refs/tags/v3.2.0.tar.gz"
	wartremover_cmd_tar := exec.Command("curl", "-L", wartremover_tar_url, "-o", "package.tar.gz")
	err := wartremover_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wartremover_zip_url := "https://github.com/wartremover/wartremover/archive/refs/tags/v3.2.0.zip"
	wartremover_cmd_zip := exec.Command("curl", "-L", wartremover_zip_url, "-o", "package.zip")
	err = wartremover_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wartremover_bin_url := "https://github.com/wartremover/wartremover/archive/refs/tags/v3.2.0.bin"
	wartremover_cmd_bin := exec.Command("curl", "-L", wartremover_bin_url, "-o", "binary.bin")
	err = wartremover_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wartremover_src_url := "https://github.com/wartremover/wartremover/archive/refs/tags/v3.2.0.src.tar.gz"
	wartremover_cmd_src := exec.Command("curl", "-L", wartremover_src_url, "-o", "source.tar.gz")
	err = wartremover_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wartremover_cmd_direct := exec.Command("./binary")
	err = wartremover_cmd_direct.Run()
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
