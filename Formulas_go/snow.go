package main

// Snow - Whitespace steganography: coded messages using whitespace
// Homepage: https://darkside.com.au/snow/

import (
	"fmt"
	
	"os/exec"
)

func installSnow() {
	// Método 1: Descargar y extraer .tar.gz
	snow_tar_url := "https://darkside.com.au/snow/snow-20130616.tar.gz"
	snow_cmd_tar := exec.Command("curl", "-L", snow_tar_url, "-o", "package.tar.gz")
	err := snow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snow_zip_url := "https://darkside.com.au/snow/snow-20130616.zip"
	snow_cmd_zip := exec.Command("curl", "-L", snow_zip_url, "-o", "package.zip")
	err = snow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snow_bin_url := "https://darkside.com.au/snow/snow-20130616.bin"
	snow_cmd_bin := exec.Command("curl", "-L", snow_bin_url, "-o", "binary.bin")
	err = snow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snow_src_url := "https://darkside.com.au/snow/snow-20130616.src.tar.gz"
	snow_cmd_src := exec.Command("curl", "-L", snow_src_url, "-o", "source.tar.gz")
	err = snow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snow_cmd_direct := exec.Command("./binary")
	err = snow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
