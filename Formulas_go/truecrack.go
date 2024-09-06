package main

// Truecrack - Brute-force password cracker for TrueCrypt
// Homepage: https://github.com/lvaccaro/truecrack

import (
	"fmt"
	
	"os/exec"
)

func installTruecrack() {
	// Método 1: Descargar y extraer .tar.gz
	truecrack_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/truecrack/truecrack_v35.tar.gz"
	truecrack_cmd_tar := exec.Command("curl", "-L", truecrack_tar_url, "-o", "package.tar.gz")
	err := truecrack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	truecrack_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/truecrack/truecrack_v35.zip"
	truecrack_cmd_zip := exec.Command("curl", "-L", truecrack_zip_url, "-o", "package.zip")
	err = truecrack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	truecrack_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/truecrack/truecrack_v35.bin"
	truecrack_cmd_bin := exec.Command("curl", "-L", truecrack_bin_url, "-o", "binary.bin")
	err = truecrack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	truecrack_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/truecrack/truecrack_v35.src.tar.gz"
	truecrack_cmd_src := exec.Command("curl", "-L", truecrack_src_url, "-o", "source.tar.gz")
	err = truecrack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	truecrack_cmd_direct := exec.Command("./binary")
	err = truecrack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
