package main

// Texapp - App.net client based on TTYtter
// Homepage: https://www.floodgap.com/software/texapp/

import (
	"fmt"
	
	"os/exec"
)

func installTexapp() {
	// Método 1: Descargar y extraer .tar.gz
	texapp_tar_url := "https://www.floodgap.com/software/texapp/dist0/0.6.11.txt"
	texapp_cmd_tar := exec.Command("curl", "-L", texapp_tar_url, "-o", "package.tar.gz")
	err := texapp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	texapp_zip_url := "https://www.floodgap.com/software/texapp/dist0/0.6.11.txt"
	texapp_cmd_zip := exec.Command("curl", "-L", texapp_zip_url, "-o", "package.zip")
	err = texapp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	texapp_bin_url := "https://www.floodgap.com/software/texapp/dist0/0.6.11.txt"
	texapp_cmd_bin := exec.Command("curl", "-L", texapp_bin_url, "-o", "binary.bin")
	err = texapp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	texapp_src_url := "https://www.floodgap.com/software/texapp/dist0/0.6.11.txt"
	texapp_cmd_src := exec.Command("curl", "-L", texapp_src_url, "-o", "source.tar.gz")
	err = texapp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	texapp_cmd_direct := exec.Command("./binary")
	err = texapp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
