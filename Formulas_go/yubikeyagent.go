package main

// YubikeyAgent - Seamless ssh-agent for YubiKeys and other PIV tokens
// Homepage: https://filippo.io/yubikey-agent

import (
	"fmt"
	
	"os/exec"
)

func installYubikeyAgent() {
	// Método 1: Descargar y extraer .tar.gz
	yubikeyagent_tar_url := "https://github.com/FiloSottile/yubikey-agent/archive/refs/tags/v0.1.6.tar.gz"
	yubikeyagent_cmd_tar := exec.Command("curl", "-L", yubikeyagent_tar_url, "-o", "package.tar.gz")
	err := yubikeyagent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yubikeyagent_zip_url := "https://github.com/FiloSottile/yubikey-agent/archive/refs/tags/v0.1.6.zip"
	yubikeyagent_cmd_zip := exec.Command("curl", "-L", yubikeyagent_zip_url, "-o", "package.zip")
	err = yubikeyagent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yubikeyagent_bin_url := "https://github.com/FiloSottile/yubikey-agent/archive/refs/tags/v0.1.6.bin"
	yubikeyagent_cmd_bin := exec.Command("curl", "-L", yubikeyagent_bin_url, "-o", "binary.bin")
	err = yubikeyagent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yubikeyagent_src_url := "https://github.com/FiloSottile/yubikey-agent/archive/refs/tags/v0.1.6.src.tar.gz"
	yubikeyagent_cmd_src := exec.Command("curl", "-L", yubikeyagent_src_url, "-o", "source.tar.gz")
	err = yubikeyagent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yubikeyagent_cmd_direct := exec.Command("./binary")
	err = yubikeyagent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: pinentry")
exec.Command("latte", "install", "pinentry")
}
