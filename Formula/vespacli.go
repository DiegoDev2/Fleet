package main

// VespaCli - Command-line tool for Vespa.ai
// Homepage: https://vespa.ai

import (
	"fmt"
	
	"os/exec"
)

func installVespaCli() {
	// Método 1: Descargar y extraer .tar.gz
	vespacli_tar_url := "https://github.com/vespa-engine/vespa/archive/refs/tags/v8.404.14.tar.gz"
	vespacli_cmd_tar := exec.Command("curl", "-L", vespacli_tar_url, "-o", "package.tar.gz")
	err := vespacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vespacli_zip_url := "https://github.com/vespa-engine/vespa/archive/refs/tags/v8.404.14.zip"
	vespacli_cmd_zip := exec.Command("curl", "-L", vespacli_zip_url, "-o", "package.zip")
	err = vespacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vespacli_bin_url := "https://github.com/vespa-engine/vespa/archive/refs/tags/v8.404.14.bin"
	vespacli_cmd_bin := exec.Command("curl", "-L", vespacli_bin_url, "-o", "binary.bin")
	err = vespacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vespacli_src_url := "https://github.com/vespa-engine/vespa/archive/refs/tags/v8.404.14.src.tar.gz"
	vespacli_cmd_src := exec.Command("curl", "-L", vespacli_src_url, "-o", "source.tar.gz")
	err = vespacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vespacli_cmd_direct := exec.Command("./binary")
	err = vespacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
