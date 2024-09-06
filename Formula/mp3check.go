package main

// Mp3check - Tool to check mp3 files for consistency
// Homepage: https://code.google.com/archive/p/mp3check/

import (
	"fmt"
	
	"os/exec"
)

func installMp3check() {
	// Método 1: Descargar y extraer .tar.gz
	mp3check_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/mp3check/mp3check-0.8.7.tgz"
	mp3check_cmd_tar := exec.Command("curl", "-L", mp3check_tar_url, "-o", "package.tar.gz")
	err := mp3check_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3check_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/mp3check/mp3check-0.8.7.tgz"
	mp3check_cmd_zip := exec.Command("curl", "-L", mp3check_zip_url, "-o", "package.zip")
	err = mp3check_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3check_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/mp3check/mp3check-0.8.7.tgz"
	mp3check_cmd_bin := exec.Command("curl", "-L", mp3check_bin_url, "-o", "binary.bin")
	err = mp3check_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3check_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/mp3check/mp3check-0.8.7.tgz"
	mp3check_cmd_src := exec.Command("curl", "-L", mp3check_src_url, "-o", "source.tar.gz")
	err = mp3check_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3check_cmd_direct := exec.Command("./binary")
	err = mp3check_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
