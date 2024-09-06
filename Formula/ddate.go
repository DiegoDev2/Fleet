package main

// Ddate - Converts boring normal dates to fun Discordian Date
// Homepage: https://github.com/bo0ts/ddate

import (
	"fmt"
	
	"os/exec"
)

func installDdate() {
	// Método 1: Descargar y extraer .tar.gz
	ddate_tar_url := "https://github.com/bo0ts/ddate/archive/refs/tags/v0.2.2.tar.gz"
	ddate_cmd_tar := exec.Command("curl", "-L", ddate_tar_url, "-o", "package.tar.gz")
	err := ddate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddate_zip_url := "https://github.com/bo0ts/ddate/archive/refs/tags/v0.2.2.zip"
	ddate_cmd_zip := exec.Command("curl", "-L", ddate_zip_url, "-o", "package.zip")
	err = ddate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddate_bin_url := "https://github.com/bo0ts/ddate/archive/refs/tags/v0.2.2.bin"
	ddate_cmd_bin := exec.Command("curl", "-L", ddate_bin_url, "-o", "binary.bin")
	err = ddate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddate_src_url := "https://github.com/bo0ts/ddate/archive/refs/tags/v0.2.2.src.tar.gz"
	ddate_cmd_src := exec.Command("curl", "-L", ddate_src_url, "-o", "source.tar.gz")
	err = ddate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddate_cmd_direct := exec.Command("./binary")
	err = ddate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
