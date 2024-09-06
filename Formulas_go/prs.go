package main

// Prs - Secure, fast & convenient password manager CLI with GPG & git sync
// Homepage: https://timvisee.com/projects/prs

import (
	"fmt"
	
	"os/exec"
)

func installPrs() {
	// Método 1: Descargar y extraer .tar.gz
	prs_tar_url := "https://github.com/timvisee/prs/archive/refs/tags/v0.5.1.tar.gz"
	prs_cmd_tar := exec.Command("curl", "-L", prs_tar_url, "-o", "package.tar.gz")
	err := prs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prs_zip_url := "https://github.com/timvisee/prs/archive/refs/tags/v0.5.1.zip"
	prs_cmd_zip := exec.Command("curl", "-L", prs_zip_url, "-o", "package.zip")
	err = prs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prs_bin_url := "https://github.com/timvisee/prs/archive/refs/tags/v0.5.1.bin"
	prs_cmd_bin := exec.Command("curl", "-L", prs_bin_url, "-o", "binary.bin")
	err = prs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prs_src_url := "https://github.com/timvisee/prs/archive/refs/tags/v0.5.1.src.tar.gz"
	prs_cmd_src := exec.Command("curl", "-L", prs_src_url, "-o", "source.tar.gz")
	err = prs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prs_cmd_direct := exec.Command("./binary")
	err = prs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
