package main

// Ficy - Icecast/Shoutcast stream grabber suite
// Homepage: https://www.thregr.org/wavexx/software/fIcy/

import (
	"fmt"
	
	"os/exec"
)

func installFicy() {
	// Método 1: Descargar y extraer .tar.gz
	ficy_tar_url := "https://www.thregr.org/wavexx/software/fIcy/releases/fIcy-1.0.21.tar.gz"
	ficy_cmd_tar := exec.Command("curl", "-L", ficy_tar_url, "-o", "package.tar.gz")
	err := ficy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ficy_zip_url := "https://www.thregr.org/wavexx/software/fIcy/releases/fIcy-1.0.21.zip"
	ficy_cmd_zip := exec.Command("curl", "-L", ficy_zip_url, "-o", "package.zip")
	err = ficy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ficy_bin_url := "https://www.thregr.org/wavexx/software/fIcy/releases/fIcy-1.0.21.bin"
	ficy_cmd_bin := exec.Command("curl", "-L", ficy_bin_url, "-o", "binary.bin")
	err = ficy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ficy_src_url := "https://www.thregr.org/wavexx/software/fIcy/releases/fIcy-1.0.21.src.tar.gz"
	ficy_cmd_src := exec.Command("curl", "-L", ficy_src_url, "-o", "source.tar.gz")
	err = ficy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ficy_cmd_direct := exec.Command("./binary")
	err = ficy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
