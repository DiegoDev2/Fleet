package main

// Xxhash - Extremely fast non-cryptographic hash algorithm
// Homepage: https://github.com/Cyan4973/xxHash

import (
	"fmt"
	
	"os/exec"
)

func installXxhash() {
	// Método 1: Descargar y extraer .tar.gz
	xxhash_tar_url := "https://github.com/Cyan4973/xxHash/archive/refs/tags/v0.8.2.tar.gz"
	xxhash_cmd_tar := exec.Command("curl", "-L", xxhash_tar_url, "-o", "package.tar.gz")
	err := xxhash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xxhash_zip_url := "https://github.com/Cyan4973/xxHash/archive/refs/tags/v0.8.2.zip"
	xxhash_cmd_zip := exec.Command("curl", "-L", xxhash_zip_url, "-o", "package.zip")
	err = xxhash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xxhash_bin_url := "https://github.com/Cyan4973/xxHash/archive/refs/tags/v0.8.2.bin"
	xxhash_cmd_bin := exec.Command("curl", "-L", xxhash_bin_url, "-o", "binary.bin")
	err = xxhash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xxhash_src_url := "https://github.com/Cyan4973/xxHash/archive/refs/tags/v0.8.2.src.tar.gz"
	xxhash_cmd_src := exec.Command("curl", "-L", xxhash_src_url, "-o", "source.tar.gz")
	err = xxhash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xxhash_cmd_direct := exec.Command("./binary")
	err = xxhash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
