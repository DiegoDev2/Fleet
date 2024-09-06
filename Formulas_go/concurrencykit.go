package main

// Concurrencykit - Aid design and implementation of concurrent systems
// Homepage: https://github.com/concurrencykit/ck

import (
	"fmt"
	
	"os/exec"
)

func installConcurrencykit() {
	// Método 1: Descargar y extraer .tar.gz
	concurrencykit_tar_url := "https://github.com/concurrencykit/ck/archive/refs/tags/0.7.2.tar.gz"
	concurrencykit_cmd_tar := exec.Command("curl", "-L", concurrencykit_tar_url, "-o", "package.tar.gz")
	err := concurrencykit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	concurrencykit_zip_url := "https://github.com/concurrencykit/ck/archive/refs/tags/0.7.2.zip"
	concurrencykit_cmd_zip := exec.Command("curl", "-L", concurrencykit_zip_url, "-o", "package.zip")
	err = concurrencykit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	concurrencykit_bin_url := "https://github.com/concurrencykit/ck/archive/refs/tags/0.7.2.bin"
	concurrencykit_cmd_bin := exec.Command("curl", "-L", concurrencykit_bin_url, "-o", "binary.bin")
	err = concurrencykit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	concurrencykit_src_url := "https://github.com/concurrencykit/ck/archive/refs/tags/0.7.2.src.tar.gz"
	concurrencykit_cmd_src := exec.Command("curl", "-L", concurrencykit_src_url, "-o", "source.tar.gz")
	err = concurrencykit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	concurrencykit_cmd_direct := exec.Command("./binary")
	err = concurrencykit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
