package main

// Aspell - Spell checker with better logic than ispell
// Homepage: http://aspell.net/

import (
	"fmt"
	
	"os/exec"
)

func installAspell() {
	// Método 1: Descargar y extraer .tar.gz
	aspell_tar_url := "https://ftp.gnu.org/gnu/aspell/aspell-0.60.8.1.tar.gz"
	aspell_cmd_tar := exec.Command("curl", "-L", aspell_tar_url, "-o", "package.tar.gz")
	err := aspell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aspell_zip_url := "https://ftp.gnu.org/gnu/aspell/aspell-0.60.8.1.zip"
	aspell_cmd_zip := exec.Command("curl", "-L", aspell_zip_url, "-o", "package.zip")
	err = aspell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aspell_bin_url := "https://ftp.gnu.org/gnu/aspell/aspell-0.60.8.1.bin"
	aspell_cmd_bin := exec.Command("curl", "-L", aspell_bin_url, "-o", "binary.bin")
	err = aspell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aspell_src_url := "https://ftp.gnu.org/gnu/aspell/aspell-0.60.8.1.src.tar.gz"
	aspell_cmd_src := exec.Command("curl", "-L", aspell_src_url, "-o", "source.tar.gz")
	err = aspell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aspell_cmd_direct := exec.Command("./binary")
	err = aspell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
