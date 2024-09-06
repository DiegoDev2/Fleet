package main

// ChibiScheme - Small footprint Scheme for use as a C Extension Language
// Homepage: https://github.com/ashinn/chibi-scheme

import (
	"fmt"
	
	"os/exec"
)

func installChibiScheme() {
	// Método 1: Descargar y extraer .tar.gz
	chibischeme_tar_url := "https://github.com/ashinn/chibi-scheme/archive/refs/tags/0.11.tar.gz"
	chibischeme_cmd_tar := exec.Command("curl", "-L", chibischeme_tar_url, "-o", "package.tar.gz")
	err := chibischeme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chibischeme_zip_url := "https://github.com/ashinn/chibi-scheme/archive/refs/tags/0.11.zip"
	chibischeme_cmd_zip := exec.Command("curl", "-L", chibischeme_zip_url, "-o", "package.zip")
	err = chibischeme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chibischeme_bin_url := "https://github.com/ashinn/chibi-scheme/archive/refs/tags/0.11.bin"
	chibischeme_cmd_bin := exec.Command("curl", "-L", chibischeme_bin_url, "-o", "binary.bin")
	err = chibischeme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chibischeme_src_url := "https://github.com/ashinn/chibi-scheme/archive/refs/tags/0.11.src.tar.gz"
	chibischeme_cmd_src := exec.Command("curl", "-L", chibischeme_src_url, "-o", "source.tar.gz")
	err = chibischeme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chibischeme_cmd_direct := exec.Command("./binary")
	err = chibischeme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
