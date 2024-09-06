package main

// PhylumCli - Command-line interface for the Phylum API
// Homepage: https://www.phylum.io

import (
	"fmt"
	
	"os/exec"
)

func installPhylumCli() {
	// Método 1: Descargar y extraer .tar.gz
	phylumcli_tar_url := "https://github.com/phylum-dev/cli/archive/refs/tags/v6.6.6.tar.gz"
	phylumcli_cmd_tar := exec.Command("curl", "-L", phylumcli_tar_url, "-o", "package.tar.gz")
	err := phylumcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phylumcli_zip_url := "https://github.com/phylum-dev/cli/archive/refs/tags/v6.6.6.zip"
	phylumcli_cmd_zip := exec.Command("curl", "-L", phylumcli_zip_url, "-o", "package.zip")
	err = phylumcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phylumcli_bin_url := "https://github.com/phylum-dev/cli/archive/refs/tags/v6.6.6.bin"
	phylumcli_cmd_bin := exec.Command("curl", "-L", phylumcli_bin_url, "-o", "binary.bin")
	err = phylumcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phylumcli_src_url := "https://github.com/phylum-dev/cli/archive/refs/tags/v6.6.6.src.tar.gz"
	phylumcli_cmd_src := exec.Command("curl", "-L", phylumcli_src_url, "-o", "source.tar.gz")
	err = phylumcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phylumcli_cmd_direct := exec.Command("./binary")
	err = phylumcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
