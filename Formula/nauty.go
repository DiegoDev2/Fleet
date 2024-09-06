package main

// Nauty - Automorphism groups of graphs and digraphs
// Homepage: https://pallini.di.uniroma1.it/

import (
	"fmt"
	
	"os/exec"
)

func installNauty() {
	// Método 1: Descargar y extraer .tar.gz
	nauty_tar_url := "https://pallini.di.uniroma1.it/nauty2_8_9.tar.gz"
	nauty_cmd_tar := exec.Command("curl", "-L", nauty_tar_url, "-o", "package.tar.gz")
	err := nauty_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nauty_zip_url := "https://pallini.di.uniroma1.it/nauty2_8_9.zip"
	nauty_cmd_zip := exec.Command("curl", "-L", nauty_zip_url, "-o", "package.zip")
	err = nauty_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nauty_bin_url := "https://pallini.di.uniroma1.it/nauty2_8_9.bin"
	nauty_cmd_bin := exec.Command("curl", "-L", nauty_bin_url, "-o", "binary.bin")
	err = nauty_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nauty_src_url := "https://pallini.di.uniroma1.it/nauty2_8_9.src.tar.gz"
	nauty_cmd_src := exec.Command("curl", "-L", nauty_src_url, "-o", "source.tar.gz")
	err = nauty_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nauty_cmd_direct := exec.Command("./binary")
	err = nauty_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
