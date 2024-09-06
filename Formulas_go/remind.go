package main

// Remind - Sophisticated calendar and alarm
// Homepage: https://dianne.skoll.ca/projects/remind/

import (
	"fmt"
	
	"os/exec"
)

func installRemind() {
	// Método 1: Descargar y extraer .tar.gz
	remind_tar_url := "https://dianne.skoll.ca/projects/remind/download/remind-05.00.05.tar.gz"
	remind_cmd_tar := exec.Command("curl", "-L", remind_tar_url, "-o", "package.tar.gz")
	err := remind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	remind_zip_url := "https://dianne.skoll.ca/projects/remind/download/remind-05.00.05.zip"
	remind_cmd_zip := exec.Command("curl", "-L", remind_zip_url, "-o", "package.zip")
	err = remind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	remind_bin_url := "https://dianne.skoll.ca/projects/remind/download/remind-05.00.05.bin"
	remind_cmd_bin := exec.Command("curl", "-L", remind_bin_url, "-o", "binary.bin")
	err = remind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	remind_src_url := "https://dianne.skoll.ca/projects/remind/download/remind-05.00.05.src.tar.gz"
	remind_cmd_src := exec.Command("curl", "-L", remind_src_url, "-o", "source.tar.gz")
	err = remind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	remind_cmd_direct := exec.Command("./binary")
	err = remind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
