package main

// GolangciLintLangserver - Language server for `golangci-lint`
// Homepage: https://github.com/nametake/golangci-lint-langserver

import (
	"fmt"
	
	"os/exec"
)

func installGolangciLintLangserver() {
	// Método 1: Descargar y extraer .tar.gz
	golangcilintlangserver_tar_url := "https://github.com/nametake/golangci-lint-langserver/archive/refs/tags/v0.0.9.tar.gz"
	golangcilintlangserver_cmd_tar := exec.Command("curl", "-L", golangcilintlangserver_tar_url, "-o", "package.tar.gz")
	err := golangcilintlangserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	golangcilintlangserver_zip_url := "https://github.com/nametake/golangci-lint-langserver/archive/refs/tags/v0.0.9.zip"
	golangcilintlangserver_cmd_zip := exec.Command("curl", "-L", golangcilintlangserver_zip_url, "-o", "package.zip")
	err = golangcilintlangserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	golangcilintlangserver_bin_url := "https://github.com/nametake/golangci-lint-langserver/archive/refs/tags/v0.0.9.bin"
	golangcilintlangserver_cmd_bin := exec.Command("curl", "-L", golangcilintlangserver_bin_url, "-o", "binary.bin")
	err = golangcilintlangserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	golangcilintlangserver_src_url := "https://github.com/nametake/golangci-lint-langserver/archive/refs/tags/v0.0.9.src.tar.gz"
	golangcilintlangserver_cmd_src := exec.Command("curl", "-L", golangcilintlangserver_src_url, "-o", "source.tar.gz")
	err = golangcilintlangserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	golangcilintlangserver_cmd_direct := exec.Command("./binary")
	err = golangcilintlangserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
