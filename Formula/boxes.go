package main

// Boxes - Draw boxes around text
// Homepage: https://boxes.thomasjensen.com/

import (
	"fmt"
	
	"os/exec"
)

func installBoxes() {
	// Método 1: Descargar y extraer .tar.gz
	boxes_tar_url := "https://github.com/ascii-boxes/boxes/archive/refs/tags/v2.3.0.tar.gz"
	boxes_cmd_tar := exec.Command("curl", "-L", boxes_tar_url, "-o", "package.tar.gz")
	err := boxes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boxes_zip_url := "https://github.com/ascii-boxes/boxes/archive/refs/tags/v2.3.0.zip"
	boxes_cmd_zip := exec.Command("curl", "-L", boxes_zip_url, "-o", "package.zip")
	err = boxes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boxes_bin_url := "https://github.com/ascii-boxes/boxes/archive/refs/tags/v2.3.0.bin"
	boxes_cmd_bin := exec.Command("curl", "-L", boxes_bin_url, "-o", "binary.bin")
	err = boxes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boxes_src_url := "https://github.com/ascii-boxes/boxes/archive/refs/tags/v2.3.0.src.tar.gz"
	boxes_cmd_src := exec.Command("curl", "-L", boxes_src_url, "-o", "source.tar.gz")
	err = boxes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boxes_cmd_direct := exec.Command("./binary")
	err = boxes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: libunistring")
	exec.Command("latte", "install", "libunistring").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
