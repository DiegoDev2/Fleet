package main

// Helidon - Command-line tool for Helidon application development
// Homepage: https://helidon.io/

import (
	"fmt"
	
	"os/exec"
)

func installHelidon() {
	// Método 1: Descargar y extraer .tar.gz
	helidon_tar_url := "https://github.com/helidon-io/helidon-build-tools/archive/refs/tags/3.0.6.tar.gz"
	helidon_cmd_tar := exec.Command("curl", "-L", helidon_tar_url, "-o", "package.tar.gz")
	err := helidon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	helidon_zip_url := "https://github.com/helidon-io/helidon-build-tools/archive/refs/tags/3.0.6.zip"
	helidon_cmd_zip := exec.Command("curl", "-L", helidon_zip_url, "-o", "package.zip")
	err = helidon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	helidon_bin_url := "https://github.com/helidon-io/helidon-build-tools/archive/refs/tags/3.0.6.bin"
	helidon_cmd_bin := exec.Command("curl", "-L", helidon_bin_url, "-o", "binary.bin")
	err = helidon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	helidon_src_url := "https://github.com/helidon-io/helidon-build-tools/archive/refs/tags/3.0.6.src.tar.gz"
	helidon_cmd_src := exec.Command("curl", "-L", helidon_src_url, "-o", "source.tar.gz")
	err = helidon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	helidon_cmd_direct := exec.Command("./binary")
	err = helidon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
