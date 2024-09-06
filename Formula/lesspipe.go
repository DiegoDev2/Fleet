package main

// Lesspipe - Input filter for the pager less
// Homepage: https://www-zeuthen.desy.de/~friebel/unix/lesspipe.html

import (
	"fmt"
	
	"os/exec"
)

func installLesspipe() {
	// Método 1: Descargar y extraer .tar.gz
	lesspipe_tar_url := "https://github.com/wofr06/lesspipe/archive/refs/tags/v2.14.tar.gz"
	lesspipe_cmd_tar := exec.Command("curl", "-L", lesspipe_tar_url, "-o", "package.tar.gz")
	err := lesspipe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lesspipe_zip_url := "https://github.com/wofr06/lesspipe/archive/refs/tags/v2.14.zip"
	lesspipe_cmd_zip := exec.Command("curl", "-L", lesspipe_zip_url, "-o", "package.zip")
	err = lesspipe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lesspipe_bin_url := "https://github.com/wofr06/lesspipe/archive/refs/tags/v2.14.bin"
	lesspipe_cmd_bin := exec.Command("curl", "-L", lesspipe_bin_url, "-o", "binary.bin")
	err = lesspipe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lesspipe_src_url := "https://github.com/wofr06/lesspipe/archive/refs/tags/v2.14.src.tar.gz"
	lesspipe_cmd_src := exec.Command("curl", "-L", lesspipe_src_url, "-o", "source.tar.gz")
	err = lesspipe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lesspipe_cmd_direct := exec.Command("./binary")
	err = lesspipe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
