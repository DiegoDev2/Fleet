package main

// WebtorrentCli - Command-line streaming torrent client
// Homepage: https://webtorrent.io/

import (
	"fmt"
	
	"os/exec"
)

func installWebtorrentCli() {
	// Método 1: Descargar y extraer .tar.gz
	webtorrentcli_tar_url := "https://registry.npmjs.org/webtorrent-cli/-/webtorrent-cli-4.1.0.tgz"
	webtorrentcli_cmd_tar := exec.Command("curl", "-L", webtorrentcli_tar_url, "-o", "package.tar.gz")
	err := webtorrentcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webtorrentcli_zip_url := "https://registry.npmjs.org/webtorrent-cli/-/webtorrent-cli-4.1.0.tgz"
	webtorrentcli_cmd_zip := exec.Command("curl", "-L", webtorrentcli_zip_url, "-o", "package.zip")
	err = webtorrentcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webtorrentcli_bin_url := "https://registry.npmjs.org/webtorrent-cli/-/webtorrent-cli-4.1.0.tgz"
	webtorrentcli_cmd_bin := exec.Command("curl", "-L", webtorrentcli_bin_url, "-o", "binary.bin")
	err = webtorrentcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webtorrentcli_src_url := "https://registry.npmjs.org/webtorrent-cli/-/webtorrent-cli-4.1.0.tgz"
	webtorrentcli_cmd_src := exec.Command("curl", "-L", webtorrentcli_src_url, "-o", "source.tar.gz")
	err = webtorrentcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webtorrentcli_cmd_direct := exec.Command("./binary")
	err = webtorrentcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
