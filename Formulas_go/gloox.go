package main

// Gloox - C++ Jabber/XMPP library that handles the low-level protocol
// Homepage: https://camaya.net/gloox/

import (
	"fmt"
	
	"os/exec"
)

func installGloox() {
	// Método 1: Descargar y extraer .tar.gz
	gloox_tar_url := "https://camaya.net/download/gloox-1.0.28.tar.bz2"
	gloox_cmd_tar := exec.Command("curl", "-L", gloox_tar_url, "-o", "package.tar.gz")
	err := gloox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gloox_zip_url := "https://camaya.net/download/gloox-1.0.28.tar.bz2"
	gloox_cmd_zip := exec.Command("curl", "-L", gloox_zip_url, "-o", "package.zip")
	err = gloox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gloox_bin_url := "https://camaya.net/download/gloox-1.0.28.tar.bz2"
	gloox_cmd_bin := exec.Command("curl", "-L", gloox_bin_url, "-o", "binary.bin")
	err = gloox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gloox_src_url := "https://camaya.net/download/gloox-1.0.28.tar.bz2"
	gloox_cmd_src := exec.Command("curl", "-L", gloox_src_url, "-o", "source.tar.gz")
	err = gloox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gloox_cmd_direct := exec.Command("./binary")
	err = gloox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libidn")
exec.Command("latte", "install", "libidn")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
