package main

// GoogleSparsehash - Extremely memory-efficient hash_map implementation
// Homepage: https://github.com/sparsehash/sparsehash

import (
	"fmt"
	
	"os/exec"
)

func installGoogleSparsehash() {
	// Método 1: Descargar y extraer .tar.gz
	googlesparsehash_tar_url := "https://github.com/sparsehash/sparsehash/archive/refs/tags/sparsehash-2.0.4.tar.gz"
	googlesparsehash_cmd_tar := exec.Command("curl", "-L", googlesparsehash_tar_url, "-o", "package.tar.gz")
	err := googlesparsehash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	googlesparsehash_zip_url := "https://github.com/sparsehash/sparsehash/archive/refs/tags/sparsehash-2.0.4.zip"
	googlesparsehash_cmd_zip := exec.Command("curl", "-L", googlesparsehash_zip_url, "-o", "package.zip")
	err = googlesparsehash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	googlesparsehash_bin_url := "https://github.com/sparsehash/sparsehash/archive/refs/tags/sparsehash-2.0.4.bin"
	googlesparsehash_cmd_bin := exec.Command("curl", "-L", googlesparsehash_bin_url, "-o", "binary.bin")
	err = googlesparsehash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	googlesparsehash_src_url := "https://github.com/sparsehash/sparsehash/archive/refs/tags/sparsehash-2.0.4.src.tar.gz"
	googlesparsehash_cmd_src := exec.Command("curl", "-L", googlesparsehash_src_url, "-o", "source.tar.gz")
	err = googlesparsehash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	googlesparsehash_cmd_direct := exec.Command("./binary")
	err = googlesparsehash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
