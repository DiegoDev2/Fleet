package main

// CodaCli - Shell integration for Panic's Coda
// Homepage: http://justinhileman.info/coda-cli/

import (
	"fmt"
	
	"os/exec"
)

func installCodaCli() {
	// Método 1: Descargar y extraer .tar.gz
	codacli_tar_url := "https://github.com/bobthecow/coda-cli/archive/refs/tags/v1.0.5.tar.gz"
	codacli_cmd_tar := exec.Command("curl", "-L", codacli_tar_url, "-o", "package.tar.gz")
	err := codacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codacli_zip_url := "https://github.com/bobthecow/coda-cli/archive/refs/tags/v1.0.5.zip"
	codacli_cmd_zip := exec.Command("curl", "-L", codacli_zip_url, "-o", "package.zip")
	err = codacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codacli_bin_url := "https://github.com/bobthecow/coda-cli/archive/refs/tags/v1.0.5.bin"
	codacli_cmd_bin := exec.Command("curl", "-L", codacli_bin_url, "-o", "binary.bin")
	err = codacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codacli_src_url := "https://github.com/bobthecow/coda-cli/archive/refs/tags/v1.0.5.src.tar.gz"
	codacli_cmd_src := exec.Command("curl", "-L", codacli_src_url, "-o", "source.tar.gz")
	err = codacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codacli_cmd_direct := exec.Command("./binary")
	err = codacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
