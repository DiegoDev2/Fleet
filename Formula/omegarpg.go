package main

// OmegaRpg - Classic Roguelike game
// Homepage: https://web.archive.org/web/20220521053542/http://www.alcyone.com/max/projects/omega/

import (
	"fmt"
	
	"os/exec"
)

func installOmegaRpg() {
	// Método 1: Descargar y extraer .tar.gz
	omegarpg_tar_url := "https://web.archive.org/web/20190318143402/http://www.alcyone.com/binaries/omega/omega-0.80.2-src.tar.gz"
	omegarpg_cmd_tar := exec.Command("curl", "-L", omegarpg_tar_url, "-o", "package.tar.gz")
	err := omegarpg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	omegarpg_zip_url := "https://web.archive.org/web/20190318143402/http://www.alcyone.com/binaries/omega/omega-0.80.2-src.zip"
	omegarpg_cmd_zip := exec.Command("curl", "-L", omegarpg_zip_url, "-o", "package.zip")
	err = omegarpg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	omegarpg_bin_url := "https://web.archive.org/web/20190318143402/http://www.alcyone.com/binaries/omega/omega-0.80.2-src.bin"
	omegarpg_cmd_bin := exec.Command("curl", "-L", omegarpg_bin_url, "-o", "binary.bin")
	err = omegarpg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	omegarpg_src_url := "https://web.archive.org/web/20190318143402/http://www.alcyone.com/binaries/omega/omega-0.80.2-src.src.tar.gz"
	omegarpg_cmd_src := exec.Command("curl", "-L", omegarpg_src_url, "-o", "source.tar.gz")
	err = omegarpg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	omegarpg_cmd_direct := exec.Command("./binary")
	err = omegarpg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
