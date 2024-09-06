package main

// QbittorrentCli - Command-line interface for qBittorrent written in Go
// Homepage: https://github.com/ludviglundgren/qbittorrent-cli

import (
	"fmt"
	
	"os/exec"
)

func installQbittorrentCli() {
	// Método 1: Descargar y extraer .tar.gz
	qbittorrentcli_tar_url := "https://github.com/ludviglundgren/qbittorrent-cli/archive/refs/tags/v2.0.0.tar.gz"
	qbittorrentcli_cmd_tar := exec.Command("curl", "-L", qbittorrentcli_tar_url, "-o", "package.tar.gz")
	err := qbittorrentcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qbittorrentcli_zip_url := "https://github.com/ludviglundgren/qbittorrent-cli/archive/refs/tags/v2.0.0.zip"
	qbittorrentcli_cmd_zip := exec.Command("curl", "-L", qbittorrentcli_zip_url, "-o", "package.zip")
	err = qbittorrentcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qbittorrentcli_bin_url := "https://github.com/ludviglundgren/qbittorrent-cli/archive/refs/tags/v2.0.0.bin"
	qbittorrentcli_cmd_bin := exec.Command("curl", "-L", qbittorrentcli_bin_url, "-o", "binary.bin")
	err = qbittorrentcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qbittorrentcli_src_url := "https://github.com/ludviglundgren/qbittorrent-cli/archive/refs/tags/v2.0.0.src.tar.gz"
	qbittorrentcli_cmd_src := exec.Command("curl", "-L", qbittorrentcli_src_url, "-o", "source.tar.gz")
	err = qbittorrentcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qbittorrentcli_cmd_direct := exec.Command("./binary")
	err = qbittorrentcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
