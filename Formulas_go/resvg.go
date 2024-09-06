package main

// Resvg - SVG rendering tool and library
// Homepage: https://github.com/RazrFalcon/resvg

import (
	"fmt"
	
	"os/exec"
)

func installResvg() {
	// Método 1: Descargar y extraer .tar.gz
	resvg_tar_url := "https://github.com/RazrFalcon/resvg/archive/refs/tags/v0.43.0.tar.gz"
	resvg_cmd_tar := exec.Command("curl", "-L", resvg_tar_url, "-o", "package.tar.gz")
	err := resvg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	resvg_zip_url := "https://github.com/RazrFalcon/resvg/archive/refs/tags/v0.43.0.zip"
	resvg_cmd_zip := exec.Command("curl", "-L", resvg_zip_url, "-o", "package.zip")
	err = resvg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	resvg_bin_url := "https://github.com/RazrFalcon/resvg/archive/refs/tags/v0.43.0.bin"
	resvg_cmd_bin := exec.Command("curl", "-L", resvg_bin_url, "-o", "binary.bin")
	err = resvg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	resvg_src_url := "https://github.com/RazrFalcon/resvg/archive/refs/tags/v0.43.0.src.tar.gz"
	resvg_cmd_src := exec.Command("curl", "-L", resvg_src_url, "-o", "source.tar.gz")
	err = resvg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	resvg_cmd_direct := exec.Command("./binary")
	err = resvg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
