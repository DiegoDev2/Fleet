package main

// Simg2img - Tool to convert Android sparse images to raw images and back
// Homepage: https://github.com/anestisb/android-simg2img

import (
	"fmt"
	
	"os/exec"
)

func installSimg2img() {
	// Método 1: Descargar y extraer .tar.gz
	simg2img_tar_url := "https://github.com/anestisb/android-simg2img/archive/refs/tags/1.1.4.tar.gz"
	simg2img_cmd_tar := exec.Command("curl", "-L", simg2img_tar_url, "-o", "package.tar.gz")
	err := simg2img_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simg2img_zip_url := "https://github.com/anestisb/android-simg2img/archive/refs/tags/1.1.4.zip"
	simg2img_cmd_zip := exec.Command("curl", "-L", simg2img_zip_url, "-o", "package.zip")
	err = simg2img_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simg2img_bin_url := "https://github.com/anestisb/android-simg2img/archive/refs/tags/1.1.4.bin"
	simg2img_cmd_bin := exec.Command("curl", "-L", simg2img_bin_url, "-o", "binary.bin")
	err = simg2img_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simg2img_src_url := "https://github.com/anestisb/android-simg2img/archive/refs/tags/1.1.4.src.tar.gz"
	simg2img_cmd_src := exec.Command("curl", "-L", simg2img_src_url, "-o", "source.tar.gz")
	err = simg2img_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simg2img_cmd_direct := exec.Command("./binary")
	err = simg2img_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
