package main

// LivekitCli - Command-line interface to LiveKit
// Homepage: https://livekit.io

import (
	"fmt"
	
	"os/exec"
)

func installLivekitCli() {
	// Método 1: Descargar y extraer .tar.gz
	livekitcli_tar_url := "https://github.com/livekit/livekit-cli/archive/refs/tags/v2.1.0.tar.gz"
	livekitcli_cmd_tar := exec.Command("curl", "-L", livekitcli_tar_url, "-o", "package.tar.gz")
	err := livekitcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	livekitcli_zip_url := "https://github.com/livekit/livekit-cli/archive/refs/tags/v2.1.0.zip"
	livekitcli_cmd_zip := exec.Command("curl", "-L", livekitcli_zip_url, "-o", "package.zip")
	err = livekitcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	livekitcli_bin_url := "https://github.com/livekit/livekit-cli/archive/refs/tags/v2.1.0.bin"
	livekitcli_cmd_bin := exec.Command("curl", "-L", livekitcli_bin_url, "-o", "binary.bin")
	err = livekitcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	livekitcli_src_url := "https://github.com/livekit/livekit-cli/archive/refs/tags/v2.1.0.src.tar.gz"
	livekitcli_cmd_src := exec.Command("curl", "-L", livekitcli_src_url, "-o", "source.tar.gz")
	err = livekitcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	livekitcli_cmd_direct := exec.Command("./binary")
	err = livekitcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
