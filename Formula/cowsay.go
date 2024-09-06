package main

// Cowsay - Configurable talking characters in ASCII art
// Homepage: https://github.com/tnalpgge/rank-amateur-cowsay

import (
	"fmt"
	
	"os/exec"
)

func installCowsay() {
	// Método 1: Descargar y extraer .tar.gz
	cowsay_tar_url := "https://github.com/tnalpgge/rank-amateur-cowsay/archive/refs/tags/cowsay-3.04.tar.gz"
	cowsay_cmd_tar := exec.Command("curl", "-L", cowsay_tar_url, "-o", "package.tar.gz")
	err := cowsay_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cowsay_zip_url := "https://github.com/tnalpgge/rank-amateur-cowsay/archive/refs/tags/cowsay-3.04.zip"
	cowsay_cmd_zip := exec.Command("curl", "-L", cowsay_zip_url, "-o", "package.zip")
	err = cowsay_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cowsay_bin_url := "https://github.com/tnalpgge/rank-amateur-cowsay/archive/refs/tags/cowsay-3.04.bin"
	cowsay_cmd_bin := exec.Command("curl", "-L", cowsay_bin_url, "-o", "binary.bin")
	err = cowsay_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cowsay_src_url := "https://github.com/tnalpgge/rank-amateur-cowsay/archive/refs/tags/cowsay-3.04.src.tar.gz"
	cowsay_cmd_src := exec.Command("curl", "-L", cowsay_src_url, "-o", "source.tar.gz")
	err = cowsay_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cowsay_cmd_direct := exec.Command("./binary")
	err = cowsay_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
