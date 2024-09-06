package main

// Ktfmt - Kotlin code formatter
// Homepage: https://facebook.github.io/ktfmt/

import (
	"fmt"
	
	"os/exec"
)

func installKtfmt() {
	// Método 1: Descargar y extraer .tar.gz
	ktfmt_tar_url := "https://github.com/facebook/ktfmt/archive/refs/tags/v0.52.tar.gz"
	ktfmt_cmd_tar := exec.Command("curl", "-L", ktfmt_tar_url, "-o", "package.tar.gz")
	err := ktfmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ktfmt_zip_url := "https://github.com/facebook/ktfmt/archive/refs/tags/v0.52.zip"
	ktfmt_cmd_zip := exec.Command("curl", "-L", ktfmt_zip_url, "-o", "package.zip")
	err = ktfmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ktfmt_bin_url := "https://github.com/facebook/ktfmt/archive/refs/tags/v0.52.bin"
	ktfmt_cmd_bin := exec.Command("curl", "-L", ktfmt_bin_url, "-o", "binary.bin")
	err = ktfmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ktfmt_src_url := "https://github.com/facebook/ktfmt/archive/refs/tags/v0.52.src.tar.gz"
	ktfmt_cmd_src := exec.Command("curl", "-L", ktfmt_src_url, "-o", "source.tar.gz")
	err = ktfmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ktfmt_cmd_direct := exec.Command("./binary")
	err = ktfmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
