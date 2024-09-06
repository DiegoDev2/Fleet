package main

// EnpassCli - Enpass command-line client
// Homepage: https://github.com/hazcod/enpass-cli

import (
	"fmt"
	
	"os/exec"
)

func installEnpassCli() {
	// Método 1: Descargar y extraer .tar.gz
	enpasscli_tar_url := "https://github.com/hazcod/enpass-cli/archive/refs/tags/v1.6.1.tar.gz"
	enpasscli_cmd_tar := exec.Command("curl", "-L", enpasscli_tar_url, "-o", "package.tar.gz")
	err := enpasscli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enpasscli_zip_url := "https://github.com/hazcod/enpass-cli/archive/refs/tags/v1.6.1.zip"
	enpasscli_cmd_zip := exec.Command("curl", "-L", enpasscli_zip_url, "-o", "package.zip")
	err = enpasscli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enpasscli_bin_url := "https://github.com/hazcod/enpass-cli/archive/refs/tags/v1.6.1.bin"
	enpasscli_cmd_bin := exec.Command("curl", "-L", enpasscli_bin_url, "-o", "binary.bin")
	err = enpasscli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enpasscli_src_url := "https://github.com/hazcod/enpass-cli/archive/refs/tags/v1.6.1.src.tar.gz"
	enpasscli_cmd_src := exec.Command("curl", "-L", enpasscli_src_url, "-o", "source.tar.gz")
	err = enpasscli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enpasscli_cmd_direct := exec.Command("./binary")
	err = enpasscli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
