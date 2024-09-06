package main

// TrezorBridge - Trezor Communication Daemon
// Homepage: https://github.com/trezor/trezord-go

import (
	"fmt"
	
	"os/exec"
)

func installTrezorBridge() {
	// Método 1: Descargar y extraer .tar.gz
	trezorbridge_tar_url := "https://github.com/trezor/trezord-go.git"
	trezorbridge_cmd_tar := exec.Command("curl", "-L", trezorbridge_tar_url, "-o", "package.tar.gz")
	err := trezorbridge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trezorbridge_zip_url := "https://github.com/trezor/trezord-go.git"
	trezorbridge_cmd_zip := exec.Command("curl", "-L", trezorbridge_zip_url, "-o", "package.zip")
	err = trezorbridge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trezorbridge_bin_url := "https://github.com/trezor/trezord-go.git"
	trezorbridge_cmd_bin := exec.Command("curl", "-L", trezorbridge_bin_url, "-o", "binary.bin")
	err = trezorbridge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trezorbridge_src_url := "https://github.com/trezor/trezord-go.git"
	trezorbridge_cmd_src := exec.Command("curl", "-L", trezorbridge_src_url, "-o", "source.tar.gz")
	err = trezorbridge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trezorbridge_cmd_direct := exec.Command("./binary")
	err = trezorbridge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
