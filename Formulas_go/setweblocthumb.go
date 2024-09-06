package main

// Setweblocthumb - Assigns custom icons to webloc files
// Homepage: https://hasseg.org/setWeblocThumb/

import (
	"fmt"
	
	"os/exec"
)

func installSetweblocthumb() {
	// Método 1: Descargar y extraer .tar.gz
	setweblocthumb_tar_url := "https://github.com/ali-rantakari/setWeblocThumb/archive/refs/tags/v1.0.0.tar.gz"
	setweblocthumb_cmd_tar := exec.Command("curl", "-L", setweblocthumb_tar_url, "-o", "package.tar.gz")
	err := setweblocthumb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	setweblocthumb_zip_url := "https://github.com/ali-rantakari/setWeblocThumb/archive/refs/tags/v1.0.0.zip"
	setweblocthumb_cmd_zip := exec.Command("curl", "-L", setweblocthumb_zip_url, "-o", "package.zip")
	err = setweblocthumb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	setweblocthumb_bin_url := "https://github.com/ali-rantakari/setWeblocThumb/archive/refs/tags/v1.0.0.bin"
	setweblocthumb_cmd_bin := exec.Command("curl", "-L", setweblocthumb_bin_url, "-o", "binary.bin")
	err = setweblocthumb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	setweblocthumb_src_url := "https://github.com/ali-rantakari/setWeblocThumb/archive/refs/tags/v1.0.0.src.tar.gz"
	setweblocthumb_cmd_src := exec.Command("curl", "-L", setweblocthumb_src_url, "-o", "source.tar.gz")
	err = setweblocthumb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	setweblocthumb_cmd_direct := exec.Command("./binary")
	err = setweblocthumb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
