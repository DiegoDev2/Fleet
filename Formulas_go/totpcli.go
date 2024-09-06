package main

// TotpCli - Authy/Google Authenticator like TOTP CLI tool written in Go
// Homepage: https://yitsushi.github.io/totp-cli/

import (
	"fmt"
	
	"os/exec"
)

func installTotpCli() {
	// Método 1: Descargar y extraer .tar.gz
	totpcli_tar_url := "https://github.com/yitsushi/totp-cli/archive/refs/tags/v1.8.7.tar.gz"
	totpcli_cmd_tar := exec.Command("curl", "-L", totpcli_tar_url, "-o", "package.tar.gz")
	err := totpcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	totpcli_zip_url := "https://github.com/yitsushi/totp-cli/archive/refs/tags/v1.8.7.zip"
	totpcli_cmd_zip := exec.Command("curl", "-L", totpcli_zip_url, "-o", "package.zip")
	err = totpcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	totpcli_bin_url := "https://github.com/yitsushi/totp-cli/archive/refs/tags/v1.8.7.bin"
	totpcli_cmd_bin := exec.Command("curl", "-L", totpcli_bin_url, "-o", "binary.bin")
	err = totpcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	totpcli_src_url := "https://github.com/yitsushi/totp-cli/archive/refs/tags/v1.8.7.src.tar.gz"
	totpcli_cmd_src := exec.Command("curl", "-L", totpcli_src_url, "-o", "source.tar.gz")
	err = totpcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	totpcli_cmd_direct := exec.Command("./binary")
	err = totpcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
