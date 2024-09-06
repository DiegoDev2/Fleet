package main

// Shpotify - Command-line interface for Spotify on a Mac
// Homepage: https://harishnarayanan.org/projects/shpotify/

import (
	"fmt"
	
	"os/exec"
)

func installShpotify() {
	// Método 1: Descargar y extraer .tar.gz
	shpotify_tar_url := "https://github.com/hnarayanan/shpotify/archive/refs/tags/2.1.tar.gz"
	shpotify_cmd_tar := exec.Command("curl", "-L", shpotify_tar_url, "-o", "package.tar.gz")
	err := shpotify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shpotify_zip_url := "https://github.com/hnarayanan/shpotify/archive/refs/tags/2.1.zip"
	shpotify_cmd_zip := exec.Command("curl", "-L", shpotify_zip_url, "-o", "package.zip")
	err = shpotify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shpotify_bin_url := "https://github.com/hnarayanan/shpotify/archive/refs/tags/2.1.bin"
	shpotify_cmd_bin := exec.Command("curl", "-L", shpotify_bin_url, "-o", "binary.bin")
	err = shpotify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shpotify_src_url := "https://github.com/hnarayanan/shpotify/archive/refs/tags/2.1.src.tar.gz"
	shpotify_cmd_src := exec.Command("curl", "-L", shpotify_src_url, "-o", "source.tar.gz")
	err = shpotify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shpotify_cmd_direct := exec.Command("./binary")
	err = shpotify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
