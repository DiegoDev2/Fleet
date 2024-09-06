package main

// LastpassCli - LastPass command-line interface tool
// Homepage: https://github.com/lastpass/lastpass-cli

import (
	"fmt"
	
	"os/exec"
)

func installLastpassCli() {
	// Método 1: Descargar y extraer .tar.gz
	lastpasscli_tar_url := "https://github.com/lastpass/lastpass-cli/releases/download/v1.6.0/lastpass-cli-1.6.0.tar.gz"
	lastpasscli_cmd_tar := exec.Command("curl", "-L", lastpasscli_tar_url, "-o", "package.tar.gz")
	err := lastpasscli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lastpasscli_zip_url := "https://github.com/lastpass/lastpass-cli/releases/download/v1.6.0/lastpass-cli-1.6.0.zip"
	lastpasscli_cmd_zip := exec.Command("curl", "-L", lastpasscli_zip_url, "-o", "package.zip")
	err = lastpasscli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lastpasscli_bin_url := "https://github.com/lastpass/lastpass-cli/releases/download/v1.6.0/lastpass-cli-1.6.0.bin"
	lastpasscli_cmd_bin := exec.Command("curl", "-L", lastpasscli_bin_url, "-o", "binary.bin")
	err = lastpasscli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lastpasscli_src_url := "https://github.com/lastpass/lastpass-cli/releases/download/v1.6.0/lastpass-cli-1.6.0.src.tar.gz"
	lastpasscli_cmd_src := exec.Command("curl", "-L", lastpasscli_src_url, "-o", "source.tar.gz")
	err = lastpasscli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lastpasscli_cmd_direct := exec.Command("./binary")
	err = lastpasscli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
	exec.Command("latte", "install", "asciidoc").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pinentry")
	exec.Command("latte", "install", "pinentry").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
}
