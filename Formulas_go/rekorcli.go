package main

// RekorCli - CLI for interacting with Rekor
// Homepage: https://docs.sigstore.dev/logging/overview/

import (
	"fmt"
	
	"os/exec"
)

func installRekorCli() {
	// Método 1: Descargar y extraer .tar.gz
	rekorcli_tar_url := "https://github.com/sigstore/rekor/archive/refs/tags/v1.3.6.tar.gz"
	rekorcli_cmd_tar := exec.Command("curl", "-L", rekorcli_tar_url, "-o", "package.tar.gz")
	err := rekorcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rekorcli_zip_url := "https://github.com/sigstore/rekor/archive/refs/tags/v1.3.6.zip"
	rekorcli_cmd_zip := exec.Command("curl", "-L", rekorcli_zip_url, "-o", "package.zip")
	err = rekorcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rekorcli_bin_url := "https://github.com/sigstore/rekor/archive/refs/tags/v1.3.6.bin"
	rekorcli_cmd_bin := exec.Command("curl", "-L", rekorcli_bin_url, "-o", "binary.bin")
	err = rekorcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rekorcli_src_url := "https://github.com/sigstore/rekor/archive/refs/tags/v1.3.6.src.tar.gz"
	rekorcli_cmd_src := exec.Command("curl", "-L", rekorcli_src_url, "-o", "source.tar.gz")
	err = rekorcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rekorcli_cmd_direct := exec.Command("./binary")
	err = rekorcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
