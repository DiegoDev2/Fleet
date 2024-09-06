package main

// NowplayingCli - Retrieves currently playing media, and simulates media actions
// Homepage: https://github.com/kirtan-shah/nowplaying-cli

import (
	"fmt"
	
	"os/exec"
)

func installNowplayingCli() {
	// Método 1: Descargar y extraer .tar.gz
	nowplayingcli_tar_url := "https://github.com/kirtan-shah/nowplaying-cli/archive/refs/tags/v1.2.1.tar.gz"
	nowplayingcli_cmd_tar := exec.Command("curl", "-L", nowplayingcli_tar_url, "-o", "package.tar.gz")
	err := nowplayingcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nowplayingcli_zip_url := "https://github.com/kirtan-shah/nowplaying-cli/archive/refs/tags/v1.2.1.zip"
	nowplayingcli_cmd_zip := exec.Command("curl", "-L", nowplayingcli_zip_url, "-o", "package.zip")
	err = nowplayingcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nowplayingcli_bin_url := "https://github.com/kirtan-shah/nowplaying-cli/archive/refs/tags/v1.2.1.bin"
	nowplayingcli_cmd_bin := exec.Command("curl", "-L", nowplayingcli_bin_url, "-o", "binary.bin")
	err = nowplayingcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nowplayingcli_src_url := "https://github.com/kirtan-shah/nowplaying-cli/archive/refs/tags/v1.2.1.src.tar.gz"
	nowplayingcli_cmd_src := exec.Command("curl", "-L", nowplayingcli_src_url, "-o", "source.tar.gz")
	err = nowplayingcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nowplayingcli_cmd_direct := exec.Command("./binary")
	err = nowplayingcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
