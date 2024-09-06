package main

// Gopass - Slightly more awesome Standard Unix Password Manager for Teams
// Homepage: https://github.com/gopasspw/gopass

import (
	"fmt"
	
	"os/exec"
)

func installGopass() {
	// Método 1: Descargar y extraer .tar.gz
	gopass_tar_url := "https://github.com/gopasspw/gopass/releases/download/v1.15.14/gopass-1.15.14.tar.gz"
	gopass_cmd_tar := exec.Command("curl", "-L", gopass_tar_url, "-o", "package.tar.gz")
	err := gopass_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gopass_zip_url := "https://github.com/gopasspw/gopass/releases/download/v1.15.14/gopass-1.15.14.zip"
	gopass_cmd_zip := exec.Command("curl", "-L", gopass_zip_url, "-o", "package.zip")
	err = gopass_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gopass_bin_url := "https://github.com/gopasspw/gopass/releases/download/v1.15.14/gopass-1.15.14.bin"
	gopass_cmd_bin := exec.Command("curl", "-L", gopass_bin_url, "-o", "binary.bin")
	err = gopass_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gopass_src_url := "https://github.com/gopasspw/gopass/releases/download/v1.15.14/gopass-1.15.14.src.tar.gz"
	gopass_cmd_src := exec.Command("curl", "-L", gopass_src_url, "-o", "source.tar.gz")
	err = gopass_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gopass_cmd_direct := exec.Command("./binary")
	err = gopass_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: terminal-notifier")
	exec.Command("latte", "install", "terminal-notifier").Run()
}
