package main

// Mpssh - Mass parallel ssh
// Homepage: https://github.com/ndenev/mpssh

import (
	"fmt"
	
	"os/exec"
)

func installMpssh() {
	// Método 1: Descargar y extraer .tar.gz
	mpssh_tar_url := "https://github.com/ndenev/mpssh/archive/refs/tags/1.3.3.tar.gz"
	mpssh_cmd_tar := exec.Command("curl", "-L", mpssh_tar_url, "-o", "package.tar.gz")
	err := mpssh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpssh_zip_url := "https://github.com/ndenev/mpssh/archive/refs/tags/1.3.3.zip"
	mpssh_cmd_zip := exec.Command("curl", "-L", mpssh_zip_url, "-o", "package.zip")
	err = mpssh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpssh_bin_url := "https://github.com/ndenev/mpssh/archive/refs/tags/1.3.3.bin"
	mpssh_cmd_bin := exec.Command("curl", "-L", mpssh_bin_url, "-o", "binary.bin")
	err = mpssh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpssh_src_url := "https://github.com/ndenev/mpssh/archive/refs/tags/1.3.3.src.tar.gz"
	mpssh_cmd_src := exec.Command("curl", "-L", mpssh_src_url, "-o", "source.tar.gz")
	err = mpssh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpssh_cmd_direct := exec.Command("./binary")
	err = mpssh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
