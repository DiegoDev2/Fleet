package main

// Screenfetch - Generate ASCII art with terminal, shell, and OS info
// Homepage: https://github.com/KittyKatt/screenFetch

import (
	"fmt"
	
	"os/exec"
)

func installScreenfetch() {
	// Método 1: Descargar y extraer .tar.gz
	screenfetch_tar_url := "https://github.com/KittyKatt/screenFetch/archive/refs/tags/v3.9.1.tar.gz"
	screenfetch_cmd_tar := exec.Command("curl", "-L", screenfetch_tar_url, "-o", "package.tar.gz")
	err := screenfetch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	screenfetch_zip_url := "https://github.com/KittyKatt/screenFetch/archive/refs/tags/v3.9.1.zip"
	screenfetch_cmd_zip := exec.Command("curl", "-L", screenfetch_zip_url, "-o", "package.zip")
	err = screenfetch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	screenfetch_bin_url := "https://github.com/KittyKatt/screenFetch/archive/refs/tags/v3.9.1.bin"
	screenfetch_cmd_bin := exec.Command("curl", "-L", screenfetch_bin_url, "-o", "binary.bin")
	err = screenfetch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	screenfetch_src_url := "https://github.com/KittyKatt/screenFetch/archive/refs/tags/v3.9.1.src.tar.gz"
	screenfetch_cmd_src := exec.Command("curl", "-L", screenfetch_src_url, "-o", "source.tar.gz")
	err = screenfetch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	screenfetch_cmd_direct := exec.Command("./binary")
	err = screenfetch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
