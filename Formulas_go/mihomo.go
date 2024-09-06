package main

// Mihomo - Another rule-based tunnel in Go, formerly known as ClashMeta
// Homepage: https://wiki.metacubex.one

import (
	"fmt"
	
	"os/exec"
)

func installMihomo() {
	// Método 1: Descargar y extraer .tar.gz
	mihomo_tar_url := "https://github.com/MetaCubeX/mihomo/archive/refs/tags/v1.18.8.tar.gz"
	mihomo_cmd_tar := exec.Command("curl", "-L", mihomo_tar_url, "-o", "package.tar.gz")
	err := mihomo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mihomo_zip_url := "https://github.com/MetaCubeX/mihomo/archive/refs/tags/v1.18.8.zip"
	mihomo_cmd_zip := exec.Command("curl", "-L", mihomo_zip_url, "-o", "package.zip")
	err = mihomo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mihomo_bin_url := "https://github.com/MetaCubeX/mihomo/archive/refs/tags/v1.18.8.bin"
	mihomo_cmd_bin := exec.Command("curl", "-L", mihomo_bin_url, "-o", "binary.bin")
	err = mihomo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mihomo_src_url := "https://github.com/MetaCubeX/mihomo/archive/refs/tags/v1.18.8.src.tar.gz"
	mihomo_cmd_src := exec.Command("curl", "-L", mihomo_src_url, "-o", "source.tar.gz")
	err = mihomo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mihomo_cmd_direct := exec.Command("./binary")
	err = mihomo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
