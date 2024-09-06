package main

// Nikto - Web server scanner
// Homepage: https://cirt.net/nikto2

import (
	"fmt"
	
	"os/exec"
)

func installNikto() {
	// Método 1: Descargar y extraer .tar.gz
	nikto_tar_url := "https://github.com/sullo/nikto/archive/refs/tags/2.5.0.tar.gz"
	nikto_cmd_tar := exec.Command("curl", "-L", nikto_tar_url, "-o", "package.tar.gz")
	err := nikto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nikto_zip_url := "https://github.com/sullo/nikto/archive/refs/tags/2.5.0.zip"
	nikto_cmd_zip := exec.Command("curl", "-L", nikto_zip_url, "-o", "package.zip")
	err = nikto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nikto_bin_url := "https://github.com/sullo/nikto/archive/refs/tags/2.5.0.bin"
	nikto_cmd_bin := exec.Command("curl", "-L", nikto_bin_url, "-o", "binary.bin")
	err = nikto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nikto_src_url := "https://github.com/sullo/nikto/archive/refs/tags/2.5.0.src.tar.gz"
	nikto_cmd_src := exec.Command("curl", "-L", nikto_src_url, "-o", "source.tar.gz")
	err = nikto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nikto_cmd_direct := exec.Command("./binary")
	err = nikto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
