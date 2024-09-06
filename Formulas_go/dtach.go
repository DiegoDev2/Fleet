package main

// Dtach - Emulates the detach feature of screen
// Homepage: https://dtach.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDtach() {
	// Método 1: Descargar y extraer .tar.gz
	dtach_tar_url := "https://downloads.sourceforge.net/project/dtach/dtach/0.9/dtach-0.9.tar.gz"
	dtach_cmd_tar := exec.Command("curl", "-L", dtach_tar_url, "-o", "package.tar.gz")
	err := dtach_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dtach_zip_url := "https://downloads.sourceforge.net/project/dtach/dtach/0.9/dtach-0.9.zip"
	dtach_cmd_zip := exec.Command("curl", "-L", dtach_zip_url, "-o", "package.zip")
	err = dtach_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dtach_bin_url := "https://downloads.sourceforge.net/project/dtach/dtach/0.9/dtach-0.9.bin"
	dtach_cmd_bin := exec.Command("curl", "-L", dtach_bin_url, "-o", "binary.bin")
	err = dtach_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dtach_src_url := "https://downloads.sourceforge.net/project/dtach/dtach/0.9/dtach-0.9.src.tar.gz"
	dtach_cmd_src := exec.Command("curl", "-L", dtach_src_url, "-o", "source.tar.gz")
	err = dtach_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dtach_cmd_direct := exec.Command("./binary")
	err = dtach_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
