package main

// Scale2x - Real-time graphics effect
// Homepage: https://www.scale2x.it/

import (
	"fmt"
	
	"os/exec"
)

func installScale2x() {
	// Método 1: Descargar y extraer .tar.gz
	scale2x_tar_url := "https://github.com/amadvance/scale2x/releases/download/v4.0/scale2x-4.0.tar.gz"
	scale2x_cmd_tar := exec.Command("curl", "-L", scale2x_tar_url, "-o", "package.tar.gz")
	err := scale2x_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scale2x_zip_url := "https://github.com/amadvance/scale2x/releases/download/v4.0/scale2x-4.0.zip"
	scale2x_cmd_zip := exec.Command("curl", "-L", scale2x_zip_url, "-o", "package.zip")
	err = scale2x_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scale2x_bin_url := "https://github.com/amadvance/scale2x/releases/download/v4.0/scale2x-4.0.bin"
	scale2x_cmd_bin := exec.Command("curl", "-L", scale2x_bin_url, "-o", "binary.bin")
	err = scale2x_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scale2x_src_url := "https://github.com/amadvance/scale2x/releases/download/v4.0/scale2x-4.0.src.tar.gz"
	scale2x_cmd_src := exec.Command("curl", "-L", scale2x_src_url, "-o", "source.tar.gz")
	err = scale2x_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scale2x_cmd_direct := exec.Command("./binary")
	err = scale2x_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
