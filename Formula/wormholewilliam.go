package main

// WormholeWilliam - End-to-end encrypted file transfer
// Homepage: https://github.com/psanford/wormhole-william

import (
	"fmt"
	
	"os/exec"
)

func installWormholeWilliam() {
	// Método 1: Descargar y extraer .tar.gz
	wormholewilliam_tar_url := "https://github.com/psanford/wormhole-william/archive/refs/tags/v1.0.7.tar.gz"
	wormholewilliam_cmd_tar := exec.Command("curl", "-L", wormholewilliam_tar_url, "-o", "package.tar.gz")
	err := wormholewilliam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wormholewilliam_zip_url := "https://github.com/psanford/wormhole-william/archive/refs/tags/v1.0.7.zip"
	wormholewilliam_cmd_zip := exec.Command("curl", "-L", wormholewilliam_zip_url, "-o", "package.zip")
	err = wormholewilliam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wormholewilliam_bin_url := "https://github.com/psanford/wormhole-william/archive/refs/tags/v1.0.7.bin"
	wormholewilliam_cmd_bin := exec.Command("curl", "-L", wormholewilliam_bin_url, "-o", "binary.bin")
	err = wormholewilliam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wormholewilliam_src_url := "https://github.com/psanford/wormhole-william/archive/refs/tags/v1.0.7.src.tar.gz"
	wormholewilliam_cmd_src := exec.Command("curl", "-L", wormholewilliam_src_url, "-o", "source.tar.gz")
	err = wormholewilliam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wormholewilliam_cmd_direct := exec.Command("./binary")
	err = wormholewilliam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
