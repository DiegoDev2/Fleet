package main

// Typespeed - Zap words flying across the screen by typing them correctly
// Homepage: https://typespeed.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installTypespeed() {
	// Método 1: Descargar y extraer .tar.gz
	typespeed_tar_url := "https://downloads.sourceforge.net/project/typespeed/typespeed/0.6.5/typespeed-0.6.5.tar.gz"
	typespeed_cmd_tar := exec.Command("curl", "-L", typespeed_tar_url, "-o", "package.tar.gz")
	err := typespeed_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typespeed_zip_url := "https://downloads.sourceforge.net/project/typespeed/typespeed/0.6.5/typespeed-0.6.5.zip"
	typespeed_cmd_zip := exec.Command("curl", "-L", typespeed_zip_url, "-o", "package.zip")
	err = typespeed_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typespeed_bin_url := "https://downloads.sourceforge.net/project/typespeed/typespeed/0.6.5/typespeed-0.6.5.bin"
	typespeed_cmd_bin := exec.Command("curl", "-L", typespeed_bin_url, "-o", "binary.bin")
	err = typespeed_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typespeed_src_url := "https://downloads.sourceforge.net/project/typespeed/typespeed/0.6.5/typespeed-0.6.5.src.tar.gz"
	typespeed_cmd_src := exec.Command("curl", "-L", typespeed_src_url, "-o", "source.tar.gz")
	err = typespeed_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typespeed_cmd_direct := exec.Command("./binary")
	err = typespeed_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
