package main

// Vhs - Your CLI home video recorder
// Homepage: https://github.com/charmbracelet/vhs

import (
	"fmt"
	
	"os/exec"
)

func installVhs() {
	// Método 1: Descargar y extraer .tar.gz
	vhs_tar_url := "https://github.com/charmbracelet/vhs/archive/refs/tags/v0.7.2.tar.gz"
	vhs_cmd_tar := exec.Command("curl", "-L", vhs_tar_url, "-o", "package.tar.gz")
	err := vhs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vhs_zip_url := "https://github.com/charmbracelet/vhs/archive/refs/tags/v0.7.2.zip"
	vhs_cmd_zip := exec.Command("curl", "-L", vhs_zip_url, "-o", "package.zip")
	err = vhs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vhs_bin_url := "https://github.com/charmbracelet/vhs/archive/refs/tags/v0.7.2.bin"
	vhs_cmd_bin := exec.Command("curl", "-L", vhs_bin_url, "-o", "binary.bin")
	err = vhs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vhs_src_url := "https://github.com/charmbracelet/vhs/archive/refs/tags/v0.7.2.src.tar.gz"
	vhs_cmd_src := exec.Command("curl", "-L", vhs_src_url, "-o", "source.tar.gz")
	err = vhs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vhs_cmd_direct := exec.Command("./binary")
	err = vhs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: ttyd")
	exec.Command("latte", "install", "ttyd").Run()
}
