package main

// Mdzk - Plain text Zettelkasten based on mdBook
// Homepage: https://mdzk.app/

import (
	"fmt"
	
	"os/exec"
)

func installMdzk() {
	// Método 1: Descargar y extraer .tar.gz
	mdzk_tar_url := "https://github.com/mdzk-rs/mdzk/archive/refs/tags/0.5.2.tar.gz"
	mdzk_cmd_tar := exec.Command("curl", "-L", mdzk_tar_url, "-o", "package.tar.gz")
	err := mdzk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdzk_zip_url := "https://github.com/mdzk-rs/mdzk/archive/refs/tags/0.5.2.zip"
	mdzk_cmd_zip := exec.Command("curl", "-L", mdzk_zip_url, "-o", "package.zip")
	err = mdzk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdzk_bin_url := "https://github.com/mdzk-rs/mdzk/archive/refs/tags/0.5.2.bin"
	mdzk_cmd_bin := exec.Command("curl", "-L", mdzk_bin_url, "-o", "binary.bin")
	err = mdzk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdzk_src_url := "https://github.com/mdzk-rs/mdzk/archive/refs/tags/0.5.2.src.tar.gz"
	mdzk_cmd_src := exec.Command("curl", "-L", mdzk_src_url, "-o", "source.tar.gz")
	err = mdzk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdzk_cmd_direct := exec.Command("./binary")
	err = mdzk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
