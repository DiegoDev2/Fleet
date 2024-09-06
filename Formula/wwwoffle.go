package main

// Wwwoffle - Better browsing for computers with intermittent connections
// Homepage: https://www.gedanken.org.uk/software/wwwoffle/

import (
	"fmt"
	
	"os/exec"
)

func installWwwoffle() {
	// Método 1: Descargar y extraer .tar.gz
	wwwoffle_tar_url := "https://www.gedanken.org.uk/software/wwwoffle/download/wwwoffle-2.9j.tgz"
	wwwoffle_cmd_tar := exec.Command("curl", "-L", wwwoffle_tar_url, "-o", "package.tar.gz")
	err := wwwoffle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wwwoffle_zip_url := "https://www.gedanken.org.uk/software/wwwoffle/download/wwwoffle-2.9j.tgz"
	wwwoffle_cmd_zip := exec.Command("curl", "-L", wwwoffle_zip_url, "-o", "package.zip")
	err = wwwoffle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wwwoffle_bin_url := "https://www.gedanken.org.uk/software/wwwoffle/download/wwwoffle-2.9j.tgz"
	wwwoffle_cmd_bin := exec.Command("curl", "-L", wwwoffle_bin_url, "-o", "binary.bin")
	err = wwwoffle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wwwoffle_src_url := "https://www.gedanken.org.uk/software/wwwoffle/download/wwwoffle-2.9j.tgz"
	wwwoffle_cmd_src := exec.Command("curl", "-L", wwwoffle_src_url, "-o", "source.tar.gz")
	err = wwwoffle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wwwoffle_cmd_direct := exec.Command("./binary")
	err = wwwoffle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
