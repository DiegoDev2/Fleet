package main

// Speedread - Simple terminal-based rapid serial visual presentation (RSVP) reader
// Homepage: https://github.com/pasky/speedread

import (
	"fmt"
	
	"os/exec"
)

func installSpeedread() {
	// Método 1: Descargar y extraer .tar.gz
	speedread_tar_url := "https://github.com/pasky/speedread/archive/refs/tags/v1.0.tar.gz"
	speedread_cmd_tar := exec.Command("curl", "-L", speedread_tar_url, "-o", "package.tar.gz")
	err := speedread_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	speedread_zip_url := "https://github.com/pasky/speedread/archive/refs/tags/v1.0.zip"
	speedread_cmd_zip := exec.Command("curl", "-L", speedread_zip_url, "-o", "package.zip")
	err = speedread_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	speedread_bin_url := "https://github.com/pasky/speedread/archive/refs/tags/v1.0.bin"
	speedread_cmd_bin := exec.Command("curl", "-L", speedread_bin_url, "-o", "binary.bin")
	err = speedread_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	speedread_src_url := "https://github.com/pasky/speedread/archive/refs/tags/v1.0.src.tar.gz"
	speedread_cmd_src := exec.Command("curl", "-L", speedread_src_url, "-o", "source.tar.gz")
	err = speedread_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	speedread_cmd_direct := exec.Command("./binary")
	err = speedread_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
