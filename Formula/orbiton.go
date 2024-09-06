package main

// Orbiton - Fast and config-free text editor and IDE limited by VT100
// Homepage: https://roboticoverlords.org/orbiton/

import (
	"fmt"
	
	"os/exec"
)

func installOrbiton() {
	// Método 1: Descargar y extraer .tar.gz
	orbiton_tar_url := "https://github.com/xyproto/orbiton/archive/refs/tags/v2.67.0.tar.gz"
	orbiton_cmd_tar := exec.Command("curl", "-L", orbiton_tar_url, "-o", "package.tar.gz")
	err := orbiton_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	orbiton_zip_url := "https://github.com/xyproto/orbiton/archive/refs/tags/v2.67.0.zip"
	orbiton_cmd_zip := exec.Command("curl", "-L", orbiton_zip_url, "-o", "package.zip")
	err = orbiton_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	orbiton_bin_url := "https://github.com/xyproto/orbiton/archive/refs/tags/v2.67.0.bin"
	orbiton_cmd_bin := exec.Command("curl", "-L", orbiton_bin_url, "-o", "binary.bin")
	err = orbiton_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	orbiton_src_url := "https://github.com/xyproto/orbiton/archive/refs/tags/v2.67.0.src.tar.gz"
	orbiton_cmd_src := exec.Command("curl", "-L", orbiton_src_url, "-o", "source.tar.gz")
	err = orbiton_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	orbiton_cmd_direct := exec.Command("./binary")
	err = orbiton_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: xorg-server")
	exec.Command("latte", "install", "xorg-server").Run()
	fmt.Println("Instalando dependencia: xclip")
	exec.Command("latte", "install", "xclip").Run()
}
