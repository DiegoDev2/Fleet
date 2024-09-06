package main

// Flvstreamer - Stream audio and video from flash & RTMP Servers
// Homepage: https://www.nongnu.org/flvstreamer/

import (
	"fmt"
	
	"os/exec"
)

func installFlvstreamer() {
	// Método 1: Descargar y extraer .tar.gz
	flvstreamer_tar_url := "https://download.savannah.gnu.org/releases/flvstreamer/source/flvstreamer-2.1c1.tar.gz"
	flvstreamer_cmd_tar := exec.Command("curl", "-L", flvstreamer_tar_url, "-o", "package.tar.gz")
	err := flvstreamer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flvstreamer_zip_url := "https://download.savannah.gnu.org/releases/flvstreamer/source/flvstreamer-2.1c1.zip"
	flvstreamer_cmd_zip := exec.Command("curl", "-L", flvstreamer_zip_url, "-o", "package.zip")
	err = flvstreamer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flvstreamer_bin_url := "https://download.savannah.gnu.org/releases/flvstreamer/source/flvstreamer-2.1c1.bin"
	flvstreamer_cmd_bin := exec.Command("curl", "-L", flvstreamer_bin_url, "-o", "binary.bin")
	err = flvstreamer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flvstreamer_src_url := "https://download.savannah.gnu.org/releases/flvstreamer/source/flvstreamer-2.1c1.src.tar.gz"
	flvstreamer_cmd_src := exec.Command("curl", "-L", flvstreamer_src_url, "-o", "source.tar.gz")
	err = flvstreamer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flvstreamer_cmd_direct := exec.Command("./binary")
	err = flvstreamer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
