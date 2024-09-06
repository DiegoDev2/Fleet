package main

// Gaze - Execute commands for you
// Homepage: https://github.com/wtetsu/gaze

import (
	"fmt"
	
	"os/exec"
)

func installGaze() {
	// Método 1: Descargar y extraer .tar.gz
	gaze_tar_url := "https://github.com/wtetsu/gaze/archive/refs/tags/v1.1.6.tar.gz"
	gaze_cmd_tar := exec.Command("curl", "-L", gaze_tar_url, "-o", "package.tar.gz")
	err := gaze_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gaze_zip_url := "https://github.com/wtetsu/gaze/archive/refs/tags/v1.1.6.zip"
	gaze_cmd_zip := exec.Command("curl", "-L", gaze_zip_url, "-o", "package.zip")
	err = gaze_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gaze_bin_url := "https://github.com/wtetsu/gaze/archive/refs/tags/v1.1.6.bin"
	gaze_cmd_bin := exec.Command("curl", "-L", gaze_bin_url, "-o", "binary.bin")
	err = gaze_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gaze_src_url := "https://github.com/wtetsu/gaze/archive/refs/tags/v1.1.6.src.tar.gz"
	gaze_cmd_src := exec.Command("curl", "-L", gaze_src_url, "-o", "source.tar.gz")
	err = gaze_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gaze_cmd_direct := exec.Command("./binary")
	err = gaze_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
