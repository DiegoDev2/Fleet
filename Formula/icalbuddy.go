package main

// IcalBuddy - Get events and tasks from the macOS calendar database
// Homepage: https://hasseg.org/icalBuddy/

import (
	"fmt"
	
	"os/exec"
)

func installIcalBuddy() {
	// Método 1: Descargar y extraer .tar.gz
	icalbuddy_tar_url := "https://github.com/dkaluta/icalBuddy64/archive/refs/tags/v1.10.1.tar.gz"
	icalbuddy_cmd_tar := exec.Command("curl", "-L", icalbuddy_tar_url, "-o", "package.tar.gz")
	err := icalbuddy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	icalbuddy_zip_url := "https://github.com/dkaluta/icalBuddy64/archive/refs/tags/v1.10.1.zip"
	icalbuddy_cmd_zip := exec.Command("curl", "-L", icalbuddy_zip_url, "-o", "package.zip")
	err = icalbuddy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	icalbuddy_bin_url := "https://github.com/dkaluta/icalBuddy64/archive/refs/tags/v1.10.1.bin"
	icalbuddy_cmd_bin := exec.Command("curl", "-L", icalbuddy_bin_url, "-o", "binary.bin")
	err = icalbuddy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	icalbuddy_src_url := "https://github.com/dkaluta/icalBuddy64/archive/refs/tags/v1.10.1.src.tar.gz"
	icalbuddy_cmd_src := exec.Command("curl", "-L", icalbuddy_src_url, "-o", "source.tar.gz")
	err = icalbuddy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	icalbuddy_cmd_direct := exec.Command("./binary")
	err = icalbuddy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
