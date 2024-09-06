package main

// Keychain - User-friendly front-end to ssh-agent(1)
// Homepage: https://www.funtoo.org/Keychain

import (
	"fmt"
	
	"os/exec"
)

func installKeychain() {
	// Método 1: Descargar y extraer .tar.gz
	keychain_tar_url := "https://github.com/funtoo/keychain/archive/refs/tags/2.8.5.tar.gz"
	keychain_cmd_tar := exec.Command("curl", "-L", keychain_tar_url, "-o", "package.tar.gz")
	err := keychain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	keychain_zip_url := "https://github.com/funtoo/keychain/archive/refs/tags/2.8.5.zip"
	keychain_cmd_zip := exec.Command("curl", "-L", keychain_zip_url, "-o", "package.zip")
	err = keychain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	keychain_bin_url := "https://github.com/funtoo/keychain/archive/refs/tags/2.8.5.bin"
	keychain_cmd_bin := exec.Command("curl", "-L", keychain_bin_url, "-o", "binary.bin")
	err = keychain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	keychain_src_url := "https://github.com/funtoo/keychain/archive/refs/tags/2.8.5.src.tar.gz"
	keychain_cmd_src := exec.Command("curl", "-L", keychain_src_url, "-o", "source.tar.gz")
	err = keychain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	keychain_cmd_direct := exec.Command("./binary")
	err = keychain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
