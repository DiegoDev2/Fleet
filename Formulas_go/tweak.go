package main

// Tweak - Command-line, ncurses library based hex editor
// Homepage: https://www.chiark.greenend.org.uk/~sgtatham/tweak/

import (
	"fmt"
	
	"os/exec"
)

func installTweak() {
	// Método 1: Descargar y extraer .tar.gz
	tweak_tar_url := "https://www.chiark.greenend.org.uk/~sgtatham/tweak/tweak-3.02.tar.gz"
	tweak_cmd_tar := exec.Command("curl", "-L", tweak_tar_url, "-o", "package.tar.gz")
	err := tweak_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tweak_zip_url := "https://www.chiark.greenend.org.uk/~sgtatham/tweak/tweak-3.02.zip"
	tweak_cmd_zip := exec.Command("curl", "-L", tweak_zip_url, "-o", "package.zip")
	err = tweak_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tweak_bin_url := "https://www.chiark.greenend.org.uk/~sgtatham/tweak/tweak-3.02.bin"
	tweak_cmd_bin := exec.Command("curl", "-L", tweak_bin_url, "-o", "binary.bin")
	err = tweak_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tweak_src_url := "https://www.chiark.greenend.org.uk/~sgtatham/tweak/tweak-3.02.src.tar.gz"
	tweak_cmd_src := exec.Command("curl", "-L", tweak_src_url, "-o", "source.tar.gz")
	err = tweak_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tweak_cmd_direct := exec.Command("./binary")
	err = tweak_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
