package main

// Luit - Filter run between arbitrary application and UTF-8 terminal emulator
// Homepage: https://invisible-island.net/luit/

import (
	"fmt"
	
	"os/exec"
)

func installLuit() {
	// Método 1: Descargar y extraer .tar.gz
	luit_tar_url := "https://invisible-mirror.net/archives/luit/luit-20240102.tgz"
	luit_cmd_tar := exec.Command("curl", "-L", luit_tar_url, "-o", "package.tar.gz")
	err := luit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luit_zip_url := "https://invisible-mirror.net/archives/luit/luit-20240102.tgz"
	luit_cmd_zip := exec.Command("curl", "-L", luit_zip_url, "-o", "package.zip")
	err = luit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luit_bin_url := "https://invisible-mirror.net/archives/luit/luit-20240102.tgz"
	luit_cmd_bin := exec.Command("curl", "-L", luit_bin_url, "-o", "binary.bin")
	err = luit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luit_src_url := "https://invisible-mirror.net/archives/luit/luit-20240102.tgz"
	luit_cmd_src := exec.Command("curl", "-L", luit_src_url, "-o", "source.tar.gz")
	err = luit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luit_cmd_direct := exec.Command("./binary")
	err = luit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
