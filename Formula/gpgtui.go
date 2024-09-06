package main

// GpgTui - Manage your GnuPG keys with ease!
// Homepage: https://github.com/orhun/gpg-tui

import (
	"fmt"
	
	"os/exec"
)

func installGpgTui() {
	// Método 1: Descargar y extraer .tar.gz
	gpgtui_tar_url := "https://github.com/orhun/gpg-tui/archive/refs/tags/v0.11.0.tar.gz"
	gpgtui_cmd_tar := exec.Command("curl", "-L", gpgtui_tar_url, "-o", "package.tar.gz")
	err := gpgtui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpgtui_zip_url := "https://github.com/orhun/gpg-tui/archive/refs/tags/v0.11.0.zip"
	gpgtui_cmd_zip := exec.Command("curl", "-L", gpgtui_zip_url, "-o", "package.zip")
	err = gpgtui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpgtui_bin_url := "https://github.com/orhun/gpg-tui/archive/refs/tags/v0.11.0.bin"
	gpgtui_cmd_bin := exec.Command("curl", "-L", gpgtui_bin_url, "-o", "binary.bin")
	err = gpgtui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpgtui_src_url := "https://github.com/orhun/gpg-tui/archive/refs/tags/v0.11.0.src.tar.gz"
	gpgtui_cmd_src := exec.Command("curl", "-L", gpgtui_src_url, "-o", "source.tar.gz")
	err = gpgtui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpgtui_cmd_direct := exec.Command("./binary")
	err = gpgtui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: gpgme")
	exec.Command("latte", "install", "gpgme").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
}
