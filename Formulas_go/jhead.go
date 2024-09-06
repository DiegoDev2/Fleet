package main

// Jhead - Extract Digicam setting info from EXIF JPEG headers
// Homepage: https://github.com/Matthias-Wandel/jhead

import (
	"fmt"
	
	"os/exec"
)

func installJhead() {
	// Método 1: Descargar y extraer .tar.gz
	jhead_tar_url := "https://github.com/Matthias-Wandel/jhead/archive/refs/tags/3.08.tar.gz"
	jhead_cmd_tar := exec.Command("curl", "-L", jhead_tar_url, "-o", "package.tar.gz")
	err := jhead_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jhead_zip_url := "https://github.com/Matthias-Wandel/jhead/archive/refs/tags/3.08.zip"
	jhead_cmd_zip := exec.Command("curl", "-L", jhead_zip_url, "-o", "package.zip")
	err = jhead_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jhead_bin_url := "https://github.com/Matthias-Wandel/jhead/archive/refs/tags/3.08.bin"
	jhead_cmd_bin := exec.Command("curl", "-L", jhead_bin_url, "-o", "binary.bin")
	err = jhead_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jhead_src_url := "https://github.com/Matthias-Wandel/jhead/archive/refs/tags/3.08.src.tar.gz"
	jhead_cmd_src := exec.Command("curl", "-L", jhead_src_url, "-o", "source.tar.gz")
	err = jhead_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jhead_cmd_direct := exec.Command("./binary")
	err = jhead_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
