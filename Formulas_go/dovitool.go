package main

// DoviTool - CLI tool for Dolby Vision metadata on video streams
// Homepage: https://github.com/quietvoid/dovi_tool/

import (
	"fmt"
	
	"os/exec"
)

func installDoviTool() {
	// Método 1: Descargar y extraer .tar.gz
	dovitool_tar_url := "https://github.com/quietvoid/dovi_tool/archive/refs/tags/2.1.2.tar.gz"
	dovitool_cmd_tar := exec.Command("curl", "-L", dovitool_tar_url, "-o", "package.tar.gz")
	err := dovitool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dovitool_zip_url := "https://github.com/quietvoid/dovi_tool/archive/refs/tags/2.1.2.zip"
	dovitool_cmd_zip := exec.Command("curl", "-L", dovitool_zip_url, "-o", "package.zip")
	err = dovitool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dovitool_bin_url := "https://github.com/quietvoid/dovi_tool/archive/refs/tags/2.1.2.bin"
	dovitool_cmd_bin := exec.Command("curl", "-L", dovitool_bin_url, "-o", "binary.bin")
	err = dovitool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dovitool_src_url := "https://github.com/quietvoid/dovi_tool/archive/refs/tags/2.1.2.src.tar.gz"
	dovitool_cmd_src := exec.Command("curl", "-L", dovitool_src_url, "-o", "source.tar.gz")
	err = dovitool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dovitool_cmd_direct := exec.Command("./binary")
	err = dovitool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
}
