package main

// Gokey - Simple vaultless password manager in Go
// Homepage: https://github.com/cloudflare/gokey

import (
	"fmt"
	
	"os/exec"
)

func installGokey() {
	// Método 1: Descargar y extraer .tar.gz
	gokey_tar_url := "https://github.com/cloudflare/gokey/archive/refs/tags/v0.1.3.tar.gz"
	gokey_cmd_tar := exec.Command("curl", "-L", gokey_tar_url, "-o", "package.tar.gz")
	err := gokey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gokey_zip_url := "https://github.com/cloudflare/gokey/archive/refs/tags/v0.1.3.zip"
	gokey_cmd_zip := exec.Command("curl", "-L", gokey_zip_url, "-o", "package.zip")
	err = gokey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gokey_bin_url := "https://github.com/cloudflare/gokey/archive/refs/tags/v0.1.3.bin"
	gokey_cmd_bin := exec.Command("curl", "-L", gokey_bin_url, "-o", "binary.bin")
	err = gokey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gokey_src_url := "https://github.com/cloudflare/gokey/archive/refs/tags/v0.1.3.src.tar.gz"
	gokey_cmd_src := exec.Command("curl", "-L", gokey_src_url, "-o", "source.tar.gz")
	err = gokey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gokey_cmd_direct := exec.Command("./binary")
	err = gokey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: go-md2man")
	exec.Command("latte", "install", "go-md2man").Run()
}
