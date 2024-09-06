package main

// ApacheBrooklynCli - Apache Brooklyn command-line interface
// Homepage: https://brooklyn.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installApacheBrooklynCli() {
	// Método 1: Descargar y extraer .tar.gz
	apachebrooklyncli_tar_url := "https://github.com/apache/brooklyn-client/archive/refs/tags/rel/apache-brooklyn-1.1.0.tar.gz"
	apachebrooklyncli_cmd_tar := exec.Command("curl", "-L", apachebrooklyncli_tar_url, "-o", "package.tar.gz")
	err := apachebrooklyncli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apachebrooklyncli_zip_url := "https://github.com/apache/brooklyn-client/archive/refs/tags/rel/apache-brooklyn-1.1.0.zip"
	apachebrooklyncli_cmd_zip := exec.Command("curl", "-L", apachebrooklyncli_zip_url, "-o", "package.zip")
	err = apachebrooklyncli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apachebrooklyncli_bin_url := "https://github.com/apache/brooklyn-client/archive/refs/tags/rel/apache-brooklyn-1.1.0.bin"
	apachebrooklyncli_cmd_bin := exec.Command("curl", "-L", apachebrooklyncli_bin_url, "-o", "binary.bin")
	err = apachebrooklyncli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apachebrooklyncli_src_url := "https://github.com/apache/brooklyn-client/archive/refs/tags/rel/apache-brooklyn-1.1.0.src.tar.gz"
	apachebrooklyncli_cmd_src := exec.Command("curl", "-L", apachebrooklyncli_src_url, "-o", "source.tar.gz")
	err = apachebrooklyncli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apachebrooklyncli_cmd_direct := exec.Command("./binary")
	err = apachebrooklyncli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
