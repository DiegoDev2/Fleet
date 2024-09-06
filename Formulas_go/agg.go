package main

// Agg - Asciicast to GIF converter
// Homepage: https://github.com/asciinema/agg

import (
	"fmt"
	
	"os/exec"
)

func installAgg() {
	// Método 1: Descargar y extraer .tar.gz
	agg_tar_url := "https://github.com/asciinema/agg/archive/refs/tags/v1.4.3.tar.gz"
	agg_cmd_tar := exec.Command("curl", "-L", agg_tar_url, "-o", "package.tar.gz")
	err := agg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	agg_zip_url := "https://github.com/asciinema/agg/archive/refs/tags/v1.4.3.zip"
	agg_cmd_zip := exec.Command("curl", "-L", agg_zip_url, "-o", "package.zip")
	err = agg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	agg_bin_url := "https://github.com/asciinema/agg/archive/refs/tags/v1.4.3.bin"
	agg_cmd_bin := exec.Command("curl", "-L", agg_bin_url, "-o", "binary.bin")
	err = agg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	agg_src_url := "https://github.com/asciinema/agg/archive/refs/tags/v1.4.3.src.tar.gz"
	agg_cmd_src := exec.Command("curl", "-L", agg_src_url, "-o", "source.tar.gz")
	err = agg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	agg_cmd_direct := exec.Command("./binary")
	err = agg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
