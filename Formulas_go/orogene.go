package main

// Orogene - `node_modules/` package manager and utility toolkit
// Homepage: https://orogene.dev

import (
	"fmt"
	
	"os/exec"
)

func installOrogene() {
	// Método 1: Descargar y extraer .tar.gz
	orogene_tar_url := "https://github.com/orogene/orogene/archive/refs/tags/v0.3.34.tar.gz"
	orogene_cmd_tar := exec.Command("curl", "-L", orogene_tar_url, "-o", "package.tar.gz")
	err := orogene_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	orogene_zip_url := "https://github.com/orogene/orogene/archive/refs/tags/v0.3.34.zip"
	orogene_cmd_zip := exec.Command("curl", "-L", orogene_zip_url, "-o", "package.zip")
	err = orogene_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	orogene_bin_url := "https://github.com/orogene/orogene/archive/refs/tags/v0.3.34.bin"
	orogene_cmd_bin := exec.Command("curl", "-L", orogene_bin_url, "-o", "binary.bin")
	err = orogene_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	orogene_src_url := "https://github.com/orogene/orogene/archive/refs/tags/v0.3.34.src.tar.gz"
	orogene_cmd_src := exec.Command("curl", "-L", orogene_src_url, "-o", "source.tar.gz")
	err = orogene_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	orogene_cmd_direct := exec.Command("./binary")
	err = orogene_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
