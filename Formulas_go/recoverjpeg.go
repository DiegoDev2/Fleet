package main

// Recoverjpeg - Tool to recover JPEG images from a file system image
// Homepage: https://rfc1149.net/devel/recoverjpeg.html

import (
	"fmt"
	
	"os/exec"
)

func installRecoverjpeg() {
	// Método 1: Descargar y extraer .tar.gz
	recoverjpeg_tar_url := "https://rfc1149.net/download/recoverjpeg/recoverjpeg-2.6.3.tar.gz"
	recoverjpeg_cmd_tar := exec.Command("curl", "-L", recoverjpeg_tar_url, "-o", "package.tar.gz")
	err := recoverjpeg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	recoverjpeg_zip_url := "https://rfc1149.net/download/recoverjpeg/recoverjpeg-2.6.3.zip"
	recoverjpeg_cmd_zip := exec.Command("curl", "-L", recoverjpeg_zip_url, "-o", "package.zip")
	err = recoverjpeg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	recoverjpeg_bin_url := "https://rfc1149.net/download/recoverjpeg/recoverjpeg-2.6.3.bin"
	recoverjpeg_cmd_bin := exec.Command("curl", "-L", recoverjpeg_bin_url, "-o", "binary.bin")
	err = recoverjpeg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	recoverjpeg_src_url := "https://rfc1149.net/download/recoverjpeg/recoverjpeg-2.6.3.src.tar.gz"
	recoverjpeg_cmd_src := exec.Command("curl", "-L", recoverjpeg_src_url, "-o", "source.tar.gz")
	err = recoverjpeg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	recoverjpeg_cmd_direct := exec.Command("./binary")
	err = recoverjpeg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
