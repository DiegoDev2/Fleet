package main

// Mint - Dependency manager that installs and runs Swift command-line tool packages
// Homepage: https://github.com/yonaskolb/Mint

import (
	"fmt"
	
	"os/exec"
)

func installMint() {
	// Método 1: Descargar y extraer .tar.gz
	mint_tar_url := "https://github.com/yonaskolb/Mint/archive/refs/tags/0.17.5.tar.gz"
	mint_cmd_tar := exec.Command("curl", "-L", mint_tar_url, "-o", "package.tar.gz")
	err := mint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mint_zip_url := "https://github.com/yonaskolb/Mint/archive/refs/tags/0.17.5.zip"
	mint_cmd_zip := exec.Command("curl", "-L", mint_zip_url, "-o", "package.zip")
	err = mint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mint_bin_url := "https://github.com/yonaskolb/Mint/archive/refs/tags/0.17.5.bin"
	mint_cmd_bin := exec.Command("curl", "-L", mint_bin_url, "-o", "binary.bin")
	err = mint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mint_src_url := "https://github.com/yonaskolb/Mint/archive/refs/tags/0.17.5.src.tar.gz"
	mint_cmd_src := exec.Command("curl", "-L", mint_src_url, "-o", "source.tar.gz")
	err = mint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mint_cmd_direct := exec.Command("./binary")
	err = mint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
