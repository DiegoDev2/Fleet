package main

// Mktorrent - Create BitTorrent metainfo files
// Homepage: https://github.com/pobrn/mktorrent/wiki

import (
	"fmt"
	
	"os/exec"
)

func installMktorrent() {
	// Método 1: Descargar y extraer .tar.gz
	mktorrent_tar_url := "https://github.com/pobrn/mktorrent/archive/refs/tags/v1.1.tar.gz"
	mktorrent_cmd_tar := exec.Command("curl", "-L", mktorrent_tar_url, "-o", "package.tar.gz")
	err := mktorrent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mktorrent_zip_url := "https://github.com/pobrn/mktorrent/archive/refs/tags/v1.1.zip"
	mktorrent_cmd_zip := exec.Command("curl", "-L", mktorrent_zip_url, "-o", "package.zip")
	err = mktorrent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mktorrent_bin_url := "https://github.com/pobrn/mktorrent/archive/refs/tags/v1.1.bin"
	mktorrent_cmd_bin := exec.Command("curl", "-L", mktorrent_bin_url, "-o", "binary.bin")
	err = mktorrent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mktorrent_src_url := "https://github.com/pobrn/mktorrent/archive/refs/tags/v1.1.src.tar.gz"
	mktorrent_cmd_src := exec.Command("curl", "-L", mktorrent_src_url, "-o", "source.tar.gz")
	err = mktorrent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mktorrent_cmd_direct := exec.Command("./binary")
	err = mktorrent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
