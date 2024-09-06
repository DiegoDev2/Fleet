package main

// Sha3sum - Keccak, SHA-3, SHAKE, and RawSHAKE checksum utilities
// Homepage: https://codeberg.org/maandree/sha3sum

import (
	"fmt"
	
	"os/exec"
)

func installSha3sum() {
	// Método 1: Descargar y extraer .tar.gz
	sha3sum_tar_url := "https://codeberg.org/maandree/sha3sum/archive/1.2.3.1.tar.gz"
	sha3sum_cmd_tar := exec.Command("curl", "-L", sha3sum_tar_url, "-o", "package.tar.gz")
	err := sha3sum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sha3sum_zip_url := "https://codeberg.org/maandree/sha3sum/archive/1.2.3.1.zip"
	sha3sum_cmd_zip := exec.Command("curl", "-L", sha3sum_zip_url, "-o", "package.zip")
	err = sha3sum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sha3sum_bin_url := "https://codeberg.org/maandree/sha3sum/archive/1.2.3.1.bin"
	sha3sum_cmd_bin := exec.Command("curl", "-L", sha3sum_bin_url, "-o", "binary.bin")
	err = sha3sum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sha3sum_src_url := "https://codeberg.org/maandree/sha3sum/archive/1.2.3.1.src.tar.gz"
	sha3sum_cmd_src := exec.Command("curl", "-L", sha3sum_src_url, "-o", "source.tar.gz")
	err = sha3sum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sha3sum_cmd_direct := exec.Command("./binary")
	err = sha3sum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libkeccak")
exec.Command("latte", "install", "libkeccak")
}
