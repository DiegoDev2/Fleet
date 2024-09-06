package main

// TelegramDownloader - Telegram Messenger downloader/tools written in Golang
// Homepage: https://docs.iyear.me/tdl/

import (
	"fmt"
	
	"os/exec"
)

func installTelegramDownloader() {
	// Método 1: Descargar y extraer .tar.gz
	telegramdownloader_tar_url := "https://github.com/iyear/tdl/archive/refs/tags/v0.17.3.tar.gz"
	telegramdownloader_cmd_tar := exec.Command("curl", "-L", telegramdownloader_tar_url, "-o", "package.tar.gz")
	err := telegramdownloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	telegramdownloader_zip_url := "https://github.com/iyear/tdl/archive/refs/tags/v0.17.3.zip"
	telegramdownloader_cmd_zip := exec.Command("curl", "-L", telegramdownloader_zip_url, "-o", "package.zip")
	err = telegramdownloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	telegramdownloader_bin_url := "https://github.com/iyear/tdl/archive/refs/tags/v0.17.3.bin"
	telegramdownloader_cmd_bin := exec.Command("curl", "-L", telegramdownloader_bin_url, "-o", "binary.bin")
	err = telegramdownloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	telegramdownloader_src_url := "https://github.com/iyear/tdl/archive/refs/tags/v0.17.3.src.tar.gz"
	telegramdownloader_cmd_src := exec.Command("curl", "-L", telegramdownloader_src_url, "-o", "source.tar.gz")
	err = telegramdownloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	telegramdownloader_cmd_direct := exec.Command("./binary")
	err = telegramdownloader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
