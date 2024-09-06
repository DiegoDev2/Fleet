package main

// WhisperkitCli - Swift native on-device speech recognition with Whisper for Apple Silicon
// Homepage: https://github.com/argmaxinc/WhisperKit

import (
	"fmt"
	
	"os/exec"
)

func installWhisperkitCli() {
	// Método 1: Descargar y extraer .tar.gz
	whisperkitcli_tar_url := "https://github.com/argmaxinc/WhisperKit/archive/refs/tags/v0.8.0.tar.gz"
	whisperkitcli_cmd_tar := exec.Command("curl", "-L", whisperkitcli_tar_url, "-o", "package.tar.gz")
	err := whisperkitcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	whisperkitcli_zip_url := "https://github.com/argmaxinc/WhisperKit/archive/refs/tags/v0.8.0.zip"
	whisperkitcli_cmd_zip := exec.Command("curl", "-L", whisperkitcli_zip_url, "-o", "package.zip")
	err = whisperkitcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	whisperkitcli_bin_url := "https://github.com/argmaxinc/WhisperKit/archive/refs/tags/v0.8.0.bin"
	whisperkitcli_cmd_bin := exec.Command("curl", "-L", whisperkitcli_bin_url, "-o", "binary.bin")
	err = whisperkitcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	whisperkitcli_src_url := "https://github.com/argmaxinc/WhisperKit/archive/refs/tags/v0.8.0.src.tar.gz"
	whisperkitcli_cmd_src := exec.Command("curl", "-L", whisperkitcli_src_url, "-o", "source.tar.gz")
	err = whisperkitcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	whisperkitcli_cmd_direct := exec.Command("./binary")
	err = whisperkitcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
