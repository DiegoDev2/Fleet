package main

// KosliCli - CLI for managing Kosli
// Homepage: https://docs.kosli.com/client_reference/

import (
	"fmt"
	
	"os/exec"
)

func installKosliCli() {
	// Método 1: Descargar y extraer .tar.gz
	koslicli_tar_url := "https://github.com/kosli-dev/cli/archive/refs/tags/v2.10.14.tar.gz"
	koslicli_cmd_tar := exec.Command("curl", "-L", koslicli_tar_url, "-o", "package.tar.gz")
	err := koslicli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	koslicli_zip_url := "https://github.com/kosli-dev/cli/archive/refs/tags/v2.10.14.zip"
	koslicli_cmd_zip := exec.Command("curl", "-L", koslicli_zip_url, "-o", "package.zip")
	err = koslicli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	koslicli_bin_url := "https://github.com/kosli-dev/cli/archive/refs/tags/v2.10.14.bin"
	koslicli_cmd_bin := exec.Command("curl", "-L", koslicli_bin_url, "-o", "binary.bin")
	err = koslicli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	koslicli_src_url := "https://github.com/kosli-dev/cli/archive/refs/tags/v2.10.14.src.tar.gz"
	koslicli_cmd_src := exec.Command("curl", "-L", koslicli_src_url, "-o", "source.tar.gz")
	err = koslicli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	koslicli_cmd_direct := exec.Command("./binary")
	err = koslicli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
