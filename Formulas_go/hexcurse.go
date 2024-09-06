package main

// Hexcurse - Ncurses-based console hex editor
// Homepage: https://github.com/LonnyGomes/hexcurse

import (
	"fmt"
	
	"os/exec"
)

func installHexcurse() {
	// Método 1: Descargar y extraer .tar.gz
	hexcurse_tar_url := "https://github.com/LonnyGomes/hexcurse/archive/refs/tags/v1.60.0.tar.gz"
	hexcurse_cmd_tar := exec.Command("curl", "-L", hexcurse_tar_url, "-o", "package.tar.gz")
	err := hexcurse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hexcurse_zip_url := "https://github.com/LonnyGomes/hexcurse/archive/refs/tags/v1.60.0.zip"
	hexcurse_cmd_zip := exec.Command("curl", "-L", hexcurse_zip_url, "-o", "package.zip")
	err = hexcurse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hexcurse_bin_url := "https://github.com/LonnyGomes/hexcurse/archive/refs/tags/v1.60.0.bin"
	hexcurse_cmd_bin := exec.Command("curl", "-L", hexcurse_bin_url, "-o", "binary.bin")
	err = hexcurse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hexcurse_src_url := "https://github.com/LonnyGomes/hexcurse/archive/refs/tags/v1.60.0.src.tar.gz"
	hexcurse_cmd_src := exec.Command("curl", "-L", hexcurse_src_url, "-o", "source.tar.gz")
	err = hexcurse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hexcurse_cmd_direct := exec.Command("./binary")
	err = hexcurse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
