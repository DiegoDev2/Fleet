package main

// Mpck - Check MP3 files for errors
// Homepage: https://checkmate.gissen.nl/

import (
	"fmt"
	
	"os/exec"
)

func installMpck() {
	// Método 1: Descargar y extraer .tar.gz
	mpck_tar_url := "https://checkmate.gissen.nl/checkmate-0.21.tar.gz"
	mpck_cmd_tar := exec.Command("curl", "-L", mpck_tar_url, "-o", "package.tar.gz")
	err := mpck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpck_zip_url := "https://checkmate.gissen.nl/checkmate-0.21.zip"
	mpck_cmd_zip := exec.Command("curl", "-L", mpck_zip_url, "-o", "package.zip")
	err = mpck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpck_bin_url := "https://checkmate.gissen.nl/checkmate-0.21.bin"
	mpck_cmd_bin := exec.Command("curl", "-L", mpck_bin_url, "-o", "binary.bin")
	err = mpck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpck_src_url := "https://checkmate.gissen.nl/checkmate-0.21.src.tar.gz"
	mpck_cmd_src := exec.Command("curl", "-L", mpck_src_url, "-o", "source.tar.gz")
	err = mpck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpck_cmd_direct := exec.Command("./binary")
	err = mpck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
