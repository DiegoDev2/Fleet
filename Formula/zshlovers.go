package main

// ZshLovers - Tips, tricks, and examples for zsh
// Homepage: https://grml.org/zsh/#zshlovers

import (
	"fmt"
	
	"os/exec"
)

func installZshLovers() {
	// Método 1: Descargar y extraer .tar.gz
	zshlovers_tar_url := "https://deb.grml.org/pool/main/z/zsh-lovers/zsh-lovers_0.10.1_all.deb"
	zshlovers_cmd_tar := exec.Command("curl", "-L", zshlovers_tar_url, "-o", "package.tar.gz")
	err := zshlovers_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshlovers_zip_url := "https://deb.grml.org/pool/main/z/zsh-lovers/zsh-lovers_0.10.1_all.deb"
	zshlovers_cmd_zip := exec.Command("curl", "-L", zshlovers_zip_url, "-o", "package.zip")
	err = zshlovers_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshlovers_bin_url := "https://deb.grml.org/pool/main/z/zsh-lovers/zsh-lovers_0.10.1_all.deb"
	zshlovers_cmd_bin := exec.Command("curl", "-L", zshlovers_bin_url, "-o", "binary.bin")
	err = zshlovers_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshlovers_src_url := "https://deb.grml.org/pool/main/z/zsh-lovers/zsh-lovers_0.10.1_all.deb"
	zshlovers_cmd_src := exec.Command("curl", "-L", zshlovers_src_url, "-o", "source.tar.gz")
	err = zshlovers_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshlovers_cmd_direct := exec.Command("./binary")
	err = zshlovers_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
