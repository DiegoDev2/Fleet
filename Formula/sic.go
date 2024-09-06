package main

// Sic - Minimal multiplexing IRC client
// Homepage: https://tools.suckless.org/sic/

import (
	"fmt"
	
	"os/exec"
)

func installSic() {
	// Método 1: Descargar y extraer .tar.gz
	sic_tar_url := "https://dl.suckless.org/tools/sic-1.3.tar.gz"
	sic_cmd_tar := exec.Command("curl", "-L", sic_tar_url, "-o", "package.tar.gz")
	err := sic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sic_zip_url := "https://dl.suckless.org/tools/sic-1.3.zip"
	sic_cmd_zip := exec.Command("curl", "-L", sic_zip_url, "-o", "package.zip")
	err = sic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sic_bin_url := "https://dl.suckless.org/tools/sic-1.3.bin"
	sic_cmd_bin := exec.Command("curl", "-L", sic_bin_url, "-o", "binary.bin")
	err = sic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sic_src_url := "https://dl.suckless.org/tools/sic-1.3.src.tar.gz"
	sic_cmd_src := exec.Command("curl", "-L", sic_src_url, "-o", "source.tar.gz")
	err = sic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sic_cmd_direct := exec.Command("./binary")
	err = sic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
