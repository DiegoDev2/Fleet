package main

// GetIplayer - Utility for downloading TV and radio programmes from BBC iPlayer
// Homepage: https://github.com/get-iplayer/get_iplayer

import (
	"fmt"
	
	"os/exec"
)

func installGetIplayer() {
	// Método 1: Descargar y extraer .tar.gz
	getiplayer_tar_url := "https://github.com/get-iplayer/get_iplayer/archive/refs/tags/v3.35.tar.gz"
	getiplayer_cmd_tar := exec.Command("curl", "-L", getiplayer_tar_url, "-o", "package.tar.gz")
	err := getiplayer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	getiplayer_zip_url := "https://github.com/get-iplayer/get_iplayer/archive/refs/tags/v3.35.zip"
	getiplayer_cmd_zip := exec.Command("curl", "-L", getiplayer_zip_url, "-o", "package.zip")
	err = getiplayer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	getiplayer_bin_url := "https://github.com/get-iplayer/get_iplayer/archive/refs/tags/v3.35.bin"
	getiplayer_cmd_bin := exec.Command("curl", "-L", getiplayer_bin_url, "-o", "binary.bin")
	err = getiplayer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	getiplayer_src_url := "https://github.com/get-iplayer/get_iplayer/archive/refs/tags/v3.35.src.tar.gz"
	getiplayer_cmd_src := exec.Command("curl", "-L", getiplayer_src_url, "-o", "source.tar.gz")
	err = getiplayer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	getiplayer_cmd_direct := exec.Command("./binary")
	err = getiplayer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: atomicparsley")
	exec.Command("latte", "install", "atomicparsley").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
