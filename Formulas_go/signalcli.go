package main

// SignalCli - CLI and dbus interface for WhisperSystems/libsignal-service-java
// Homepage: https://github.com/AsamK/signal-cli

import (
	"fmt"
	
	"os/exec"
)

func installSignalCli() {
	// Método 1: Descargar y extraer .tar.gz
	signalcli_tar_url := "https://github.com/AsamK/signal-cli/archive/refs/tags/v0.13.5.tar.gz"
	signalcli_cmd_tar := exec.Command("curl", "-L", signalcli_tar_url, "-o", "package.tar.gz")
	err := signalcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	signalcli_zip_url := "https://github.com/AsamK/signal-cli/archive/refs/tags/v0.13.5.zip"
	signalcli_cmd_zip := exec.Command("curl", "-L", signalcli_zip_url, "-o", "package.zip")
	err = signalcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	signalcli_bin_url := "https://github.com/AsamK/signal-cli/archive/refs/tags/v0.13.5.bin"
	signalcli_cmd_bin := exec.Command("curl", "-L", signalcli_bin_url, "-o", "binary.bin")
	err = signalcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	signalcli_src_url := "https://github.com/AsamK/signal-cli/archive/refs/tags/v0.13.5.src.tar.gz"
	signalcli_cmd_src := exec.Command("curl", "-L", signalcli_src_url, "-o", "source.tar.gz")
	err = signalcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	signalcli_cmd_direct := exec.Command("./binary")
	err = signalcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gradle")
exec.Command("latte", "install", "gradle")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
	fmt.Println("Instalando dependencia: openjdk@21")
exec.Command("latte", "install", "openjdk@21")
}
