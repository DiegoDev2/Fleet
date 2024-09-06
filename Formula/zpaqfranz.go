package main

// Zpaqfranz - Deduplicating command-line archiver and backup tool
// Homepage: https://github.com/fcorbelli/zpaqfranz

import (
	"fmt"
	
	"os/exec"
)

func installZpaqfranz() {
	// Método 1: Descargar y extraer .tar.gz
	zpaqfranz_tar_url := "https://github.com/fcorbelli/zpaqfranz/archive/refs/tags/60.6.tar.gz"
	zpaqfranz_cmd_tar := exec.Command("curl", "-L", zpaqfranz_tar_url, "-o", "package.tar.gz")
	err := zpaqfranz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zpaqfranz_zip_url := "https://github.com/fcorbelli/zpaqfranz/archive/refs/tags/60.6.zip"
	zpaqfranz_cmd_zip := exec.Command("curl", "-L", zpaqfranz_zip_url, "-o", "package.zip")
	err = zpaqfranz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zpaqfranz_bin_url := "https://github.com/fcorbelli/zpaqfranz/archive/refs/tags/60.6.bin"
	zpaqfranz_cmd_bin := exec.Command("curl", "-L", zpaqfranz_bin_url, "-o", "binary.bin")
	err = zpaqfranz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zpaqfranz_src_url := "https://github.com/fcorbelli/zpaqfranz/archive/refs/tags/60.6.src.tar.gz"
	zpaqfranz_cmd_src := exec.Command("curl", "-L", zpaqfranz_src_url, "-o", "source.tar.gz")
	err = zpaqfranz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zpaqfranz_cmd_direct := exec.Command("./binary")
	err = zpaqfranz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
