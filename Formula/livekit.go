package main

// Livekit - Scalable, high-performance WebRTC server
// Homepage: https://livekit.io

import (
	"fmt"
	
	"os/exec"
)

func installLivekit() {
	// Método 1: Descargar y extraer .tar.gz
	livekit_tar_url := "https://github.com/livekit/livekit/archive/refs/tags/v1.7.2.tar.gz"
	livekit_cmd_tar := exec.Command("curl", "-L", livekit_tar_url, "-o", "package.tar.gz")
	err := livekit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	livekit_zip_url := "https://github.com/livekit/livekit/archive/refs/tags/v1.7.2.zip"
	livekit_cmd_zip := exec.Command("curl", "-L", livekit_zip_url, "-o", "package.zip")
	err = livekit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	livekit_bin_url := "https://github.com/livekit/livekit/archive/refs/tags/v1.7.2.bin"
	livekit_cmd_bin := exec.Command("curl", "-L", livekit_bin_url, "-o", "binary.bin")
	err = livekit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	livekit_src_url := "https://github.com/livekit/livekit/archive/refs/tags/v1.7.2.src.tar.gz"
	livekit_cmd_src := exec.Command("curl", "-L", livekit_src_url, "-o", "source.tar.gz")
	err = livekit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	livekit_cmd_direct := exec.Command("./binary")
	err = livekit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
